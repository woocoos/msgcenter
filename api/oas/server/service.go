package server

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/web/handler"
	"github.com/woocoos/msgcenter/api/oas"
	"github.com/woocoos/msgcenter/dispatch"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/metrics"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/silence"
	"net/http"
	"regexp"
	"runtime"
	"sort"
	"sync"
	"time"
)

type (
	getAlertStatusFn func(label.Fingerprint) alert.MarkerStatus
	setAlertStatusFn func(label.LabelSet)
)

var (
	// check implemented by the API.
	_ oas.GeneralServer  = (*Service)(nil)
	_ oas.ReceiverServer = (*Service)(nil)
	_ oas.SilenceServer  = (*Service)(nil)
	_ oas.AlertServer    = (*Service)(nil)

	logger = log.Component("api")
)

type Service struct {
	*Options

	uptime time.Time

	// mtx protects alertmanagerConfig, setAlertStatus and route.
	mtx sync.RWMutex

	route          *dispatch.Route
	setAlertStatus setAlertStatusFn

	m *metrics.Alerts
}

func RegisterHandlers(router *gin.RouterGroup, srv *Service) {
	RegisterGeneralHandlers(router, srv)
	RegisterReceiverHandlers(router, srv)
	RegisterSilenceHandlers(router, srv)
	RegisterAlertHandlers(router, srv)
}

func NewService(opts *Options) (*Service, error) {
	if err := opts.Validate(); err != nil {
		return nil, fmt.Errorf("invalid API options: %s", err)
	}
	s := &Service{
		Options: opts,
		uptime:  time.Now(),
		m:       metrics.NewAlerts("v2"),
	}
	return s, nil
}

func (s *Service) Update(cfg *profile.Config, fn func(label.LabelSet)) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.setAlertStatus = fn
	s.route = dispatch.NewRoute(cfg.Route, nil)
}

// GetStatus returns the status of the Alertmanager.
// Notice: return all config include all tenants,so it's not safe for production
func (s *Service) GetStatus(c *gin.Context) (resp *oas.AlertmanagerStatus, err error) {
	original := s.Coordinator.ProfileString()
	resp = &oas.AlertmanagerStatus{
		Uptime: s.uptime,
		VersionInfo: oas.VersionInfo{
			Version:   conf.Global().Version(),
			Revision:  "",
			Branch:    "",
			BuildUser: "",
			BuildDate: "",
			GoVersion: runtime.Version(),
		},
		Config: oas.AlertmanagerConfig{Original: original},
	}
	return
}

func (s *Service) GetReceivers(c *gin.Context) ([]*oas.Receiver, error) {
	pr := s.Coordinator.GetReceivers()
	receivers := make([]*oas.Receiver, 0, len(pr))
	for i := range pr {
		receivers = append(receivers, &oas.Receiver{Name: pr[i].Name})
	}
	return receivers, nil
}

func (s *Service) GetAlerts(c *gin.Context, request *oas.GetAlertsRequest) (res oas.GettableAlerts, err error) {
	var (
		receiverFilter *regexp.Regexp
	)

	matchers, err := parseFilter(request.Body.Filter)
	if err != nil {
		return nil, err
	}

	if request.Body.Receiver != "" {
		receiverFilter, err = regexp.Compile("^(?:" + request.Body.Receiver + ")$")
		if err != nil {
			return nil, err
		}
	}

	alerts := s.Alerts.GetPending()
	defer alerts.Close()

	alertFilter := s.alertFilter(matchers, request.Body.Silenced, request.Body.Inhibited, request.Body.Active)
	now := time.Now()

	s.mtx.RLock()
	for a := range alerts.Next() {
		if err = alerts.Err(); err != nil {
			break
		}

		routes := s.route.Match(a.Labels)
		receivers := make([]string, 0, len(routes))
		for _, r := range routes {
			receivers = append(receivers, r.RouteOpts.Receiver)
		}

		if receiverFilter != nil && !receiversMatchFilter(receivers, receiverFilter) {
			continue
		}

		if !alertFilter(a, now) {
			continue
		}

		alert := AlertToOpenAPIAlert(a, s.StatusFunc(a.Fingerprint()), receivers)

		res = append(res, alert)
	}
	s.mtx.RUnlock()

	if err != nil {
		return
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Fingerprint < res[j].Fingerprint
	})
	return
}

func (s *Service) alertFilter(matchers []*label.Matcher, silenced, inhibited, active bool) func(a *alert.Alert, now time.Time) bool {
	return func(a *alert.Alert, now time.Time) bool {
		if !a.EndsAt.IsZero() && a.EndsAt.Before(now) {
			return false
		}

		// Set alert's current status based on its label set.
		s.setAlertStatus(a.Labels)

		// Get alert's current status after seeing if it is suppressed.
		status := s.StatusFunc(a.Fingerprint())

		if !active && status.State == alert.AlertStateActive {
			return false
		}

		if !silenced && len(status.SilencedBy) != 0 {
			return false
		}

		if !inhibited && len(status.InhibitedBy) != 0 {
			return false
		}

		return alertMatchesFilterLabels(a, matchers)
	}
}

func parseFilter(filter []string) ([]*label.Matcher, error) {
	matchers := make([]*label.Matcher, 0, len(filter))
	for _, matcherString := range filter {
		if matcherString == "" {
			continue
		}
		matcher, err := label.ParseMatcher(matcherString)
		if err != nil {
			return nil, err
		}

		matchers = append(matchers, matcher)
	}
	return matchers, nil
}

func alertMatchesFilterLabels(a *alert.Alert, matchers []*label.Matcher) bool {
	sms := make(map[string]string)
	for name, value := range a.Labels {
		sms[string(name)] = string(value)
	}
	return matchFilterLabels(matchers, sms)
}

func matchFilterLabels(matchers []*label.Matcher, sms map[string]string) bool {
	for _, m := range matchers {
		v, prs := sms[m.Name]
		switch m.Type {
		case label.MatchNotRegexp, label.MatchNotEqual:
			if m.Value == "" && prs {
				continue
			}
			if !m.Matches(v) {
				return false
			}
		default:
			if m.Value == "" && !prs {
				continue
			}
			if !m.Matches(v) {
				return false
			}
		}
	}

	return true
}

func receiversMatchFilter(receivers []string, filter *regexp.Regexp) bool {
	for _, r := range receivers {
		if filter.MatchString(r) {
			return true
		}
	}

	return false
}

func (s *Service) PostAlerts(c *gin.Context, req *oas.PostAlertsRequest) error {
	alerts := OpenAPIAlertsToAlerts(req.PostableAlerts)
	now := time.Now()

	resolveTimeout := s.Coordinator.ResolveTimeout()

	for _, alert := range alerts {
		alert.UpdatedAt = now

		// Ensure StartsAt is set.
		if alert.StartsAt.IsZero() {
			if alert.EndsAt.IsZero() {
				alert.StartsAt = now
			} else {
				alert.StartsAt = alert.EndsAt
			}
		}
		// If no end time is defined, set a timeout after which an alert
		// is marked resolved if it is not updated.
		if alert.EndsAt.IsZero() {
			alert.Timeout = true
			alert.EndsAt = now.Add(resolveTimeout)
		}
		if alert.EndsAt.After(time.Now()) {
			s.m.Firing().Inc()
		} else {
			s.m.Resolved().Inc()
		}
	}

	// Make a best effort to insert all alerts that are valid.
	var (
		validAlerts = make([]*alert.Alert, 0, len(alerts))
	)
	for _, a := range alerts {
		removeEmptyLabels(a.Labels)
		if err := a.Validate(); err != nil {
			c.Error(err)
			s.m.Invalid().Inc()
			continue
		}
		validAlerts = append(validAlerts, a)
	}
	if err := s.Alerts.Put(validAlerts...); err != nil {
		return err
	}

	return nil
}

func removeEmptyLabels(ls label.LabelSet) {
	for k, v := range ls {
		if string(v) == "" {
			delete(ls, k)
		}
	}
}

func (s *Service) DeleteSilence(c *gin.Context, req *oas.DeleteSilenceRequest) error {
	if err := s.Silences.Expire(req.UriParams.SilenceID); err != nil {
		if errors.Is(err, silence.ErrNotFound) {
			c.Status(http.StatusNotFound)
			return nil
		}
		return err
	}
	return nil
}

func (s *Service) GetSilence(c *gin.Context, req *oas.GetSilenceRequest) (*oas.GettableSilence, error) {
	sils, _, err := s.Silences.Query(silence.QIDs([]int{req.UriParams.SilenceID}))
	if err != nil {
		return nil, err
	}
	if len(sils) == 0 {
		c.Status(http.StatusNotFound)
		return nil, nil
	}

	sil, err := GettableSilenceFromProto(sils[0])
	if err != nil {
		return nil, err
	}
	return sil, nil
}

func (s *Service) GetSilences(c *gin.Context, req *oas.GetSilencesRequest) (vals oas.GettableSilences, err error) {
	var matchers []*label.Matcher
	for _, matcherStr := range req.Filter {
		matcher, err := label.ParseMatcher(matcherStr)
		if err != nil {
			return nil, err
		}
		matchers = append(matchers, matcher)
	}
	psils, _, err := s.Silences.Query()
	if err != nil {
		return nil, err
	}
	vals = make([]*oas.GettableSilence, 0, len(psils))
	for _, sil := range psils {
		if !CheckSilenceMatchesFilterLabels(sil, matchers) {
			continue
		}
		val, err := GettableSilenceFromProto(sil)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	SortSilences(vals)
	return
}

func (s *Service) PostSilences(c *gin.Context, req *oas.PostSilencesRequest) (res *oas.PostSilencesResponse, err error) {
	sil, err := PostableSilenceToEnt(&req.PostableSilence)
	if err != nil {
		return nil, err
	}

	sid, err := s.Silences.Set(&silence.Entry{
		ID:        sil.ID,
		UpdatedAt: sil.UpdatedAt,
		Matchers:  sil.Matchers,
		StartsAt:  sil.StartsAt,
		EndsAt:    sil.EndsAt,
		State:     sil.State,
	})
	if err != nil {
		if errors.Is(err, silence.ErrNotFound) {
			return nil, err
		}
		handler.AbortWithError(c, http.StatusInternalServerError, err)
		return nil, err
	}

	res = &oas.PostSilencesResponse{
		SilenceID: sid,
	}
	return
}

// AlertToOpenAPIAlert converts internal alerts, alert types, and receivers to *open_api_models.GettableAlert.
func AlertToOpenAPIAlert(alert *alert.Alert, status alert.MarkerStatus, receivers []string) *oas.GettableAlert {
	apiReceivers := make([]*oas.Receiver, 0, len(receivers))
	for i := range receivers {
		apiReceivers = append(apiReceivers, &oas.Receiver{Name: receivers[i]})
	}

	aa := &oas.GettableAlert{
		Alert: &oas.Alert{
			GeneratorURL: alert.GeneratorURL,
			Labels:       ModelLabelSetToAPILabelSet(alert.Labels),
		},
		Annotations: ModelLabelSetToAPILabelSet(alert.Annotations),
		StartsAt:    alert.StartsAt,
		UpdatedAt:   alert.UpdatedAt,
		EndsAt:      alert.EndsAt,
		Fingerprint: alert.Fingerprint().String(),
		Receivers:   apiReceivers,
		Status: oas.AlertStatus{
			State:       string(status.State),
			SilencedBy:  status.SilencedBy,
			InhibitedBy: status.InhibitedBy,
		},
	}
	return aa
}

// OpenAPIAlertsToAlerts converts open_api_models.PostableAlerts to []*types.Alert.
func OpenAPIAlertsToAlerts(apiAlerts oas.PostableAlerts) []*alert.Alert {
	alerts := []*alert.Alert{}
	for _, apiAlert := range apiAlerts {
		a := alert.Alert{
			Labels:       APILabelSetToModelLabelSet(apiAlert.Labels),
			Annotations:  APILabelSetToModelLabelSet(apiAlert.Annotations),
			StartsAt:     apiAlert.StartsAt,
			EndsAt:       apiAlert.EndsAt,
			GeneratorURL: apiAlert.GeneratorURL,
		}
		alerts = append(alerts, &a)
	}

	return alerts
}

// ModelLabelSetToAPILabelSet converts prometheus_model.LabelSet to open_api_models.LabelSet.
func ModelLabelSetToAPILabelSet(modelLabelSet label.LabelSet) oas.LabelSet {
	apiLabelSet := oas.LabelSet{}
	for key, value := range modelLabelSet {
		apiLabelSet[string(key)] = value
	}

	return apiLabelSet
}

// APILabelSetToModelLabelSet converts open_api_models.LabelSet to prometheus_model.LabelSet.
func APILabelSetToModelLabelSet(apiLabelSet oas.LabelSet) label.LabelSet {
	modelLabelSet := label.LabelSet{}
	for key, value := range apiLabelSet {
		modelLabelSet[label.LabelName(key)] = value
	}

	return modelLabelSet
}

// PostableSilenceToEnt converts *open_api_models.PostableSilenc to *silencepb.Silence.
func PostableSilenceToEnt(s *oas.PostableSilence) (*ent.Silence, error) {
	sil := &ent.Silence{
		ID:        s.ID,
		StartsAt:  s.StartsAt,
		EndsAt:    s.EndsAt,
		Comments:  s.Comment,
		CreatedBy: s.CreatedBy,
		TenantID:  s.TenantID,
	}
	for _, m := range s.Matchers {
		matcher, err := label.NewMatcher(label.MatchEqual, m.Name, m.Value)
		if err != nil {
			return nil, err
		}
		sil.Matchers = append(sil.Matchers, matcher)
	}
	return sil, nil
}

// GettableSilenceFromProto converts *silencepb.Silence to open_api_models.GettableSilence.
func GettableSilenceFromProto(s *silence.Entry) (*oas.GettableSilence, error) {
	start := s.StartsAt
	end := s.EndsAt
	updated := s.UpdatedAt
	state := string(alert.CalcSilenceState(start, end))
	sil := &oas.GettableSilence{
		Silence: &oas.Silence{
			StartsAt: start,
			EndsAt:   end,
			//Comment:   s.Comments,
			//CreatedBy: s.CreatedBy,
		},
		ID:        s.ID,
		UpdatedAt: updated,
		Status: oas.SilenceStatus{
			State: state,
		},
	}

	for _, m := range s.Matchers {
		matcher := &oas.Matcher{
			Name:  m.Name,
			Value: m.Value,
		}
		switch m.Type {
		case label.MatchEqual:
			matcher.IsEqual = true
			matcher.IsRegex = false
		case label.MatchNotEqual:
			matcher.IsEqual = false
			matcher.IsRegex = true
		case label.MatchRegexp:
			matcher.IsEqual = true
			matcher.IsRegex = true
		case label.MatchNotRegexp:
			matcher.IsEqual = false
			matcher.IsRegex = true
		default:
			return sil, fmt.Errorf(
				"unknown matcher type for matcher '%v' in silence '%v'",
				m.Name,
				s.ID,
			)
		}
		sil.Matchers = append(sil.Matchers, matcher)
	}

	return sil, nil
}

// CheckSilenceMatchesFilterLabels returns true if
// a given silence matches a list of matchers.
// A silence matches a filter (list of matchers) if
// for all matchers in the filter, there exists a matcher in the silence
// such that their names, types, and values are equivalent.
func CheckSilenceMatchesFilterLabels(s *silence.Entry, matchers []*label.Matcher) bool {
	for _, matcher := range matchers {
		found := false
		for _, m := range s.Matchers {
			if matcher.Name == m.Name &&
				(matcher.Type == label.MatchEqual && m.Type == label.MatchEqual ||
					matcher.Type == label.MatchRegexp && m.Type == label.MatchRegexp ||
					matcher.Type == label.MatchNotEqual && m.Type == label.MatchNotEqual ||
					matcher.Type == label.MatchNotRegexp && m.Type == label.MatchNotRegexp) &&
				matcher.Value == m.Value {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

var silenceStateOrder = map[alert.SilenceState]int{
	alert.SilenceStateActive:  1,
	alert.SilenceStatePending: 2,
	alert.SilenceStateExpired: 3,
}

// SortSilences sorts first according to the state "active, pending, expired"
// then by end time or start time depending on the state.
// active silences should show the next to expire first
// pending silences are ordered based on which one starts next
// expired are ordered based on which one expired most recently
func SortSilences(sils oas.GettableSilences) {
	sort.Slice(sils, func(i, j int) bool {
		state1 := alert.SilenceState(sils[i].Status.State)
		state2 := alert.SilenceState(sils[j].Status.State)
		if state1 != state2 {
			return silenceStateOrder[state1] < silenceStateOrder[state2]
		}
		switch state1 {
		case alert.SilenceStateActive:
			endsAt1 := sils[i].Silence.EndsAt
			endsAt2 := sils[j].Silence.EndsAt
			return endsAt1.Before(endsAt2)
		case alert.SilenceStatePending:
			startsAt1 := sils[i].Silence.StartsAt
			startsAt2 := sils[j].Silence.StartsAt
			return startsAt1.Before(startsAt2)
		case alert.SilenceStateExpired:
			endsAt1 := sils[i].Silence.EndsAt
			endsAt2 := sils[j].Silence.EndsAt
			return endsAt1.After(endsAt2)
		}
		return false
	})
}

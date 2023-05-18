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
	"github.com/woocoos/msgcenter/metrics"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/silence"
	"github.com/woocoos/msgcenter/silence/silencepb"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	logger log.ComponentLogger
	m      *metrics.Alerts
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
		m:       metrics.NewAlerts("v2", opts.Registry),
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
	alerts := OpenAPIAlertsToAlerts(req.Body)
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

func (s *Service) DeleteSilence(context *gin.Context, request *oas.DeleteSilenceRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetSilence(context *gin.Context, request *oas.GetSilenceRequest) (*oas.GettableSilence, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetSilences(context *gin.Context, request *oas.GetSilencesRequest) (oas.GettableSilences, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) PostSilences(c *gin.Context, req *oas.PostSilencesRequest) (res *oas.PostSilencesResponse, err error) {
	sil, err := PostableSilenceToProto(&req.Body)
	if err != nil {
		return nil, err
	}

	sid, err := s.Silences.Set(sil)
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

	if aa.Status.SilencedBy == nil {
		aa.Status.SilencedBy = []string{}
	}

	if aa.Status.InhibitedBy == nil {
		aa.Status.InhibitedBy = []string{}
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

// PostableSilenceToProto converts *open_api_models.PostableSilenc to *silencepb.Silence.
func PostableSilenceToProto(s *oas.PostableSilence) (*silencepb.Silence, error) {
	sil := &silencepb.Silence{
		Id:        s.ID,
		StartsAt:  timestamppb.New(s.StartsAt),
		EndsAt:    timestamppb.New(s.EndsAt),
		Comment:   s.Comment,
		CreatedBy: s.CreatedBy,
	}
	for _, m := range s.Matchers {
		matcher := &silencepb.Matcher{
			Name:    m.Name,
			Pattern: m.Value,
		}
		isEqual := m.IsEqual
		isRegex := m.IsRegex

		switch {
		case isEqual && !isRegex:
			matcher.Type = silencepb.Matcher_EQUAL
		case !isEqual && !isRegex:
			matcher.Type = silencepb.Matcher_NOT_EQUAL
		case isEqual && isRegex:
			matcher.Type = silencepb.Matcher_REGEXP
		case !isEqual && isRegex:
			matcher.Type = silencepb.Matcher_NOT_REGEXP
		}
		sil.Matchers = append(sil.Matchers, matcher)
	}
	return sil, nil
}

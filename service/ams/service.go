package ams

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/knockout-go/pkg/identity"
	"github.com/woocoos/msgcenter/api/graphql/model"
	"github.com/woocoos/msgcenter/dispatch"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msgalert"
	"github.com/woocoos/msgcenter/ent/msgchannel"
	"github.com/woocoos/msgcenter/ent/msgevent"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/ent/nlog"
	"github.com/woocoos/msgcenter/ent/org"
	"github.com/woocoos/msgcenter/ent/predicate"
	"github.com/woocoos/msgcenter/ent/user"
	"github.com/woocoos/msgcenter/ent/useraddr"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/service"
	"go.uber.org/zap"
)

var logger = log.Component("ams")

type Option func(*Service)

type Service struct {
	client *ent.Client
	am     *service.AlertManager
}

func NewService(opt ...Option) *Service {
	r := &Service{}
	for _, o := range opt {
		o(r)
	}
	return r
}

func WithClient(client *ent.Client) Option {
	return func(s *Service) {
		s.client = client
	}
}

func WithAlertManager(am *service.AlertManager) Option {
	return func(r *Service) {
		r.am = am
	}
}

// formatContext 持有批量预查询的数据，避免重复查询
type formatContext struct {
	channelComments map[string]string // receiver -> comments
	eventComments   map[string]string // alertName -> comments
	userMap         map[int]*model.UserInfo
}

// buildFormatContext 批量查询通道备注、事件备注、用户信息
func (s *Service) buildFormatContext(ctx context.Context, tid int, receivers []string, alertNames []string, userIDs []int) *formatContext {
	fc := &formatContext{
		channelComments: make(map[string]string),
		eventComments:   make(map[string]string),
		userMap:         make(map[int]*model.UserInfo),
	}
	// 批量查询通道备注
	if len(receivers) > 0 {
		channels, _ := s.client.MsgChannel.Query().Where(
			msgchannel.NameIn(receivers...), msgchannel.TenantID(tid),
		).All(ctx)
		for _, ch := range channels {
			if ch.Comments != "" {
				fc.channelComments[ch.Name] = ch.Comments
			}
		}
	}
	// 批量查询事件备注
	if len(alertNames) > 0 {
		events, _ := s.client.MsgEvent.Query().Where(
			msgevent.NameIn(alertNames...),
		).All(ctx)
		for _, e := range events {
			if e.Comments != "" {
				fc.eventComments[e.Name] = e.Comments
			}
		}
	}
	// 批量查询用户信息（含联系方式）
	if len(userIDs) > 0 {
		users, _ := s.client.User.Query().Where(user.IDIn(userIDs...)).
			WithAddresses(func(q *ent.UserAddrQuery) {
				q.Where(useraddr.AddrTypeEQ(useraddr.AddrTypeContact))
			}).
			All(ctx)
		for _, u := range users {
			uid := strconv.Itoa(u.ID)
			info := &model.UserInfo{
				Name:   &u.DisplayName,
				UserID: &uid,
			}
			if len(u.Edges.Addresses) > 0 {
				info.Email = &u.Edges.Addresses[0].Email
				info.Mobile = &u.Edges.Addresses[0].Mobile
			}
			fc.userMap[u.ID] = info
		}
	}
	return fc
}

func (s *Service) FormatMsgAlerts(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, alertName *string, userID *string, receiverType *profile.ReceiverType, orderBy *ent.MsgAlertOrder, where *ent.MsgAlertWhereInput) (*model.FormatMsgAlertConnection, error) {
	msgalert.LabelsIsNil()
	w := make([]predicate.MsgAlert, 0)
	if alertName != nil {
		an := func(s *sql.Selector) {
			s.Where(sqljson.ValueEQ(msgalert.FieldLabels, *alertName, sqljson.Path(label.AlertNameLabel)))
		}
		w = append(w, an)
	}
	if receiverType != nil {
		ry := func(s *sql.Selector) {
			s.Where(sqljson.ValueContains(msgalert.FieldLabels, receiverType.String(), sqljson.Path("receiver")))
		}
		w = append(w, msgalert.Or(ry, msgalert.HasNlogWith(nlog.ReceiverTypeEQ(*receiverType))))
	}
	if userID != nil {
		usr := func(s *sql.Selector) {
			s.Where(sqljson.ValueContains(msgalert.FieldLabels, userID, sqljson.Path(label.ToUserIDLabel)))
		}
		w = append(w, usr)
	}
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	o, err := s.client.Org.Get(ctx, tid)
	if err != nil {
		return nil, err
	}
	// 查询所有子租户
	w = append(w, msgalert.HasOrgWith(org.Or(org.PathContains(o.Path), org.Path(o.Path))))
	msgAlerts, err := s.client.MsgAlert.Query().Where(w...).Paginate(schemax.SkipTenantPrivacy(ctx), after, first, before, last,
		ent.WithMsgAlertOrder(orderBy), ent.WithMsgAlertFilter(where.Filter))
	if err != nil {
		return nil, err
	}
	// 第一遍：匹配路由，收集 receivers、alertNames、userIDs 去重
	type alertRoute struct {
		edge      *ent.MsgAlertEdge
		routes    []*dispatch.Route
		receivers []string
		alertName string
		userIDs   []int
	}
	items := make([]alertRoute, 0, len(msgAlerts.Edges))
	receiverSet := make(map[string]struct{})
	alertNameSet := make(map[string]struct{})
	userIDSet := make(map[int]struct{})
	for _, edge := range msgAlerts.Edges {
		if edge.Node.Labels == nil {
			continue
		}
		labels := *edge.Node.Labels
		rs := s.am.Route.Match(labels)
		if len(rs) == 0 {
			continue
		}
		receivers := make([]string, len(rs))
		for i, r := range rs {
			receivers[i] = r.RouteOpts.Receiver
			receiverSet[receivers[i]] = struct{}{}
		}
		alertName := labels[label.LabelName(label.AlertNameLabel)]
		alertNameSet[alertName] = struct{}{}
		uids, _ := label.UserIDsFromLabels(labels)
		item := alertRoute{edge: edge, routes: rs, receivers: receivers, alertName: alertName, userIDs: uids}
		for _, uid := range uids {
			userIDSet[uid] = struct{}{}
		}
		items = append(items, item)
	}
	// 批量查询
	receivers := make([]string, 0, len(receiverSet))
	for r := range receiverSet {
		receivers = append(receivers, r)
	}
	alertNames := make([]string, 0, len(alertNameSet))
	for n := range alertNameSet {
		alertNames = append(alertNames, n)
	}
	allUIDs := make([]int, 0, len(userIDSet))
	for uid := range userIDSet {
		allUIDs = append(allUIDs, uid)
	}
	fc := s.buildFormatContext(ctx, tid, receivers, alertNames, allUIDs)
	// 第二遍：构建结果
	edges := make([]*model.FormatMsgAlertEdge, 0, len(items))
	for _, item := range items {
		fa, err := s.renderMsgAlert(ctx, item.edge.Node, item.routes[0], fc, item.userIDs)
		if err != nil {
			return nil, err
		}
		if fa == nil {
			continue
		}
		fa.HasMultiMsg = len(item.routes) > 1
		fa.Modes = ptr(strings.Join(item.receivers, ","))
		// 聚合所有路由的通道备注
		if comments := collectComments(item.receivers, fc.channelComments); comments != "" {
			fa.MsgChannelComments = ptr(comments)
		}
		edges = append(edges, &model.FormatMsgAlertEdge{Cursor: item.edge.Cursor, Node: fa})
	}
	return &model.FormatMsgAlertConnection{
		Edges:      edges,
		PageInfo:   &msgAlerts.PageInfo,
		TotalCount: msgAlerts.TotalCount,
	}, nil
}

func (s *Service) renderMsgAlert(ctx context.Context, msgAlert *ent.MsgAlert, route *dispatch.Route, fc *formatContext, userIDs []int) (*model.FormatMsgAlert, error) {
	if msgAlert.Labels == nil {
		return nil, nil
	}
	labels := *msgAlert.Labels
	msgTemplateTitle := ""
	a := s.convertMsgAlert(msgAlert)
	// 获取模板信息
	routeOpt := route.RouteOpts
	msgTemp, err := s.findMsgTemplate(ctx, routeOpt.Receiver, a)
	if err != nil {
		logger.Error("not find msg template", zap.Error(err), zap.Int("alertID", msgAlert.ID))
		return nil, nil
	}
	// 模板标题
	if msgTemp != nil {
		data := s.am.Coordinator.Template.Data("", nil, []*alert.Alert{&a}...)
		msgTemplateTitle, err = s.am.Coordinator.Template.ExecuteHTMLString(msgTemp.Subject, data)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	// 从预查询的 map 中获取备注
	alertName := labels[label.LabelName(label.AlertNameLabel)]
	msgEventComments := fc.eventComments[alertName]
	msgChannelComments := fc.channelComments[routeOpt.Receiver]
	// 从预查询的 userMap 中获取用户信息
	users := make([]*model.UserInfo, 0, len(userIDs))
	for _, uid := range userIDs {
		if info, ok := fc.userMap[uid]; ok {
			users = append(users, info)
		}
	}
	if len(users) == 0 {
		// 如果没有用户，则返回消息体的邮箱
		annotations := *msgAlert.Annotations
		to := annotations["to"]
		if to != "" {
			users = append(users, &model.UserInfo{
				Email: &to,
			})
		}
	}
	return &model.FormatMsgAlert{
		ID:                 msgAlert.ID,
		TenantID:           msgAlert.TenantID,
		StartsAt:           msgAlert.StartsAt,
		EndsAt:             &msgAlert.EndsAt,
		State:              msgAlert.State,
		ReceiverType:       msgTemp.ReceiverType,
		Receiver:           routeOpt.Receiver,
		MsgEventComments:   ptr(msgEventComments),
		MsgChannelComments: ptr(msgChannelComments),
		MsgTemplateTitle:   &msgTemplateTitle,
		Users:              users,
	}, nil
}

func (s *Service) convertMsgAlert(msgAlert *ent.MsgAlert) alert.Alert {
	return alert.Alert{
		Labels:       *msgAlert.Labels,
		Annotations:  *msgAlert.Annotations,
		StartsAt:     msgAlert.StartsAt,
		EndsAt:       msgAlert.EndsAt,
		GeneratorURL: msgAlert.URL,
		Timeout:      msgAlert.Timeout,
	}
}

func (s *Service) findMsgTemplate(ctx context.Context, receiver string, a alert.Alert) (*ent.MsgTemplate, error) {
	var msgTemp *ent.MsgTemplate
	var err error
	var rt profile.ReceiverType
	if strings.HasPrefix(receiver, profile.ReceiverWebhook.String()) {
		// webhook
		rt = profile.ReceiverWebhook
	} else if strings.HasPrefix(receiver, profile.ReceiverEmail.String()) {
		// email
		rt = profile.ReceiverEmail
	} else if strings.HasPrefix(receiver, profile.ReceiverMessage.String()) {
		// message
		rt = profile.ReceiverMessage
	} else if strings.HasPrefix(receiver, profile.ReceiverUmeng.String()) {
		// umeng
		rt = profile.ReceiverUmeng
	} else {
		// unknown
		return nil, fmt.Errorf("unknown receiver")
	}
	msgTemp, err = s.am.Coordinator.FindTemplate(ctx, s.client, rt, a.Labels)
	if err != nil {
		return nil, err
	}
	return msgTemp, nil
}

func (s *Service) FormatMsgAlertMore(ctx context.Context, msgAlertID int) ([]*model.FormatMsgAlert, error) {
	ma, err := s.client.MsgAlert.Get(schemax.SkipTenantPrivacy(ctx), msgAlertID)
	if err != nil {
		return nil, err
	}
	if ma.Labels == nil {
		return nil, nil
	}
	labels := *ma.Labels
	alertName := labels[label.LabelName(label.AlertNameLabel)]
	// 获取路由
	rs := s.am.Route.Match(labels)
	receivers := make([]string, len(rs))
	for i, r := range rs {
		receivers[i] = r.RouteOpts.Receiver
	}
	userIDs, _ := label.UserIDsFromLabels(labels)
	tid, _ := identity.TenantIDFromContext(ctx)
	// 批量查询
	fc := s.buildFormatContext(ctx, tid, receivers, []string{alertName}, userIDs)
	msgAlerts := make([]*model.FormatMsgAlert, 0, len(rs))
	for _, route := range rs {
		fa, err := s.renderMsgAlert(ctx, ma, route, fc, userIDs)
		if err != nil {
			return nil, err
		}
		if fa == nil {
			continue
		}
		msgAlerts = append(msgAlerts, fa)
	}
	return msgAlerts, nil
}

func (s *Service) RenderMsgAlert(ctx context.Context, msgAlertID int, receiver string) (*string, error) {
	msgAlert, err := s.client.MsgAlert.Query().Where(msgalert.ID(msgAlertID)).Only(schemax.SkipTenantPrivacy(ctx))
	if err != nil {
		return nil, err
	}
	a := s.convertMsgAlert(msgAlert)
	msgTemp, err := s.findMsgTemplate(ctx, receiver, a)
	if err != nil {
		return nil, err
	}
	tplStr := ""
	data := notify.GetTemplateData(ctx, s.am.Coordinator.Template, []*alert.Alert{&a})
	if msgTemp.Format == msgtemplate.FormatHTML {
		tplStr, err = s.am.Coordinator.Template.ExecuteHTMLString(msgTemp.Body, data)
	} else if msgTemp.Format == msgtemplate.FormatTxt {
		tplStr, err = s.am.Coordinator.Template.ExecuteTextString(msgTemp.Body, data)
	} else {
		return nil, fmt.Errorf("unknown format: %s", msgTemp.Format)
	}
	return &tplStr, nil
}

func ptr(s string) *string { return &s }

func collectComments(receivers []string, m map[string]string) string {
	comments := make([]string, 0, len(receivers))
	for _, r := range receivers {
		if c, ok := m[r]; ok {
			comments = append(comments, c)
		}
	}
	return strings.Join(comments, ",")
}

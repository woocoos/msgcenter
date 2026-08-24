package ams

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
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
)

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
	formatMsgAlerts := make([]*model.FormatMsgAlertEdge, 0)
	// 遍历消息列表，转换成格式化消息列表
	for _, msgAlert := range msgAlerts.Edges {
		if msgAlert.Node.Labels == nil {
			continue
		}
		labels := *msgAlert.Node.Labels
		// 获取路由
		rs := s.am.Route.Match(labels)
		if len(rs) == 0 {
			continue
		}
		hasMultiMsg := false
		if len(rs) > 1 {
			hasMultiMsg = true
		}
		route := rs[0]
		formatMsgAlert, err := s.formatMsgAlert(ctx, msgAlert.Node, route)
		if err != nil {
			return nil, err
		}
		if formatMsgAlert != nil {
			formatMsgAlert.HasMultiMsg = hasMultiMsg
		}
		formatMsgAlerts = append(formatMsgAlerts, &model.FormatMsgAlertEdge{
			Cursor: msgAlert.Cursor,
			Node:   formatMsgAlert,
		})
	}
	return &model.FormatMsgAlertConnection{
		Edges:      formatMsgAlerts,
		PageInfo:   &msgAlerts.PageInfo,
		TotalCount: msgAlerts.TotalCount,
	}, nil
}

func (s *Service) formatMsgAlert(ctx context.Context, msgAlert *ent.MsgAlert, route *dispatch.Route) (*model.FormatMsgAlert, error) {
	if msgAlert.Labels == nil {
		return nil, nil
	}
	labels := *msgAlert.Labels
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	msgChannelComments := ""
	msgTemplateTitle := ""
	a := s.convertMsgAlert(msgAlert)
	// 获取模板信息
	routeOpt := route.RouteOpts
	msgTemp, err := s.findMsgTemplate(ctx, routeOpt.Receiver, a)
	if err != nil {
		return nil, err
	}
	// 模板标题
	if msgTemp != nil {
		data := notify.GetTemplateData(ctx, s.am.Coordinator.Template, []*alert.Alert{&a})
		msgTemplateTitle, err = s.am.Coordinator.Template.ExecuteHTMLString(msgTemp.Subject, data)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	// 获取消息通道描述
	msgChannel, err := s.client.MsgChannel.Query().Where(
		msgchannel.Name(routeOpt.Receiver), msgchannel.TenantID(tid),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	if msgChannel != nil {
		msgChannelComments = msgChannel.Comments
	}
	// 消息事件
	alertName := labels[label.LabelName(label.AlertNameLabel)]
	msgEvent, err := s.client.MsgEvent.Query().Where(
		msgevent.Name(alertName),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	msgEventComments := ""
	if msgEvent != nil {
		msgEventComments = msgEvent.Comments
	}
	// 判断消息是否订阅
	users := make([]*model.UserInfo, 0)
	// 取消息体的user
	uids, err := service.UserIDsFromLabels(labels)
	if err != nil {
		return nil, err
	}
	if uids != nil {
		us, err := s.client.User.Query().Where(user.IDIn(uids...)).All(ctx)
		if err != nil {
			return nil, err
		}
		for _, u := range us {
			addr, err := u.QueryAddresses().Where(useraddr.AddrTypeEQ(useraddr.AddrTypeContact)).Only(ctx)
			if err != nil {
				return nil, err
			}
			uid := strconv.Itoa(u.ID)
			users = append(users, &model.UserInfo{
				Name:   &u.DisplayName,
				Email:  &addr.Email,
				Mobile: &addr.Mobile,
				UserID: &uid,
			})
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
		MsgEventComments:   &msgEventComments,
		MsgChannelComments: &msgChannelComments,
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
	if strings.HasPrefix(receiver, profile.ReceiverWebhook.String()) {
		// webhook
		msgTemp, err = s.am.Coordinator.FindTemplate(ctx, s.client, profile.ReceiverWebhook, a.Labels)
	} else if strings.HasPrefix(receiver, profile.ReceiverEmail.String()) {
		// email
		msgTemp, err = s.am.Coordinator.FindTemplate(ctx, s.client, profile.ReceiverEmail, a.Labels)
	} else if strings.HasPrefix(receiver, profile.ReceiverMessage.String()) {
		// message
		msgTemp, err = s.am.Coordinator.FindTemplate(ctx, s.client, profile.ReceiverMessage, a.Labels)
	} else {
		// unknown
		return nil, fmt.Errorf("unknown receiver")
	}
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
	msgAlerts := make([]*model.FormatMsgAlert, 0)
	// 获取路由
	rs := s.am.Route.Match(labels)
	for _, route := range rs {
		formatMsgAlert, err := s.formatMsgAlert(ctx, ma, route)
		if err != nil {
			return nil, err
		}
		msgAlerts = append(msgAlerts, formatMsgAlert)
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

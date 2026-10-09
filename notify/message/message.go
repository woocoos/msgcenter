package message

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
	"github.com/tsingsun/woocoo/pkg/log"
	ecx "github.com/woocoos/knockout-go/ent/clientx"
	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/pkg/push"
	"github.com/woocoos/msgcenter/template"
	"go.uber.org/zap"
)

var (
	logger = log.Component("message")
)

// Notifier is internal message notifier
type Notifier struct {
	config        *profile.MessageConfig
	tmpl          *template.Template
	customTplFunc notify.CustomerConfigFunc[profile.MessageConfig]
	client        *ent.Client
	rdb           redis.UniversalClient
}

func New(cfg *profile.MessageConfig, tmpl *template.Template,
	client *ent.Client, rdb redis.UniversalClient,
	fn notify.CustomerConfigFunc[profile.MessageConfig]) (*Notifier, error) {
	return &Notifier{
		config:        cfg,
		tmpl:          tmpl,
		customTplFunc: fn,
		client:        client,
		rdb:           rdb,
	}, nil
}

func (n *Notifier) SendResolved() bool {
	return false
}

// CustomConfig returns a custom config for the notifier.
func (n *Notifier) CustomConfig(ctx context.Context) (*profile.MessageConfig, error) {
	if n.customTplFunc == nil {
		return n.config, nil
	}
	labels, ok := notify.GroupLabels(ctx)
	if !ok {
		return n.config, nil
	}
	cfg := n.config.Clone()
	err := n.customTplFunc(ctx, cfg, labels)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// Notify implements the Notifier interface.
//
// Notice: the caller must ensure that the tenant id and user id are valid.
func (n *Notifier) Notify(ctx context.Context, alerts ...*alert.Alert) (retry bool, err error) {
	ts, _ := notify.Tenant(ctx)
	tid, err := strconv.Atoi(ts)
	if err != nil {
		return false, err
	}
	data := notify.GetTemplateData(ctx, n.tmpl, alerts)
	tmpl := notify.TmplText(n.tmpl, data, &err)

	config, err := n.CustomConfig(ctx)
	if err != nil {
		return false, err
	}
	if config.To == "" {
		return false, errors.New("to is empty")
	}

	// 检查是否有 alert 带有 skipStore label
	skipStore := false
	for _, a := range alerts {
		if _, ok := a.Labels[label.SkipStoreLabel]; ok {
			skipStore = true
			break
		}
	}

	var pushData = push.Data{
		Topic: "message",
	}

	// 提前渲染模板（skipStore 和非 skipStore 共用）
	var title, body, redirect string
	var formatStr string
	if config.Subject != "" {
		title = tmpl(config.Subject)
		if err != nil {
			return false, fmt.Errorf("execute 'Title' template: %w", err)
		}
	}
	if config.Text != "" {
		body = tmpl(config.Text)
		if err != nil {
			return false, fmt.Errorf("execute 'context' template: %w", err)
		}
		formatStr = "text"
	} else if config.HTML != "" {
		body = tmpl(config.HTML)
		if err != nil {
			return false, fmt.Errorf("execute 'context' template: %w", err)
		}
		formatStr = "html"
	}
	if config.Redirect != "" {
		redirect = tmpl(config.Redirect)
	}
	format := msgtemplate.Format(formatStr)

	alertNameStr := data.CommonLabels[label.AlertNameLabel]

	// db error ,don't try
	err = ecx.WithTx(ctx, func(ctx context.Context) (ecx.Transactor, error) {
		return n.client.Tx(ctx)
	}, func(itx ecx.Transactor) error {
		tx := itx.(*ent.Tx)
		nctx := schemax.SkipTenantPrivacy(ctx)

		var rowID int
		if !skipStore {
			msg := tx.MsgInternal.Create().
				SetCreatedBy(0).
				SetCategory(config.Extras["category"]).
				SetReceiverType(profile.ReceiverMessage).
				SetSubject(title).
				SetBody(body).
				SetFormat(formatStr).
				SetTenantID(tid)
			if idStr := data.CommonAnnotations[label.AlertIDAnnotation]; idStr != "" {
				if aid, err := strconv.Atoi(idStr); err == nil {
					msg.SetAlertID(aid)
				}
			}
			if redirect != "" {
				msg.SetRedirect(redirect)
			}
			row, err := msg.Save(nctx)
			if err != nil {
				return err
			}
			rowID = row.ID

			// 创建 MsgInternalTo 记录
			msggtos := make([]*ent.MsgInternalToCreate, 0)
			for _, uid := range strings.Split(config.To, ",") {
				suid, err := strconv.Atoi(uid)
				if err != nil {
					logger.Error("invalid user id", zap.String("userID", uid))
					continue
				}
				msggtos = append(msggtos,
					tx.MsgInternalTo.Create().SetTenantID(tid).SetUserID(suid).SetMsgInternalID(row.ID),
				)
				pushData.Audience.UserIDs = append(pushData.Audience.UserIDs, suid)
			}
			if len(msggtos) > 0 {
				_, err = tx.MsgInternalTo.CreateBulk(msggtos...).Save(nctx)
				if err != nil {
					return err
				}
			}
		} else {
			// skipStore 模式下只收集推送用户，不创建 DB 记录
			for _, uid := range strings.Split(config.To, ",") {
				suid, err := strconv.Atoi(uid)
				if err != nil {
					logger.Error("invalid user id", zap.String("userID", uid))
					continue
				}
				pushData.Audience.UserIDs = append(pushData.Audience.UserIDs, suid)
			}
		}

		// 构建推送消息（skipStore 和非 skipStore 共用）
		extras := map[label.LabelName]string{
			"action":    "internal",
			"alertName": alertNameStr,
		}
		if rowID > 0 {
			extras["actionID"] = strconv.Itoa(rowID)
		}
		pushData.Message = push.Message{
			Title:   title,
			Format:  format,
			Content: body,
			Extras:  extras,
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	n.notifyRedis(ctx, &pushData)
	return
}

// only log error
func (n *Notifier) notifyRedis(ctx context.Context, data *push.Data) {
	md, err := push.Marshal(data)
	if err != nil {
		log.Errorf("notifyRedis: marshal:%v", err)
	}
	if err := n.rdb.Publish(ctx, data.Topic, md).Err(); err != nil {
		log.Errorf("notifyRedis: publish redis:%v", err)
	}
}

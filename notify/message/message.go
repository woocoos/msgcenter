package message

import (
	"context"
	"errors"
	"fmt"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/template"
	"go.uber.org/zap"
	"strconv"
	"strings"
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
}

func New(cfg *profile.MessageConfig, tmpl *template.Template, client *ent.Client,
	fn notify.CustomerConfigFunc[profile.MessageConfig]) (*Notifier, error) {
	return &Notifier{
		config:        cfg,
		tmpl:          tmpl,
		customTplFunc: fn,
		client:        client,
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
	err = ecx.WithTx(ctx, func(ctx context.Context) (ecx.Transactor, error) {
		return n.client.Tx(ctx)
	}, func(itx ecx.Transactor) error {
		tx := itx.(*ent.Tx)
		msg := tx.MsgInternal.Create().SetCreatedBy(0)
		if config.Subject != "" {
			msg.SetSubject(tmpl(config.Subject))
			if err != nil {
				return fmt.Errorf("execute 'Title' template: %w", err)
			}
		}
		if config.Text != "" {
			msg.SetBody(tmpl(config.Text))
			if err != nil {
				return fmt.Errorf("execute 'context' template: %w", err)
			}
			msg.SetFormat("text")
		} else if config.HTML != "" {
			msg.SetBody(tmpl(config.HTML))
			if err != nil {
				return fmt.Errorf("execute 'context' template: %w", err)
			}
			msg.SetFormat("html")
		}
		if config.Redirect != "" {
			msg.SetRedirect(tmpl(config.Redirect))
		}
		msg.SetTenantID(tid)
		nctx := schemax.SkipTenantPrivacy(ctx)
		row, err := msg.Save(nctx)
		if err != nil {
			return err
		}
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
		}
		if len(msggtos) > 0 {
			_, err = tx.MsgInternalTo.CreateBulk(msggtos...).Save(nctx)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return
}

package service

import (
	"context"
	"errors"
	"net/mail"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msgevent"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/ent/msgtype"
	"github.com/woocoos/msgcenter/ent/user"
	"github.com/woocoos/msgcenter/ent/useraddr"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/service/kosdk"
)

var (
	ErrTenantIDNotFound = errors.New("tenant id not found")
)

func tenantIDFromLabels(set label.LabelSet) (int, error) {
	tv, ok := set[label.TenantLabel]
	if !ok {
		return 0, ErrTenantIDNotFound
	}
	tid, err := strconv.Atoi(tv)
	if err != nil {
		return 0, err
	}
	if tid == 0 {
		return 0, ErrTenantIDNotFound
	}
	return tid, nil
}

// UserIDsFromLabels returns the user IDs from the labels.
func UserIDsFromLabels(set label.LabelSet) ([]int, error) {
	ul, ok := set[label.ToUserIDLabel]
	if !ok {
		return nil, nil
	}
	ids := strings.Split(ul, ",")
	uis := make([]int, 0, len(ids))
	for _, id := range ids {
		uid, _ := strconv.Atoi(id)
		if uid == 0 {
			continue
		}
		uis = append(uis, uid)
	}
	return uis, nil
}

// findTemplate find template from database
// Priority: user-level template > tenant-level template > global default template
func findTemplate(ctx context.Context, basedir, attdir string, client *ent.Client, rt profile.ReceiverType,
	labels label.LabelSet) (*ent.MsgTemplate, error) {
	tid, err := tenantIDFromLabels(labels)
	if err != nil {
		if errors.Is(err, ErrTenantIDNotFound) {
			return nil, &ent.NotFoundError{}
		}
		return nil, err
	}
	en := labels[label.AlertNameLabel]

	// Try user-level template first (only if single user ID in labels)
	userIDs, _ := UserIDsFromLabels(labels)
	if len(userIDs) == 1 {
		uid := userIDs[0]
		event, err := client.MsgTemplate.Query().Where(
			msgtemplate.TenantID(tid),
			msgtemplate.UserID(uid),
			msgtemplate.StatusEQ(typex.SimpleStatusActive),
			msgtemplate.HasEventWith(msgevent.Name(en), msgevent.StatusEQ(typex.SimpleStatusActive)),
			msgtemplate.ReceiverTypeEQ(rt),
		).Only(ctx)
		if err == nil {
			return processTemplateAttachments(event, basedir, attdir)
		}
		if !ent.IsNotFound(err) {
			return nil, err
		}
		// User-level template not found, fall through to tenant-level
	}

	// Try tenant-level template
	event, err := client.MsgTemplate.Query().Where(
		msgtemplate.TenantID(tid),
		msgtemplate.UserIDIsNil(),
		msgtemplate.StatusEQ(typex.SimpleStatusActive),
		msgtemplate.HasEventWith(msgevent.Name(en), msgevent.StatusEQ(typex.SimpleStatusActive)),
		msgtemplate.ReceiverTypeEQ(rt),
	).Only(ctx)
	if err == nil {
		return processTemplateAttachments(event, basedir, attdir)
	}
	if !ent.IsNotFound(err) {
		return nil, err
	}

	// Fall back to global default template (tenant_id IS NULL)
	event, err = client.MsgTemplate.Query().Where(
		msgtemplate.TenantIDIsNil(),
		msgtemplate.UserIDIsNil(),
		msgtemplate.StatusEQ(typex.SimpleStatusActive),
		msgtemplate.HasEventWith(msgevent.Name(en), msgevent.StatusEQ(typex.SimpleStatusActive)),
		msgtemplate.ReceiverTypeEQ(rt),
	).Only(ctx)
	if err != nil {
		return nil, err
	}
	return processTemplateAttachments(event, basedir, attdir)
}

// processTemplateAttachments replaces remote attachment paths with local paths
func processTemplateAttachments(event *ent.MsgTemplate, basedir, attdir string) (*ent.MsgTemplate, error) {
	if event == nil {
		return nil, nil
	}
	if event.Attachments != nil && len(event.Attachments) > 0 {
		as := make([]string, len(event.Attachments))
		for i, attacher := range event.Attachments {
			if strings.Contains(attacher, "://") {
				path, err := kosdk.DefaultFilePath(event.TenantID, attacher, basedir, attdir)
				if err != nil {
					return nil, err
				}
				as[i] = path
			} else {
				// 按约定，默认模版附件只存储附件名称+后缀
				as[i] = filepath.Join(basedir, strconv.Itoa(event.TenantID), attdir, attacher)
			}
		}
		event.Attachments = as
	}
	return event, nil
}

func usersFromLabels(client *ent.Client, set label.LabelSet) ([]*ent.User, error) {
	ul, ok := set[label.ToUserIDLabel]
	if !ok {
		return nil, nil
	}
	ids := strings.Split(ul, ",")
	uis := make([]int, 0, len(ids))
	for _, id := range ids {
		uid, _ := strconv.Atoi(id)
		if uid == 0 {
			continue
		}
		uis = append(uis, uid)
	}
	return client.User.Query().Where(user.IDIn(uis...)).All(schemax.SkipTenantPrivacy(context.Background()))
}

// use for email.Email.customTplFunc.
// 1. Support load template from database
// 2. Get user info's email address if exist user id in label. The user info email replaces template to address.
func overrideEmailConfig(basedir, attdir string, client *ent.Client) notify.CustomerConfigFunc[profile.EmailConfig] {
	return func(ctx context.Context, cfg *profile.EmailConfig, set label.LabelSet,
	) error {
		data, err := findTemplate(ctx, basedir, attdir, client, profile.ReceiverEmail, set)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}
		eus, err := usersFromLabels(client, set)
		if err != nil {
			return err
		}
		if len(eus) > 0 {
			tos := make([]string, 0, len(eus))
			for _, u := range eus {
				addr, err := u.QueryAddresses().Where(useraddr.AddrTypeEQ(useraddr.AddrTypeContact)).Only(ctx)
				if err != nil {
					return err
				}
				if addr.Email != "" {
					ma := mail.Address{
						Name:    u.DisplayName,
						Address: addr.Email,
					}
					tos = append(tos, ma.String())
				}
			}
			cfg.To = strings.Join(tos, ",")
		} else {
			cfg.To = data.To
		}
		if data.From != "" {
			cfg.From = data.From
		}
		cfg.Subject = data.Subject
		if data.Format == msgtemplate.FormatHTML {
			cfg.HTML = data.Body
		} else {
			cfg.Text = data.Body
		}
		if data.Cc != "" {
			cfg.Headers["Cc"] = data.Cc
		}
		if data.Bcc != "" {
			cfg.Headers["Bcc"] = data.Bcc
		}
		if data.Attachments != nil && len(data.Attachments) > 0 {
			cfg.Headers["Attachments"] = strings.Join(data.Attachments, ",")
		}
		return nil
	}
}

func overrideWebHookConfig(basedir, attdir string, client *ent.Client) notify.CustomerConfigFunc[profile.WebhookConfig] {
	return func(ctx context.Context, cfg *profile.WebhookConfig, set label.LabelSet,
	) error {
		data, err := findTemplate(ctx, basedir, attdir, client, profile.ReceiverWebhook, set)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}
		cfg.Subject = data.Subject
		cfg.Body = data.Body
		return nil
	}
}

func overrideUmengConfig(basedir, attdir string, client *ent.Client) notify.CustomerConfigFunc[profile.UmengConfig] {
	return func(ctx context.Context, cfg *profile.UmengConfig, set label.LabelSet,
	) error {
		// Umeng config does not support template overrides from database.
		// Template rendering is handled in the notifier using default templates.
		return nil
	}
}

func overrideMessageConfig(basedir, attdir string, client *ent.Client) notify.CustomerConfigFunc[profile.MessageConfig] {
	return func(ctx context.Context, cfg *profile.MessageConfig, set label.LabelSet,
	) error {
		ul, ok := set[label.ToUserIDLabel]
		if !ok {
			return errors.New("user id is empty")
		}
		cfg.To = ul

		data, err := findTemplate(ctx, basedir, attdir, client, profile.ReceiverMessage, set)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}
		ev, err := data.QueryEvent().WithMsgType(func(query *ent.MsgTypeQuery) {
			query.Select(msgtype.FieldCategory)
		}).Only(ctx)
		if err != nil {
			// must have
			return err
		}
		mt, err := ev.MsgType(ctx)
		if err != nil {
			return err
		}
		cfg.Extras["category"] = mt.Category
		cfg.Subject = data.Subject
		if data.Format == msgtemplate.FormatHTML {
			cfg.HTML = data.Body
		} else {
			cfg.Text = data.Body
		}
		return nil
	}
}

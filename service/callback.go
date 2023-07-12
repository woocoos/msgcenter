package service

import (
	"context"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msgalert"
	"github.com/woocoos/msgcenter/nflog"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/provider/mem"
	"go.uber.org/zap"
	"strconv"
	"time"
)

var _ mem.AlertStoreCallback = (*AlertCallback)(nil)

// AlertCallback 是Alert处理回调类,由于消息中心是系统服务,对于alert无法有明显的上下文,因此需要在回调中处理租户时采用的忽略.
type AlertCallback struct {
	db *ent.Client
}

func (a *AlertCallback) PreStore(alert *alert.Alert, existing bool) error {
	if existing {
		return nil
	}
	fp := alert.Fingerprint()
	c := a.db.MsgAlert.Create().SetLabels(&alert.Labels).SetAnnotations(&alert.Annotations).
		SetStartsAt(alert.StartsAt).SetEndsAt(alert.EndsAt).SetURL(alert.GeneratorURL).
		SetTimeout(alert.Timeout).SetFingerprint(strconv.Itoa(int(fp)))
	if s, ok := alert.Labels[label.TenantLabel]; ok {
		tid, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		c.SetTenantID(tid)
	} else {
		c.SetTenantID(0)
	}
	//id := alert.Fingerprint()
	return c.Exec(schemax.SkipTenantKey(context.Background()))
}

func (a *AlertCallback) PostStore(alert *alert.Alert, existing bool) {

}

func (a *AlertCallback) PostDelete(alert *alert.Alert) {
	fp := strconv.Itoa(int(alert.Fingerprint()))
	c := a.db.MsgAlert.Update().Where(msgalert.Fingerprint(fp), msgalert.Deleted(false)).
		SetDeleted(true)
	if s, ok := alert.Labels[label.TenantLabel]; ok {
		tid, err := strconv.Atoi(s)
		if err != nil {
			logger.Error("delete alert error", zap.Error(err), zap.Any("alert", alert))
			return
		}
		c.Where(msgalert.TenantID(tid))
	}
	err := c.Exec(schemax.SkipTenantKey(context.Background()))
	if err != nil {
		logger.Error("delete alert error", zap.Error(err))
	}
}

type NlogCallback struct {
	db *ent.Client
}

func (n NlogCallback) LoadData() ([]*nflog.Entry, error) {
	//query := n.db.Silence.Query()
	//query.Where(nlog.sen)
	return nil, nil
}

func (n NlogCallback) CreateLog(ctx context.Context, r *nflog.Receiver, gkey string,
	firingAlerts, resolvedAlerts []uint64, expiresAt time.Time) (int, error) {
	var tenantID int
	ts, _ := notify.Tenant(ctx)
	if ts != "" {
		tid, err := strconv.Atoi(ts)
		if err != nil {
			return 0, err
		}
		tenantID = tid
	}
	row, err := n.db.Nlog.Create().SetTenantID(tenantID).SetReceiver(r.Name).
		SetGroupKey(gkey).SetReceiverType(profile.ReceiverType(r.Integration)).SetIdx(int(r.Index)).
		SetExpiresAt(expiresAt).SetSendAt(time.Now()).Save(schemax.SkipTenantKey(context.Background()))

	if err != nil {
		return 0, err
	}
	return row.ID, nil
}
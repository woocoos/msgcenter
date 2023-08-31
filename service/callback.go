package service

import (
	"context"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msgalert"
	"github.com/woocoos/msgcenter/ent/nlog"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/provider/mem"
	"go.uber.org/zap"
	"strconv"
	"time"
)

var (
	_ mem.AlertStoreCallback = (*AlertCallback)(nil)
	_ notify.NLogCallback    = (*NlogCallback)(nil)
)

// AlertCallback 是Alert处理回调类,由于消息中心是系统服务,对于alert无法有明显的上下文,因此需要在回调中处理租户时采用的忽略.
type AlertCallback struct {
	db *ent.Client
}

func (a *AlertCallback) PreStore(alert *alert.Alert, existing bool) error {
	if existing {
		return nil
	}
	fp := alert.Fingerprint().String()
	c := a.db.MsgAlert.Create().SetLabels(&alert.Labels).SetAnnotations(&alert.Annotations).
		SetStartsAt(alert.StartsAt).SetURL(alert.GeneratorURL).
		SetTimeout(alert.Timeout).SetFingerprint(fp)
	if alert.EndsAt.IsZero() {
		c.SetNillableEndsAt(nil)
	} else {
		c.SetEndsAt(alert.EndsAt)
	}
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
	return c.Exec(schemax.SkipTenantPrivacy(context.Background()))
}

func (a *AlertCallback) PostStore(alert *alert.Alert, existing bool) {

}

func (a *AlertCallback) PostDelete(alert *alert.Alert) {
	fp := alert.Fingerprint().String()
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
	err := c.Exec(schemax.SkipTenantPrivacy(context.Background()))
	if err != nil {
		logger.Error("delete alert error", zap.Error(err))
	}
}

type NlogCallback struct {
	db *ent.Client
}

func (n NlogCallback) LoadData() ([]*notify.LogEntry, error) {
	// expireAt > now will be evicted from nlog cache
	ds, err := n.db.Nlog.Query().Where(nlog.ExpiresAtGT(time.Now())).
		WithAlerts().
		All(schemax.SkipTenantPrivacy(context.Background()))
	if err != nil {
		return nil, err
	}
	var logs []*notify.LogEntry
	for _, d := range ds {
		nas, err := d.Alerts(context.Background())
		if err != nil {
			return nil, err
		}
		var firingAlerts []uint64
		var resolvedAlerts []uint64
		for _, na := range nas {
			if na.State == alert.AlertResolved {
				fp, err := label.StringToFingerprint(na.Fingerprint)
				if err != nil {
					return nil, err
				}
				resolvedAlerts = append(resolvedAlerts, uint64(fp))
			} else {
				fp, err := label.StringToFingerprint(na.Fingerprint)
				if err != nil {
					return nil, err
				}
				firingAlerts = append(firingAlerts, uint64(fp))
			}
		}
		logs = append(logs, &notify.LogEntry{
			ID:             d.ID,
			Receiver:       d.Receiver,
			ReceiverType:   d.ReceiverType,
			GroupKey:       d.GroupKey,
			FiringAlerts:   firingAlerts,
			ResolvedAlerts: resolvedAlerts,
			ExpiresAt:      d.ExpiresAt,
		})
	}
	return logs, nil
}

func (n NlogCallback) updateAlerts(ctx context.Context, alerts []uint64, state alert.AlertStatus) (ids []int, err error) {
	if len(alerts) == 0 {
		return
	}
	fps := make([]string, len(alerts))
	for i, fp := range alerts {
		fps[i] = label.Fingerprint(fp).String()
	}
	ids, err = n.db.MsgAlert.Query().Where(msgalert.FingerprintIn(fps...),
		msgalert.StateEQ(alert.AlertNone)).IDs(ctx)
	if err != nil {
		return

	}
	err = n.db.MsgAlert.Update().Where(msgalert.IDIn(ids...)).SetState(state).Exec(ctx)
	return
}

func (n NlogCallback) CreateLog(ctx context.Context, r *profile.ReceiverKey, gkey string,
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
	tctx := schemax.SkipTenantPrivacy(ctx)
	var alertids []int
	if len(firingAlerts) > 0 {
		ids, err := n.updateAlerts(tctx, firingAlerts, alert.AlertFiring)
		if err != nil {
			return 0, err
		}
		alertids = append(alertids, ids...)
	}
	if len(resolvedAlerts) > 0 {
		ids, err := n.updateAlerts(tctx, firingAlerts, alert.AlertResolved)
		if err != nil {
			return 0, err
		}
		alertids = append(alertids, ids...)
	}
	row, err := n.db.Nlog.Create().SetTenantID(tenantID).SetReceiver(r.Name).
		SetGroupKey(gkey).SetReceiverType(profile.ReceiverType(r.Integration)).SetIdx(int(r.Index)).
		SetExpiresAt(expiresAt).SetSendAt(time.Now()).
		AddAlertIDs(alertids...).
		Save(tctx)

	if err != nil {
		return 0, err
	}
	return row.ID, nil
}

// EvictLog evict log from nlog cache. the rule is expireAt > now.so need not do anything
func (n NlogCallback) EvictLog(ctx context.Context, ids []int) {
	return
}

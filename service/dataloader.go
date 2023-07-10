package service

import (
	"context"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/silence"
	"github.com/woocoos/msgcenter/pkg/alert"
	"time"
)

// SilencesDataLoad is a data loader for silences.
func SilencesDataLoad(client *ent.Client) func(ids ...int) ([]*ent.Silence, error) {
	if client == nil {
		return nil
	}
	return func(ids ...int) ([]*ent.Silence, error) {
		query := client.Silence.Query()
		if len(ids) == 0 {
			query.Where(silence.EndsAtGT(time.Now()), silence.StateNotIn(alert.SilenceStateExpired))
		} else {
			query.Where(silence.IDIn(ids...))
		}
		return query.All(schemax.SkipTenantKey(context.Background()))
	}
}

package service

import (
	"context"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/silence"
	"github.com/woocoos/msgcenter/pkg/alert"
	sil "github.com/woocoos/msgcenter/silence"
	"time"
)

// SilencesDataLoad is a data loader for silences.
func SilencesDataLoad(client *ent.Client) func(ids ...int) ([]*sil.Entry, error) {
	if client == nil {
		return nil
	}
	return func(ids ...int) ([]*sil.Entry, error) {
		query := client.Silence.Query()
		if len(ids) == 0 {
			query.Where(silence.EndsAtGT(time.Now()), silence.StateNotIn(alert.SilenceStateExpired))
		} else {
			query.Where(silence.IDIn(ids...))
		}
		ds, err := query.Select(silence.FieldID, silence.FieldUpdatedAt, silence.FieldState, silence.FieldMatchers,
			silence.FieldStartsAt, silence.FieldEndsAt).
			All(schemax.SkipTenantKey(context.Background()))
		if err != nil {
			return nil, err
		}
		vals := make([]*sil.Entry, len(ds))
		for i, row := range ds {
			vals[i] = &sil.Entry{
				ID:        row.ID,
				UpdatedAt: row.UpdatedAt,
				State:     row.State,
				Matchers:  row.Matchers,
				StartsAt:  row.StartsAt,
				EndsAt:    row.EndsAt,
			}
		}
		return vals, nil
	}
}

package service

import (
	"context"
	"time"

	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msgsilence"
	"github.com/woocoos/msgcenter/service/silence"
)

// SilencesDataLoad is a data loader for silences.
func SilencesDataLoad(client *ent.Client) func(ids ...int) ([]*silence.Entry, error) {
	if client == nil {
		return nil
	}
	return func(ids ...int) ([]*silence.Entry, error) {
		query := client.MsgSilence.Query()
		if len(ids) == 0 {
			query.Where(msgsilence.EndsAtGT(time.Now()), msgsilence.StateNotIn(silence.SilenceStateExpired))
		} else {
			query.Where(msgsilence.IDIn(ids...))
		}
		ds, err := query.Select(msgsilence.FieldID, msgsilence.FieldUpdatedAt, msgsilence.FieldState, msgsilence.FieldMatchers,
			msgsilence.FieldStartsAt, msgsilence.FieldEndsAt).
			All(schemax.SkipTenantPrivacy(context.Background()))
		if err != nil {
			return nil, err
		}
		vals := make([]*silence.Entry, len(ds))
		for i, row := range ds {
			vals[i] = &silence.Entry{
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

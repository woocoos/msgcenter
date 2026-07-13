package notify

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/woocoos/msgcenter/pkg/alert"
)

func TestFanoutStage_Exec_AllSuccess(t *testing.T) {
	t.Parallel()

	var called atomic.Int32
	stage := FanoutStage{
		StageFunc(func(ctx context.Context, alerts ...*alert.Alert) (context.Context, []*alert.Alert, error) {
			called.Add(1)
			return ctx, alerts, nil
		}),
		StageFunc(func(ctx context.Context, alerts ...*alert.Alert) (context.Context, []*alert.Alert, error) {
			called.Add(1)
			return ctx, alerts, nil
		}),
	}

	ctx := context.Background()
	_, _, err := stage.Exec(ctx)
	require.NoError(t, err)
	assert.Equal(t, int32(2), called.Load())
}

func TestFanoutStage_Exec_CollectsAllErrors(t *testing.T) {
	t.Parallel()

	err1 := errors.New("error from stage 1")
	err2 := errors.New("error from stage 2")
	stage := FanoutStage{
		StageFunc(func(ctx context.Context, alerts ...*alert.Alert) (context.Context, []*alert.Alert, error) {
			return ctx, nil, err1
		}),
		StageFunc(func(ctx context.Context, alerts ...*alert.Alert) (context.Context, []*alert.Alert, error) {
			return ctx, nil, err2
		}),
	}

	ctx := context.Background()
	_, _, err := stage.Exec(ctx)
	require.Error(t, err)
	assert.ErrorIs(t, err, err1)
	assert.ErrorIs(t, err, err2)
}

func TestFanoutStage_Exec_ConcurrentSafety(t *testing.T) {
	t.Parallel()

	// Run many stages concurrently to exercise the mutex protection.
	// With the race detector, this would catch unsynchronized writes.
	const numStages = 50
	stages := make(FanoutStage, numStages)
	for i := range stages {
		stages[i] = StageFunc(func(ctx context.Context, alerts ...*alert.Alert) (context.Context, []*alert.Alert, error) {
			return ctx, nil, errors.New("fail")
		})
	}

	ctx := context.Background()
	_, _, err := stages.Exec(ctx)
	require.Error(t, err)
}

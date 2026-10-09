package notify

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/metrics"
	"github.com/woocoos/msgcenter/pkg/profile"
)

func init() {
	// 初始化 metrics，避免测试时 nil pointer
	if metrics.Nflog == nil {
		metrics.Nflog = metrics.NewNflogMetrics()
	}
}

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

// mockNLogCallback 用于测试的 NLogCallback 实现
type mockNLogCallback struct {
	createLogCalled atomic.Int32
	createLogFunc   func(ctx context.Context, r *profile.ReceiverKey, gkey string, firingAlerts, resolvedAlerts []uint64, expiresAt time.Time) (int, error)
}

func (m *mockNLogCallback) LoadData() ([]*LogEntry, error) {
	return nil, nil
}

func (m *mockNLogCallback) CreateLog(ctx context.Context, r *profile.ReceiverKey, gkey string, firingAlerts, resolvedAlerts []uint64, expiresAt time.Time) (int, error) {
	m.createLogCalled.Add(1)
	if m.createLogFunc != nil {
		return m.createLogFunc(ctx, r, gkey, firingAlerts, resolvedAlerts, expiresAt)
	}
	return 0, nil
}

func (m *mockNLogCallback) EvictLog(ctx context.Context, ids []string) {}

// TestSetNotifiesStage_SkipStore 验证 skipStore=Y 时：
// 1. NLogCallback.CreateLog 仍被调用（用于创建内存 entry）
// 2. 内存 entry 正常创建（DedupStage 可以查询到）
// 3. 数据库写入由 CreateLog 内部根据 context 中的 skipStore 标志决定是否跳过
func TestSetNotifiesStage_SkipStore(t *testing.T) {
	t.Parallel()

	// 创建 mock callback，模拟 skipStore 时返回 (0, nil)
	callback := &mockNLogCallback{
		createLogFunc: func(ctx context.Context, r *profile.ReceiverKey, gkey string, firingAlerts, resolvedAlerts []uint64, expiresAt time.Time) (int, error) {
			// 模拟 service/callback.go 中 NLogCallback.CreateLog 的行为
			if SkipStore(ctx) {
				return 0, nil // 跳过 DB 写入
			}
			return 123, nil // 正常返回 ID
		},
	}

	// 创建 Log
	nflog := &Log{
		LogOptions: LogOptions{
			Retention:           time.Hour * 120,
			NLogCallback:        callback,
			MaintenanceInterval: time.Minute * 15,
		},
		st: make(state),
	}

	recv := &profile.ReceiverKey{
		Name:        "test-receiver",
		Integration: "message",
		Index:       0,
	}
	gkey := "test-group-key"
	firingAlerts := []uint64{12345}

	stage := SetNotifiesStage{
		nflog: nflog,
		recv:  recv,
	}

	// 测试 1: skipStore=Y 时，CreateLog 仍被调用，内存 entry 正常创建
	t.Run("skipStore=Y creates memory entry", func(t *testing.T) {
		callback.createLogCalled.Store(0)

		ctx := context.Background()
		ctx = WithGroupKey(ctx, gkey)
		ctx = WithFiringAlerts(ctx, firingAlerts)
		ctx = WithResolvedAlerts(ctx, nil)
		ctx = WithRepeatInterval(ctx, time.Hour)

		alerts := []*alert.Alert{
			{
				Labels: label.LabelSet{
					label.SkipStoreLabel: "Y",
				},
			},
		}

		_, _, err := stage.Exec(ctx, alerts...)
		require.NoError(t, err)

		// 验证 CreateLog 被调用
		assert.Equal(t, int32(1), callback.createLogCalled.Load(), "CreateLog should be called")

		// 验证内存 entry 被创建
		entries, err := nflog.Query(QReceiver(recv), QGroupKey(gkey))
		require.NoError(t, err)
		require.Len(t, entries, 1, "memory entry should be created")

		// 验证 entry 的 ID 为 0（因为 skipStore 时 CreateLog 返回 0）
		assert.Equal(t, 0, entries[0].ID, "entry ID should be 0 when skipStore=Y")
		assert.Equal(t, gkey, entries[0].GroupKey)
		assert.Equal(t, firingAlerts, entries[0].FiringAlerts)
	})

	// 测试 2: 无 skipStore 时，正常创建 entry
	t.Run("without skipStore creates entry with ID", func(t *testing.T) {
		// 重新创建 Log 以清除状态
		nflog2 := &Log{
			LogOptions: LogOptions{
				Retention:           time.Hour * 120,
				NLogCallback:        callback,
				MaintenanceInterval: time.Minute * 15,
			},
			st: make(state),
		}
		stage2 := SetNotifiesStage{
			nflog: nflog2,
			recv:  recv,
		}
		callback.createLogCalled.Store(0)

		ctx := context.Background()
		ctx = WithGroupKey(ctx, gkey)
		ctx = WithFiringAlerts(ctx, firingAlerts)
		ctx = WithResolvedAlerts(ctx, nil)
		ctx = WithRepeatInterval(ctx, time.Hour)
		// 不设置 skipStore

		alerts := []*alert.Alert{
			{
				Labels: label.LabelSet{},
			},
		}

		_, _, err := stage2.Exec(ctx, alerts...)
		require.NoError(t, err)

		// 验证 CreateLog 被调用
		assert.Equal(t, int32(1), callback.createLogCalled.Load(), "CreateLog should be called")

		// 验证内存 entry 被创建，ID 为 123（mock 返回值）
		entries, err := nflog2.Query(QReceiver(recv), QGroupKey(gkey))
		require.NoError(t, err)
		require.Len(t, entries, 1, "memory entry should be created")
		assert.Equal(t, 123, entries[0].ID, "entry ID should be 123 when skipStore is not set")
	})
}

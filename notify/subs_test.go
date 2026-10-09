package notify

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/service/provider"
)

// mockAlerts 实现 provider.Alerts 接口，用于测试。
type mockAlerts struct {
	provider.Alerts
	getFunc   func(fp label.Fingerprint) (*alert.Alert, error)
	putFunc   func(ctx context.Context, alerts ...*alert.Alert) error
	putAlerts []*alert.Alert // 记录 Put 的 alerts
}

func (m *mockAlerts) Get(fp label.Fingerprint) (*alert.Alert, error) {
	if m.getFunc != nil {
		return m.getFunc(fp)
	}
	return nil, fmt.Errorf("alert not found: %s", fp)
}

func (m *mockAlerts) Put(ctx context.Context, alerts ...*alert.Alert) error {
	m.putAlerts = append(m.putAlerts, alerts...)
	if m.putFunc != nil {
		return m.putFunc(ctx, alerts...)
	}
	return nil
}

// mockSubscriber 实现 Subscriber 接口，返回固定的用户列表。
type mockSubscriber struct {
	users []UserInfo
}

func (m *mockSubscriber) SubUsers(_ context.Context, _ *alert.Alert) ([]UserInfo, error) {
	return m.users, nil
}

// TestEventSubscribeStage_CloneEndsAtFallback 测试当 mem.Alerts.Get 失败时，
// clone 的 EndsAt 和 StartsAt 使用兜底值（now+5min / now）。
//
// 复现场景：
// 1. flush 清空了 firing alert 副本的 EndsAt（设为零值）
// 2. EventSubscribeStage 创建 clone，clone 继承零值 EndsAt
// 3. clone 尝试从 mem.Alerts 查找原始 alert 恢复 EndsAt
// 4. 但原始 alert 在 aggrGroup store 中，不在 mem.Alerts 中 → Get 失败
// 5. 兜底逻辑触发：EndsAt = now+5min, StartsAt = now
func TestEventSubscribeStage_CloneEndsAtFallback(t *testing.T) {
	t.Parallel()

	// 模拟 flush 后的 alert：EndsAt 和 StartsAt 都是零值
	now := time.Now()
	originalAlert := &alert.Alert{
		Labels: label.LabelSet{
			label.AlertNameLabel:     "test-alert",
			label.TenantLabel:        "1000",
			label.SkipSubscribeLabel: "Y", // 跳过订阅检查
		},
		Annotations: label.LabelSet{},
		StartsAt:    time.Time{}, // 零值，模拟 flush 清空
		EndsAt:      time.Time{}, // 零值，模拟 flush 清空
		UpdatedAt:   now,
	}

	// mockAlerts.Get 返回错误，模拟原始 alert 不在 mem.Alerts 中
	mock := &mockAlerts{
		getFunc: func(fp label.Fingerprint) (*alert.Alert, error) {
			return nil, fmt.Errorf("alert not found in mem.Alerts")
		},
	}

	subscriber := &mockSubscriber{
		users: []UserInfo{
			{UserID: "101106", Name: "Test User"},
		},
	}

	stage := NewEventSubscribeStage(mock, subscriber)

	// 注意：由于 alert 有 SkipSubscribeLabel，Exec 会直接返回
	// 我们需要去掉这个 label 来测试 exec 内部逻辑
	originalAlert.Labels = label.LabelSet{
		label.AlertNameLabel: "test-alert",
		label.TenantLabel:    "1000",
	}

	ctx := context.Background()
	_, _, err := stage.Exec(ctx, originalAlert)
	require.NoError(t, err)

	// 验证 Put 被调用（clone + resolved original）
	require.NotEmpty(t, mock.putAlerts, "Put should be called")

	// 找到 clone（带 to.user_id label 的）
	var clone *alert.Alert
	for _, a := range mock.putAlerts {
		if uid, ok := a.Labels[label.ToUserIDLabel]; ok && uid == "101106" {
			clone = a
			break
		}
	}
	require.NotNil(t, clone, "clone for user 101106 should exist")

	// 验证兜底逻辑：EndsAt 和 StartsAt 不应该是零值
	assert.False(t, clone.EndsAt.IsZero(), "clone EndsAt should not be zero after fallback")
	assert.False(t, clone.StartsAt.IsZero(), "clone StartsAt should not be zero after fallback")

	// 验证 EndsAt 在未来（firing 状态）
	assert.True(t, clone.EndsAt.After(time.Now()), "clone EndsAt should be in the future (firing)")

	// 验证 EndsAt 大约是 now+5min（允许 10s 误差）
	expectedEndsAt := now.Add(5 * time.Minute)
	assert.WithinDuration(t, expectedEndsAt, clone.EndsAt, 10*time.Second,
		"clone EndsAt should be approximately now+5min")

	t.Logf("clone EndsAt: %v (expected ~%v)", clone.EndsAt, expectedEndsAt)
	t.Logf("clone StartsAt: %v", clone.StartsAt)
}

// TestEventSubscribeStage_CloneEndsAtFromMemAlerts 测试当 mem.Alerts.Get 成功时，
// clone 的 EndsAt 从原始 alert 恢复，不使用兜底值。
func TestEventSubscribeStage_CloneEndsAtFromMemAlerts(t *testing.T) {
	t.Parallel()

	now := time.Now()
	originalEndsAt := now.Add(10 * time.Minute)
	originalStartsAt := now.Add(-1 * time.Minute)

	// 模拟 flush 后的 alert：EndsAt 和 StartsAt 都是零值
	flushedAlert := &alert.Alert{
		Labels: label.LabelSet{
			label.AlertNameLabel: "test-alert",
			label.TenantLabel:    "1000",
		},
		Annotations: label.LabelSet{},
		StartsAt:    time.Time{}, // 零值
		EndsAt:      time.Time{}, // 零值
		UpdatedAt:   now,
	}

	// mem.Alerts 中的原始 alert（有正确的 EndsAt）
	originalInMem := &alert.Alert{
		Labels: label.LabelSet{
			label.AlertNameLabel: "test-alert",
			label.TenantLabel:    "1000",
		},
		StartsAt:  originalStartsAt,
		EndsAt:    originalEndsAt,
		UpdatedAt: now,
	}

	// mockAlerts.Get 返回原始 alert
	mock := &mockAlerts{
		getFunc: func(fp label.Fingerprint) (*alert.Alert, error) {
			// 返回原始 alert（有正确的 EndsAt）
			return originalInMem, nil
		},
	}

	subscriber := &mockSubscriber{
		users: []UserInfo{
			{UserID: "101083", Name: "Test User"},
		},
	}

	stage := NewEventSubscribeStage(mock, subscriber)

	ctx := context.Background()
	_, _, err := stage.Exec(ctx, flushedAlert)
	require.NoError(t, err)

	// 找到 clone
	var clone *alert.Alert
	for _, a := range mock.putAlerts {
		if uid, ok := a.Labels[label.ToUserIDLabel]; ok && uid == "101083" {
			clone = a
			break
		}
	}
	require.NotNil(t, clone, "clone should exist")

	// 验证 EndsAt 从原始 alert 恢复，不是兜底值
	assert.Equal(t, originalEndsAt, clone.EndsAt,
		"clone EndsAt should be restored from original alert, not fallback")
	assert.Equal(t, originalStartsAt, clone.StartsAt,
		"clone StartsAt should be restored from original alert")

	t.Logf("clone EndsAt: %v (from original: %v)", clone.EndsAt, originalEndsAt)
}

// TestEventSubscribeStage_CloneEndsAtFallbackWhenResolved 测试场景：
// 当 mem.Alerts.Get 成功但返回的 alert 已 resolved 时，
// clone 应使用兜底值而非已过期原始 alert 的 EndsAt。
//
// 复现场景：
// 1. Pipeline A 的 EventSubscribeStage 先执行，Put resolved original 到 mem.Alerts
// 2. Pipeline B 的 EventSubscribeStage 后执行，Get 返回已 resolved 的 alert
// 3. 修复后：检查 Resolved()，使用兜底值 now+5min
func TestEventSubscribeStage_CloneEndsAtFallbackWhenResolved(t *testing.T) {
	t.Parallel()

	now := time.Now()
	// 模拟已 resolved 的 alert（EndsAt 已过期）
	resolvedEndsAt := now.Add(-2 * time.Minute) // 2 分钟前已过期
	resolvedStartsAt := now.Add(-7 * time.Minute)

	// 模拟 flush 后的 alert：EndsAt 和 StartsAt 都是零值
	flushedAlert := &alert.Alert{
		Labels: label.LabelSet{
			label.AlertNameLabel: "test-alert",
			label.TenantLabel:    "1000",
		},
		Annotations: label.LabelSet{},
		StartsAt:    time.Time{}, // 零值
		EndsAt:      time.Time{}, // 零值
		UpdatedAt:   now,
	}

	// mem.Alerts 中的 alert 已被其他 pipeline 标记为 resolved
	resolvedInMem := &alert.Alert{
		Labels: label.LabelSet{
			label.AlertNameLabel: "test-alert",
			label.TenantLabel:    "1000",
		},
		StartsAt:  resolvedStartsAt,
		EndsAt:    resolvedEndsAt, // 已过期
		UpdatedAt: now,
	}

	// 验证 mock 的 alert 确实是 resolved
	require.True(t, resolvedInMem.Resolved(), "mock alert should be resolved")

	// mockAlerts.Get 返回已 resolved 的 alert
	mock := &mockAlerts{
		getFunc: func(fp label.Fingerprint) (*alert.Alert, error) {
			return resolvedInMem, nil // Get 成功，但 alert 已 resolved
		},
	}

	subscriber := &mockSubscriber{
		users: []UserInfo{
			{UserID: "101106", Name: "Test User"},
		},
	}

	stage := NewEventSubscribeStage(mock, subscriber)

	ctx := context.Background()
	_, _, err := stage.Exec(ctx, flushedAlert)
	require.NoError(t, err)

	// 找到 clone
	var clone *alert.Alert
	for _, a := range mock.putAlerts {
		if uid, ok := a.Labels[label.ToUserIDLabel]; ok && uid == "101106" {
			clone = a
			break
		}
	}
	require.NotNil(t, clone, "clone should exist")

	// 验证 clone 没有使用已 resolved 的原始 alert 的 EndsAt
	assert.NotEqual(t, resolvedEndsAt, clone.EndsAt,
		"clone EndsAt should NOT be from resolved original alert")

	// 验证 clone 使用了兜底值（now+5min）
	assert.False(t, clone.EndsAt.IsZero(), "clone EndsAt should not be zero after fallback")
	assert.True(t, clone.EndsAt.After(now), "clone EndsAt should be in the future (firing)")

	expectedEndsAt := now.Add(5 * time.Minute)
	assert.WithinDuration(t, expectedEndsAt, clone.EndsAt, 10*time.Second,
		"clone EndsAt should be approximately now+5min (fallback)")

	t.Logf("resolved original EndsAt: %v (should NOT be used)", resolvedEndsAt)
	t.Logf("clone EndsAt: %v (fallback to now+5min)", clone.EndsAt)
}

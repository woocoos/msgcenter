package msg

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/woocoos/msgcenter/api/graphql/model"
)

func newTestPubSub(t *testing.T) (*PubSub, *miniredis.Miniredis) {
	t.Helper()
	mr, err := miniredis.Run()
	require.NoError(t, err)

	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	serverID := "test-server-1"
	pb := NewPubSub(rdb, serverID)
	t.Cleanup(func() {
		rdb.Close()
		mr.Close()
	})
	return pb, mr
}

func TestNewPubSub(t *testing.T) {
	pb, _ := newTestPubSub(t)

	assert.NotNil(t, pb.client)
	assert.Equal(t, "test-server-1", pb.serverID)
	assert.NotNil(t, pb.active)
	assert.Empty(t, pb.active)
	assert.NotNil(t, pb.conns)
}

func TestRegisterAndHasDeviceConnection(t *testing.T) {
	pb, mr := newTestPubSub(t)

	pb.registerDevice("device-1")

	assert.True(t, pb.HasDeviceConnection("device-1"))
	assert.False(t, pb.HasDeviceConnection("device-unknown"))

	// 验证 Redis key 存在且值正确
	val, err := pb.client.Get(context.Background(), deviceConnKey+"device-1").Result()
	require.NoError(t, err)
	assert.Equal(t, "test-server-1", val)

	// 验证 TTL 已设置
	ttl := mr.TTL(deviceConnKey + "device-1")
	assert.True(t, ttl > 0)
}

func TestUnregisterDevice(t *testing.T) {
	pb, _ := newTestPubSub(t)

	pb.registerDevice("device-1")
	assert.True(t, pb.HasDeviceConnection("device-1"))

	pb.unregisterDevice("device-1")
	assert.False(t, pb.HasDeviceConnection("device-1"))
}

func TestUnregisterDeviceSafety(t *testing.T) {
	// 验证 Lua 脚本安全性：只有 key 指向本实例时才删除
	pb1, _ := newTestPubSub(t)
	mr2, _ := miniredis.Run()
	rdb2 := redis.NewClient(&redis.Options{Addr: mr2.Addr()})
	defer func() {
		rdb2.Close()
		mr2.Close()
	}()

	// 手动在 Redis 中设置 key 指向另一个 server
	pb1.client.Set(context.Background(), deviceConnKey+"device-1", "other-server", deviceConnTTL)

	// pb1 的 serverID 是 "test-server-1"，但 key 值是 "other-server"
	pb1.unregisterDevice("device-1")

	// key 应该仍然存在（不被误删）
	val, err := pb1.client.Get(context.Background(), deviceConnKey+"device-1").Result()
	require.NoError(t, err)
	assert.Equal(t, "other-server", val)
}

func TestHasDeviceConnectionCrossServer(t *testing.T) {
	// 两个服务实例共享同一个 Redis
	mr, err := miniredis.Run()
	require.NoError(t, err)
	defer mr.Close()

	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	defer rdb.Close()

	pb1 := NewPubSub(rdb, "server-1")
	pb2 := NewPubSub(rdb, "server-2")

	pb1.registerDevice("device-1")

	// server-1 注册的设备，server-1 能查到
	assert.True(t, pb1.HasDeviceConnection("device-1"))
	// server-2 查不到（因为 Redis key 值是 "server-1"，不匹配 "server-2"）
	assert.False(t, pb2.HasDeviceConnection("device-1"))
}

func TestRemoveConn(t *testing.T) {
	pb, _ := newTestPubSub(t)

	connID := uuid.New()
	filter := &model.MessageFilter{
		TenantID: 1,
		UserID:   100,
		DeviceID: "device-1",
	}
	conn := pb.AddConnBy(connID, filter)
	pb.registerDevice("device-1")

	assert.True(t, pb.HasDeviceConnection("device-1"))
	assert.Len(t, pb.conns, 1)
	_ = conn

	ctx := context.WithValue(context.Background(), connectionIDKey, connID)
	err := pb.RemoveConn(ctx)
	require.NoError(t, err)

	// 连接已移除
	assert.Empty(t, pb.conns)
	// 设备已注销
	assert.False(t, pb.HasDeviceConnection("device-1"))
}

func TestRemoveConnWrongID(t *testing.T) {
	// 验证修复后的 bug：不匹配的 connID 不应移除任何连接
	pb, _ := newTestPubSub(t)

	connID := uuid.New()
	filter := &model.MessageFilter{DeviceID: "device-1"}
	pb.AddConnBy(connID, filter)
	pb.registerDevice("device-1")

	// 用错误的 connID 调用 RemoveConn
	wrongID := uuid.New()
	ctx := context.WithValue(context.Background(), connectionIDKey, wrongID)
	err := pb.RemoveConn(ctx)
	require.NoError(t, err)

	// 连接不应被移除
	assert.Len(t, pb.conns, 1)
	assert.True(t, pb.HasDeviceConnection("device-1"))
}

func TestRemoveConnNoDeviceID(t *testing.T) {
	pb, _ := newTestPubSub(t)

	connID := uuid.New()
	filter := &model.MessageFilter{} // 无 DeviceID
	pb.AddConnBy(connID, filter)

	ctx := context.WithValue(context.Background(), connectionIDKey, connID)
	err := pb.RemoveConn(ctx)
	require.NoError(t, err)

	assert.Empty(t, pb.conns)
}

func TestStopCleansUpDevices(t *testing.T) {
	pb, _ := newTestPubSub(t)

	pb.registerDevice("device-1")
	pb.registerDevice("device-2")
	assert.True(t, pb.HasDeviceConnection("device-1"))
	assert.True(t, pb.HasDeviceConnection("device-2"))

	err := pb.Stop(context.Background())
	require.NoError(t, err)

	assert.False(t, pb.HasDeviceConnection("device-1"))
	assert.False(t, pb.HasDeviceConnection("device-2"))
}

func TestRefreshLoop(t *testing.T) {
	pb, mr := newTestPubSub(t)

	pb.registerDevice("device-1")

	// 快进时间，让 TTL 减少
	mr.FastForward(60 * time.Second)

	// 刷新前 TTL 应该减少了
	ttl := mr.TTL(deviceConnKey + "device-1")
	assert.True(t, ttl > 0)
	assert.True(t, ttl < deviceConnTTL)

	// 手动调用 RefreshAll 模拟刷新
	pb.RefreshAll()

	// TTL 应该被重置
	ttl = mr.TTL(deviceConnKey + "device-1")
	assert.True(t, ttl > 60*time.Second)
}

func TestHasDeviceConnectionNoRedis(t *testing.T) {
	// 无 Redis 时回退到本地内存
	pb := NewPubSub(nil, "server-1")

	pb.active["device-1"] = true
	assert.True(t, pb.HasDeviceConnection("device-1"))
	assert.False(t, pb.HasDeviceConnection("device-2"))
}

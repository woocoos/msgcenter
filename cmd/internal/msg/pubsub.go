package msg

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/knockout-go/pkg/identity"
	"github.com/woocoos/msgcenter/api/graphql"
	"github.com/woocoos/msgcenter/api/graphql/model"
	"github.com/woocoos/msgcenter/pkg/push"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

var logger = log.Component("msgcenter.push")

const (
	deviceConnKey     = "mc:ws:dev:"
	deviceConnTTL     = 90 * time.Second
	appCodeHeaderKey  = "X-App-Code"
	deviceIDHeaderKey = "X-Device-ID"
)

// Connection 对应客户端连接,共享队列机制.连接在用户真正订阅时才会创建连接.
type Connection struct {
	ID          uuid.UUID
	Filter      model.MessageFilter
	Subscribers map[string]chan *model.Message

	mu sync.RWMutex
}

// Find 查找指定 topic 的订阅 channel
func (c *Connection) Find(topic string) (chan *model.Message, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	ch, ok := c.Subscribers[topic]
	return ch, ok
}

// PubSub 订阅管理器
type PubSub struct {
	conns    []*Connection
	client   redis.UniversalClient
	mu       sync.RWMutex
	serverID string
	active   map[string]bool // deviceId -> active
}

func NewPubSub(client redis.UniversalClient, serverID string) *PubSub {
	return &PubSub{
		client:   client,
		conns:    make([]*Connection, 0, 100),
		serverID: serverID,
		active:   make(map[string]bool),
	}
}

func (pb *PubSub) GetFilter(ctx context.Context) (*model.MessageFilter, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	gctx, err := gql.FromIncomingContext(ctx)
	if err != nil {
		return nil, err
	}
	filter := model.MessageFilter{
		TenantID: tid,
		UserID:   uid,
		AppCode:  gctx.Request.Header.Get(appCodeHeaderKey),
		DeviceID: gctx.Request.Header.Get(deviceIDHeaderKey),
	}
	return &filter, nil
}

// RemoveConn 移除指定连接并清理设备注册信息
func (pb *PubSub) RemoveConn(conn *Connection) {
	pb.mu.Lock()
	deviceID := conn.Filter.DeviceID
	delete(pb.active, deviceID)
	for i, c := range pb.conns {
		if c.ID == conn.ID {
			pb.conns = append(pb.conns[:i], pb.conns[i+1:]...)
			break
		}
	}
	pb.mu.Unlock()

	if deviceID != "" && pb.client != nil {
		script := redis.NewScript(`
			if redis.call("GET", KEYS[1]) == ARGV[1] then
				return redis.call("DEL", KEYS[1])
			else
				return 0
			end
		`)
		script.Run(context.Background(), pb.client, []string{deviceConnKey + deviceID}, pb.serverID)
	}
}

// Subscribe 根据topic订阅消息.
func (pb *PubSub) Subscribe(ctx context.Context, topic string) (chan *model.Message, error) {
	filter, err := pb.GetFilter(ctx)
	if err != nil {
		return nil, err
	}
	return pb.subscribe(ctx, filter, topic)
}

func (pb *PubSub) subscribe(ctx context.Context, filter *model.MessageFilter, topic string) (chan *model.Message, error) {
	ch := make(chan *model.Message, 100)
	conn := &Connection{
		ID:          uuid.New(),
		Filter:      *filter,
		Subscribers: make(map[string]chan *model.Message),
	}
	conn.Subscribers[topic] = ch

	pb.mu.Lock()
	pb.conns = append(pb.conns, conn)
	pb.mu.Unlock()

	if filter.DeviceID != "" {
		pb.registerDevice(filter.DeviceID)
	}

	go func() {
		<-ctx.Done()
		pb.RemoveConn(conn)
	}()
	return ch, nil
}

// HasDeviceConnection 检查指定设备是否已有活跃的连接
func (pb *PubSub) HasDeviceConnection(deviceID string) bool {
	if pb.client == nil {
		pb.mu.RLock()
		defer pb.mu.RUnlock()
		_, ok := pb.active[deviceID]
		return ok
	}
	val, err := pb.client.Get(context.Background(), deviceConnKey+deviceID).Result()
	if err != nil {
		return false
	}
	return val == pb.serverID
}

func (pb *PubSub) registerDevice(deviceID string) {
	pb.mu.Lock()
	pb.active[deviceID] = true
	pb.mu.Unlock()

	if pb.client != nil {
		pb.client.Set(context.Background(), deviceConnKey+deviceID, pb.serverID, deviceConnTTL)
	}
}

func (pb *PubSub) unregisterDevice(deviceID string) {
	pb.mu.Lock()
	delete(pb.active, deviceID)
	pb.mu.Unlock()

	if pb.client != nil {
		script := redis.NewScript(`
			if redis.call("GET", KEYS[1]) == ARGV[1] then
				return redis.call("DEL", KEYS[1])
			else
				return 0
			end
		`)
		script.Run(context.Background(), pb.client, []string{deviceConnKey + deviceID}, pb.serverID)
	}
}

func (pb *PubSub) Start(ctx context.Context) error {
	if pb.client != nil {
		go pb.subRedis(ctx)
		go pb.refreshLoop(ctx)
	}
	return nil
}

func (pb *PubSub) Stop(ctx context.Context) error {
	pb.mu.RLock()
	ids := make([]string, 0, len(pb.active))
	for id := range pb.active {
		ids = append(ids, id)
	}
	pb.mu.RUnlock()
	for _, id := range ids {
		pb.unregisterDevice(id)
	}
	return pb.client.Close()
}

func (pb *PubSub) refreshLoop(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			pb.RefreshAll()
		case <-ctx.Done():
			return
		}
	}
}

// RefreshAll 刷新所有活跃连接的 TTL
func (pb *PubSub) RefreshAll() {
	if pb.client == nil {
		return
	}
	pb.mu.RLock()
	ids := make([]string, 0, len(pb.active))
	for id := range pb.active {
		ids = append(ids, id)
	}
	pb.mu.RUnlock()

	for _, id := range ids {
		pb.client.Expire(context.Background(), deviceConnKey+id, deviceConnTTL)
	}
}

// 连接redis订阅
func (pb *PubSub) subRedis(ctx context.Context) {
	topics := graphql.SubTopics()
	ch := pb.client.Subscribe(context.Background(), topics...)
	for {
		select {
		case msg := <-ch.Channel():
			if msg != nil {
				switch msg.Channel {
				case string(graphql.SubTopicMessage):
					pb.handlerMessage(msg.Payload)
				}
			}
		case <-ctx.Done():
			ch.Close()
			return
		}
	}
}

func (pb *PubSub) handlerMessage(body string) {
	data, err := push.Unmarshal([]byte(body))
	if err != nil {
		logger.Error("msg handle error", zap.Error(err))
		return
	}
	// 先收集匹配的 channel,避免持锁发送导致死锁
	var targets []chan *model.Message
	for _, conn := range pb.conns {
		ch, ok := conn.Find(data.Topic)
		if ok && match(conn.Filter, data.Audience) {
			targets = append(targets, ch)
		}
	}

	msg := convertMessage(data)
	for _, ch := range targets {
		select {
		case ch <- msg:
		default:
			logger.Warn("channel full, drop message")
		}
	}
}

func convertMessage(data *push.Data) *model.Message {
	extras := make(map[string]string)
	if data.Message.Extras != nil {
		for k, v := range data.Message.Extras {
			extras[string(k)] = v
		}
	}
	msg := &model.Message{
		Topic:   data.Topic,
		Title:   data.Message.Title,
		Content: data.Message.Content,
		Format:  string(data.Message.Format),
		SendAt:  time.Now(),
		Extras:  extras,
	}
	return msg
}

// 根据消息的订阅信息匹配
func match(filter model.MessageFilter, audience push.Audience) bool {
	if filter.AppCode != audience.AppCode {
		return false
	}
	// user id
	if slices.Index(audience.UserIDs, filter.UserID) == -1 {
		return false
	}
	//
	if len(audience.DeviceIDs) > 0 && filter.DeviceID == "" {
		return false
	} else if filter.DeviceID != "" && len(audience.DeviceIDs) > 0 {
		if slices.Index(audience.DeviceIDs, filter.DeviceID) == -1 {
			return false
		}
	}
	return true
}

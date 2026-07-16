package msg

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/knockout-go/pkg/identity"
	"github.com/woocoos/msgcenter/api/graphql"
	"github.com/woocoos/msgcenter/api/graphql/model"
	"github.com/woocoos/msgcenter/pkg/push"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
	"sync"
	"time"
)

var logger = log.Component("push")

const (
	connectionIDKey = "woocoos/msg/conn_id"
	deviceConnKey   = "mc:ws:dev:"
	deviceConnTTL   = 90 * time.Second
)

// Connection 对应客户端连接,共享队列机制.连接在用户真正订阅时才会创建连接.
type Connection struct {
	ID          uuid.UUID
	Filter      model.MessageFilter
	Subscribers map[string]chan *model.Message
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
	initPayload := transport.GetInitPayload(ctx)
	filter := model.MessageFilter{
		TenantID: tid,
		UserID:   uid,
		AppCode:  initPayload.GetString("appCode"),
		DeviceID: initPayload.GetString("deviceId"),
	}
	return &filter, nil
}

// GetConn 从上下文获取客户端连接
func (pb *PubSub) GetConn(ctx context.Context, connID uuid.UUID) (*Connection, error) {
	for _, v := range pb.conns {
		if v.ID == connID {
			return v, nil
		}
	}
	return nil, nil
}

func (pb *PubSub) AddConnBy(id uuid.UUID, filter *model.MessageFilter) *Connection {
	s := &Connection{
		ID:          id,
		Filter:      *filter,
		Subscribers: make(map[string]chan *model.Message),
	}
	pb.conns = append(pb.conns, s)
	return s
}

// RemoveConn 移除连接,忽略对于连接ID不匹配的错误
func (pb *PubSub) RemoveConn(ctx context.Context) error {
	connID, ok := ctx.Value(connectionIDKey).(uuid.UUID)
	if !ok {
		return nil
	}
	pb.mu.Lock()
	var deviceID string
	for i, conn := range pb.conns {
		if conn.ID == connID {
			deviceID = conn.Filter.DeviceID
			delete(pb.active, deviceID)
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
	return nil
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
	connID, ok := ctx.Value(connectionIDKey).(uuid.UUID)
	if !ok {
		return nil, errors.New("ws connection id not found")
	}
	ch := make(chan *model.Message, 100)
	conn, err := pb.GetConn(ctx, connID)
	if err != nil {
		return nil, err
	}
	if conn == nil {
		conn = pb.AddConnBy(connID, filter)
		if filter.DeviceID != "" {
			pb.registerDevice(filter.DeviceID)
		}
	}
	pb.mu.Lock()
	defer pb.mu.Unlock()

	conn.Subscribers[topic] = ch
	return ch, nil
}

// HasDeviceConnection 检查指定设备是否已有活跃的WebSocket连接
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
			switch msg.Channel {
			case string(graphql.SubTopicMessage):
				pb.handlerMessage(msg.Payload)
			}
		case <-ctx.Done():
			ch.Close()
			return
		}
	}
}

func (pb *PubSub) handlerMessage(body string) {
	pb.mu.RLock()
	defer pb.mu.RUnlock()
	data, err := push.Unmarshal([]byte(body))
	if err != nil {
		logger.Error("msg handle error", zap.Error(err))
	}
	for _, conn := range pb.conns {
		ch, ok := conn.Subscribers[data.Topic]
		if ok && match(conn.Filter, data.Audience) {
			msg := convertMessage(data)
			ch <- msg
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

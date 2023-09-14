package msg

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/redis/go-redis/v9"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/msgcenter/api/graphql"
	"github.com/woocoos/msgcenter/api/graphql/model"
	"github.com/woocoos/msgcenter/pkg/push"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
	"strconv"
	"sync"
	"time"
)

var logger = log.Component("push")

// Connection 对应客户端连接,共享队列机制.
type Connection struct {
	Filter      model.MessageFilter
	Subscribers map[string]chan *model.Message
}

func (c *Connection) Key() string {
	return c.Filter.AppCode + ":" + strconv.Itoa(c.Filter.UserID) + ":" + c.Filter.DeviceID
}

// PubSub 订阅管理器
type PubSub struct {
	conns  []*Connection
	client redis.UniversalClient
	mu     sync.RWMutex
}

func NewPubSub(client redis.UniversalClient) *PubSub {
	return &PubSub{
		client: client,
		conns:  make([]*Connection, 0, 100),
	}
}

func (pb *PubSub) GetFilter(ctx context.Context) (*model.MessageFilter, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	_, err = identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	initPayload := transport.GetInitPayload(ctx)
	filter := model.MessageFilter{
		UserID:   uid,
		AppCode:  initPayload.GetString("appCode"),
		DeviceID: initPayload.GetString("deviceId"),
	}
	return &filter, nil
}

// GetConn 从上下文获取客户端连接
func (pb *PubSub) GetConn(ctx context.Context, filter *model.MessageFilter) (*Connection, error) {
	for _, v := range pb.conns {
		if v.Filter == *filter {
			return v, nil
		}
	}
	return nil, nil
}

func (pb *PubSub) AddConnBy(filter *model.MessageFilter) *Connection {
	s := &Connection{
		Filter:      *filter,
		Subscribers: make(map[string]chan *model.Message),
	}
	pb.conns = append(pb.conns, s)
	return s
}

func (pb *PubSub) RemoveConn(ctx context.Context) {
	filter, err := pb.GetFilter(ctx)
	if err != nil {
		log.Errorf("RemoveConn:%v", err)
		return
	}
	current, err := pb.GetConn(ctx, filter)
	if err != nil {
		return
	}
	for i, conn := range pb.conns {
		if conn.Key() == current.Key() {
			pb.conns = append(pb.conns[:i], pb.conns[i+1:]...)
		}
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
	conn, err := pb.GetConn(ctx, filter)
	if err != nil {
		return nil, err
	}
	if conn == nil {
		conn = pb.AddConnBy(filter)
	}
	pb.mu.Lock()
	defer pb.mu.Unlock()

	conn.Subscribers[topic] = ch
	return ch, nil
}

func (pb *PubSub) Publish(tars []chan *model.Message, message *model.Message) {
	for _, tar := range tars {
		tar <- message
	}
}

func (pb *PubSub) Start(ctx context.Context) error {
	if pb.client != nil {
		go func() {
			pb.subRedis(ctx)
		}()
	}
	return nil
}

func (pb *PubSub) Stop(ctx context.Context) error {
	return pb.client.Close()
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
	msg := &model.Message{
		Topic:   data.Topic,
		Title:   data.Message.Title,
		Content: data.Message.Content,
		Format:  string(data.Message.Format),
		SendAt:  time.Now(),
	}
	return msg
}

func match(filter model.MessageFilter, audience push.Audience) bool {
	if filter.AppCode != audience.AppCode {
		return false
	}
	// user id
	if slices.Index(audience.UserIDs, filter.UserID) == -1 {
		return false
	}
	//
	if filter.DeviceID != "" {
		if slices.Index(audience.DeviceIDs, filter.DeviceID) == -1 {
			return false
		}
	}
	return true
}

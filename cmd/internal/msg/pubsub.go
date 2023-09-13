package msg

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/redis/go-redis/v9"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/msgcenter/api/graphql/model"
	"strconv"
	"sync"
)

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
	mu     sync.Mutex
}

func NewPubSub(client redis.UniversalClient) *PubSub {
	return &PubSub{
		client: client,
		conns:  make([]*Connection, 0, 100),
	}
}

func (cm *PubSub) GetFilter(ctx context.Context) (*model.MessageFilter, error) {
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
func (cm *PubSub) GetConn(ctx context.Context, filter *model.MessageFilter) (*Connection, error) {
	for _, v := range cm.conns {
		if v.Filter == *filter {
			return v, nil
		}
	}
	return nil, nil
}

func (cm *PubSub) AddConnBy(filter *model.MessageFilter) *Connection {
	s := &Connection{
		Filter:      *filter,
		Subscribers: make(map[string]chan *model.Message),
	}
	cm.conns = append(cm.conns, s)
	return s
}

func (cm *PubSub) RemoveConn(ctx context.Context) {
	filter, err := cm.GetFilter(ctx)
	if err != nil {
		log.Errorf("RemoveConn:%v", err)
		return
	}
	current, err := cm.GetConn(ctx, filter)
	if err != nil {
		return
	}
	for i, conn := range cm.conns {
		if conn.Key() == current.Key() {
			cm.conns = append(cm.conns[:i], cm.conns[i+1:]...)
		}
	}
}

// Subscribe 根据topic订阅消息.
func (cm *PubSub) Subscribe(ctx context.Context, topic string) (chan *model.Message, error) {
	filter, err := cm.GetFilter(ctx)
	if err != nil {
		return nil, err
	}
	return cm.subscribe(ctx, filter, topic)
}

func (cm *PubSub) subscribe(ctx context.Context, filter *model.MessageFilter, topic string) (chan *model.Message, error) {
	ch := make(chan *model.Message)
	conn, err := cm.GetConn(ctx, filter)
	if err != nil {
		return nil, err
	}
	if conn == nil {
		conn = cm.AddConnBy(filter)
	}
	cm.mu.Lock()
	defer cm.mu.Unlock()

	conn.Subscribers[topic] = ch
	return ch, nil
}

func (cm *PubSub) Publish(ss []*Connection, topic string, message *model.Message) {
	for _, s := range ss {
		go func(s *Connection) {
			select {
			case s.Subscribers[topic] <- message:
			default:
				delete(s.Subscribers, topic)
			}
		}(s)
	}
}

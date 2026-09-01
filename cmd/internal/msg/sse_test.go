package msg

import (
	"bufio"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/alicebob/miniredis/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo/pkg/security"
	"github.com/woocoos/knockout-go/pkg/identity"
	"github.com/woocoos/msgcenter/api/graphql"
	"github.com/woocoos/msgcenter/pkg/push"
)

type sseSuite struct {
	suite.Suite
	pubsub *PubSub
	rdb    *redis.Client
	mr     *miniredis.Miniredis
	server *httptest.Server
}

func TestSSESubscription(t *testing.T) {
	suite.Run(t, new(sseSuite))
}

func (s *sseSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)

	var err error
	s.mr, err = miniredis.Run()
	s.Require().NoError(err)

	s.rdb = redis.NewClient(&redis.Options{Addr: s.mr.Addr()})
	s.pubsub = NewPubSub(s.rdb, "test-server")
	s.pubsub.Start(context.Background())

	// 构建 GraphQL schema 和 SSE handler
	schema := graphql.NewSchema(graphql.WithPubSub(s.pubsub))
	gqlSrv := handler.New(schema)
	gqlSrv.AddTransport(transport.SSE{KeepAlivePingInterval: 30 * time.Second})
	gqlSrv.AddTransport(transport.POST{})

	// 包装 handler: 注入 identity 和 gin.Context（模拟中间件行为）
	baseHandler := gqlSrv
	wrappedHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = security.WithContext(ctx, security.NewGenericPrincipalByClaims(jwt.MapClaims{"sub": "1"}))
		ctx = identity.WithTenantID(ctx, 1)
		ginCtx := &gin.Context{Request: r}
		ctx = context.WithValue(ctx, gin.ContextKey, ginCtx)
		gqlSrv.ServeHTTP(w, r.WithContext(ctx))
	})
	_ = baseHandler

	s.server = httptest.NewServer(wrappedHandler)
}

func (s *sseSuite) TearDownSuite() {
	if s.server != nil {
		s.server.Close()
	}
	if s.pubsub != nil {
		s.pubsub.Stop(context.Background())
	}
	if s.rdb != nil {
		s.rdb.Close()
	}
	if s.mr != nil {
		s.mr.Close()
	}
}

func (s *sseSuite) TestSSESubscription() {
	// 发送 SSE 订阅请求
	req, err := http.NewRequest(http.MethodPost, s.server.URL+"/query",
		strings.NewReader(`{"query":"subscription { message { topic title content } }"}`))
	s.Require().NoError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("X-Tenant-ID", "1")
	req.Header.Set("X-App-Code", "test-app")
	req.Header.Set("X-Device-ID", "test-device-1")
	req.Header.Set("Authorization", "Bearer test-token")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.StatusCode)
	s.Require().Equal("text/event-stream", resp.Header.Get("Content-Type"))
	s.Require().Equal("keep-alive", resp.Header.Get("Connection"))
	defer resp.Body.Close()

	br := bufio.NewReader(resp.Body)

	// 读取 SSE 初始连接标记 (":\n\n")
	line, err := br.ReadString('\n')
	s.Require().NoError(err)
	s.Equal(":\n", line)
	line, err = br.ReadString('\n')
	s.Require().NoError(err)
	s.Equal("\n", line)

	// 等待订阅注册完成
	time.Sleep(300 * time.Millisecond)

	// 通过 Redis 发布第一条测试消息
	msg := &push.Data{
		Topic: "message",
		Audience: push.Audience{
			AppCode:   "test-app",
			UserIDs:   []int{1},
			DeviceIDs: []string{"test-device-1"},
		},
		Message: push.Message{
			Title:   "SSE测试标题",
			Content: "SSE测试内容",
			Format:  "text",
		},
	}
	body, err := json.Marshal(msg)
	s.Require().NoError(err)
	s.mr.Publish("message", string(body))

	// 读取 SSE 事件: "event: next\n"
	line, err = br.ReadString('\n')
	s.Require().NoError(err)
	s.Equal("event: next\n", line)

	// 读取数据行: "data: {...}\n"
	line, err = br.ReadString('\n')
	s.Require().NoError(err)
	s.True(strings.HasPrefix(line, "data: "))
	s.Contains(line, "SSE测试标题")
	s.Contains(line, "SSE测试内容")
	s.Contains(line, `"topic":"message"`)

	// 读取事件间空行: "\n"
	line, err = br.ReadString('\n')
	s.Require().NoError(err)
	s.Equal("\n", line)

	// 发布第二条消息,验证持续订阅
	msg.Message.Title = "第二条消息"
	body, _ = json.Marshal(msg)
	s.mr.Publish("message", string(body))

	line, err = br.ReadString('\n')
	s.Require().NoError(err)
	s.Equal("event: next\n", line)

	line, err = br.ReadString('\n')
	s.Require().NoError(err)
	s.Contains(line, "第二条消息")

	br.ReadString('\n') // 空行
}

func (s *sseSuite) TestSSESubscriptionFilterMismatch() {
	// 订阅者 appCode=filter-app, 但消息发往 other-app, 不应收到消息
	req, err := http.NewRequest(http.MethodPost, s.server.URL+"/query",
		strings.NewReader(`{"query":"subscription { message { topic title } }"}`))
	s.Require().NoError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("X-Tenant-ID", "1")
	req.Header.Set("X-App-Code", "filter-app")
	req.Header.Set("X-Device-ID", "filter-device-1")
	req.Header.Set("Authorization", "Bearer test-token")

	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()

	br := bufio.NewReader(resp.Body)
	br.ReadString('\n') // :\n
	br.ReadString('\n') // \n

	time.Sleep(300 * time.Millisecond)

	// 发布不匹配的消息（appCode 不同）
	msg := &push.Data{
		Topic: "message",
		Audience: push.Audience{
			AppCode: "other-app",
			UserIDs: []int{1},
		},
		Message: push.Message{
			Title:   "不应收到",
			Content: "内容",
			Format:  "text",
		},
	}
	body, _ := json.Marshal(msg)
	s.mr.Publish("message", string(body))

	// 设置短超时,验证不会收到消息
	done := make(chan string, 1)
	go func() {
		line, _ := br.ReadString('\n')
		done <- line
	}()

	select {
	case line := <-done:
		// 如果收到数据行则失败（不应收到不匹配的消息）
		if strings.HasPrefix(line, "data:") {
			s.Fail("不应收到不匹配的消息")
		}
		// 如果收到 ping 注释行是正常的心跳,不算失败
	case <-time.After(500 * time.Millisecond):
		// 超时期望行为: 没有收到任何消息
	}
}

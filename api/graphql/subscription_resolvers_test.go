package graphql

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/web/handler"
	"github.com/woocoos/msgcenter/api/graphql/model"
	"go.uber.org/zap"
)

// mockPubSub 模拟 PubSub 接口
type mockPubSub struct {
	subscribeCalled bool
	lastTopic       string
}

func (m *mockPubSub) Subscribe(ctx context.Context, topic string) (chan *model.Message, error) {
	m.subscribeCalled = true
	m.lastTopic = topic
	ch := make(chan *model.Message, 1)
	return ch, nil
}

func (m *mockPubSub) HasDeviceConnection(deviceID string) bool {
	return false
}

// TestSubscriptionResolver_Message_LogFieldsInjection 验证 subscription resolver 注入审计日志字段
func TestSubscriptionResolver_Message_LogFieldsInjection(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		setupContext   func(*gin.Context)
		expectedFields map[string]interface{}
	}{
		{
			name: "with userId and tenantId",
			setupContext: func(c *gin.Context) {
				// 初始化 LogCarrier
				logCarrier := handler.GetLogCarrierFromGinContext(c)
				if logCarrier == nil {
					logCarrier = &log.FieldCarrier{Fields: make([]zap.Field, 0)}
					c.Set(handler.AccessLogComponentName, logCarrier)
				}
			},
			expectedFields: map[string]interface{}{},
		},
		{
			name: "with deviceId",
			setupContext: func(c *gin.Context) {
				// 设置 deviceId header
				c.Request.Header.Set("X-Device-ID", "test-device-123")

				// 初始化 LogCarrier
				logCarrier := handler.GetLogCarrierFromGinContext(c)
				if logCarrier == nil {
					logCarrier = &log.FieldCarrier{Fields: make([]zap.Field, 0)}
					c.Set(handler.AccessLogComponentName, logCarrier)
				}
			},
			expectedFields: map[string]interface{}{
				"deviceId": "test-device-123",
			},
		},
		{
			name: "without deviceId header",
			setupContext: func(c *gin.Context) {
				// 不设置 X-Device-ID header

				logCarrier := handler.GetLogCarrierFromGinContext(c)
				if logCarrier == nil {
					logCarrier = &log.FieldCarrier{Fields: make([]zap.Field, 0)}
					c.Set(handler.AccessLogComponentName, logCarrier)
				}
			},
			expectedFields: map[string]interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建 gin context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(http.MethodGet, "/graphql", nil)

			// 设置测试上下文
			tt.setupContext(c)

			// 创建 resolver
			mockPub := &mockPubSub{}
			resolver := &subscriptionResolver{
				&Resolver{
					PubSub: mockPub,
				},
			}

			// 使用 gin context 作为 GraphQL context
			// FromIncomingContext 会尝试从 context 中获取 gin.Context
			ctx := context.WithValue(c.Request.Context(), gin.ContextKey, c)

			// 执行 resolver
			ch, err := resolver.Message(ctx)

			// 验证结果
			require.NoError(t, err)
			require.NotNil(t, ch)
			assert.True(t, mockPub.subscribeCalled)
			assert.Equal(t, "message", mockPub.lastTopic)

			// 验证日志字段注入
			logCarrier := handler.GetLogCarrierFromGinContext(c)
			require.NotNil(t, logCarrier, "LogCarrier should not be nil")

			// 检查期望的字段是否存在
			fieldMap := make(map[string]zap.Field)
			for _, field := range logCarrier.Fields {
				fieldMap[field.Key] = field
			}

			for key, expectedValue := range tt.expectedFields {
				field, exists := fieldMap[key]
				assert.True(t, exists, "Field %s should exist", key)

				// 验证字段值
				switch v := expectedValue.(type) {
				case int:
					assert.Equal(t, int64(v), field.Integer, "Field %s value mismatch", key)
				case string:
					assert.Equal(t, v, field.String, "Field %s value mismatch", key)
				}
			}
		})
	}
}

// TestSubscriptionResolver_Message_NoLogCarrier 验证没有 LogCarrier 时不 panic
func TestSubscriptionResolver_Message_NoLogCarrier(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// 创建 gin context，不初始化 LogCarrier
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/graphql", nil)

	// 创建 resolver
	mockPub := &mockPubSub{}
	resolver := &subscriptionResolver{
		&Resolver{
			PubSub: mockPub,
		},
	}

	// 使用 gin context 作为 GraphQL context
	ctx := context.WithValue(c.Request.Context(), gin.ContextKey, c)

	// 执行 resolver - 应该不 panic
	ch, err := resolver.Message(ctx)

	require.NoError(t, err)
	require.NotNil(t, ch)
	assert.True(t, mockPub.subscribeCalled)

	// LogCarrier 应该仍然为 nil
	logCarrier := handler.GetLogCarrierFromGinContext(c)
	assert.Nil(t, logCarrier)
}

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/woocoos/msgcenter/api/oas"
	"github.com/woocoos/msgcenter/pkg/push"
)

type PushService struct {
	rdb redis.UniversalClient
}

func NewPushService(rdb redis.UniversalClient) *PushService {
	return &PushService{rdb: rdb}
}

func (s *PushService) PostPush(c *gin.Context, request *oas.PostPushRequest) error {
	if err := request.PushData.Validate(); err != nil {
		return err
	}
	msg, err := push.Marshal(request.PushData)
	if err != nil {
		return err
	}
	return s.rdb.Publish(c, request.PushData.Topic, msg).Err()
}

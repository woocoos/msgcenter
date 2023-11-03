package oas

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/woocoos/msgcenter/pkg/push"
)

type PushService struct {
	rdb redis.UniversalClient
}

func NewPushService(rdb redis.UniversalClient) *PushService {
	return &PushService{rdb: rdb}
}

func (s *PushService) PostPush(c *gin.Context, request *PostPushRequest) error {
	if err := request.Data.Validate(); err != nil {
		return err
	}
	msg, err := push.Marshal(request.Data)
	if err != nil {
		return err
	}
	return s.rdb.Publish(c, request.Topic, msg).Err()
}

package note

import (
	"Noteus/global"
	"context"
)

type NousService struct{}

var ctx = context.Background()

func (s *NousService) GetKeyList() (keys []string, err error) {
	keys, err = global.GVA_REDIS.HKeys(ctx, "notes:nous").Result()
	return
}

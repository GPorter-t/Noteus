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

func (s *NousService) GetItem(key string) (item map[string]string, err error) {
	value, err := global.GVA_REDIS.HGet(ctx, "notes:nous", key).Result()
	item = make(map[string]string)
	item[key] = value
	return
}

package note

import (
	"Noteus/global"
	"Noteus/model/note"
	"context"
)

type NousService struct{}

var ctx = context.Background()

func (s *NousService) GetAll() (keys []string, err error) {
	keys, err = global.GVA_REDIS.HKeys(ctx, "notes:nous").Result()
	return
}

func (s *NousService) GetItem(key string) (value string, err error) {
	value, err = global.GVA_REDIS.HGet(ctx, "notes:nous", key).Result()
	return
}

func (s *NousService) PostItem(item note.Nous) (i map[string]string, err error) {
	_, err = global.GVA_REDIS.HSet(ctx, "notes:nous", item.Key, item.Value).Result()
	i = make(map[string]string)
	i[item.Key] = item.Value
	return
}

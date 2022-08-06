package note

import (
	"Noteus/global"
	"Noteus/model/note"
	"context"
	"math/rand"
)

type NousService struct{}

var ctx = context.Background()

func (s *NousService) GetAll() (keys []string, err error) {
	keys, err = global.GVA_REDIS.HKeys(ctx, "notes:nous").Result()
	return
}

func (s *NousService) GetItem(key string) (id string, value string, err error) {
	go func() {
		if global.GVA_STORE.LRU.Len("notes:nous") == 0 {
			res, e := s.GetAll()
			if e != nil {
				err = e
				return
			}
			for i := 0; i < global.GVA_CONFIG.System.LruMaxSize; i++ {
				index := rand.Intn(len(res))
				v, _ := global.GVA_REDIS.HGet(ctx, "notes:nous", res[index]).Result()
				global.GVA_STORE.LRU.Add("notes:nous", res[index], v)
			}
		}
	}()

	if key == "" {
		k, v, ok := global.GVA_STORE.LRU.GetBack("notes:nous")
		if ok {
			id = k
			value = v.(string)
		}
		return
	}
	value, err = global.GVA_REDIS.HGet(ctx, "notes:nous", key).Result()

	return
}

func (s *NousService) PostItem(item note.Nous) (i map[string]string, err error) {
	_, err = global.GVA_REDIS.HSet(ctx, "notes:nous", item.Key, item.Value).Result()
	i = make(map[string]string)
	i[item.Key] = item.Value
	return
}

func (s *NousService) DeleteItem(key string) (err error) {
	ok, err := global.GVA_REDIS.HExists(ctx, "notes:nous", key).Result()
	if err != nil {
		return
	}
	if ok {
		_, err = global.GVA_REDIS.HDel(ctx, "notes:nous", key).Result()
		if err != nil {
			return
		}
	}
	return
}

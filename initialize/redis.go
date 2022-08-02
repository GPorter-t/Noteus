package initialize

import (
	"Noteus/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func Redis() {
	redisCfg := global.GVA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis connect ping failed: ", err)
	} else {
		fmt.Println("redis connect ping succeeded: ", pong)
		global.GVA_REDIS = client
	}
}

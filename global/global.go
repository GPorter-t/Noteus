package global

import (
	"Noteus/config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
)

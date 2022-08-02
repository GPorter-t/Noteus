package core

import (
	"Noteus/global"
	"fmt"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("./config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Errorf("Fatal Error config file: %s\n", err)
	}

	v.WatchConfig()
	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}

package core

import (
	"Noteus/global"
	"Noteus/initialize"
	"fmt"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	if global.GVA_CONFIG.System.UseRedis {
		initialize.Redis()
	}

	Router := initialize.Routers()
	address := fmt.Sprintf("%s:%d", global.GVA_CONFIG.System.Host, global.GVA_CONFIG.System.Port)

	s := initServer(address, Router)
	s.ListenAndServe().Error()
}

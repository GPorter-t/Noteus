package main

import (
	"Noteus/core"
	"Noteus/global"
	"Noteus/initialize"
)

func main() {
	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	global.GVA_LOG.Info("Starting...")
	core.RunServer()
}

package main

import (
	"Noteus/core"
	"Noteus/global"
)

func main() {
	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	global.GVA_LOG.Info("Starting...")
	core.RunServer()
}

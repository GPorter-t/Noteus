package main

import (
	"Noteus/core"
	"Noteus/global"
)

func main() {
	global.GVA_VP = core.Viper()
	core.RunServer()
}

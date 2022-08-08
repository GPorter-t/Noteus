package system

import "Noteus/service"

type ApiGroups struct {
	SystemApi
	DBApi
}

var (
	systemService = service.ServiceGroupApp.SystemServiceGroup
	initDBService = service.ServiceGroupApp.SystemServiceGroup
)

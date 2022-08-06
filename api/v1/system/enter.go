package system

import "Noteus/service"

type ApiGroups struct {
	SystemApi
}

var (
	systemService = service.ServiceGroupApp.SystemService
)

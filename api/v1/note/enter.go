package note

import "Noteus/service"

type ApiGroups struct {
	NousApi
}

var (
	nousService = service.ServiceGroupApp.NoteServiceGroup
)

package service

import (
	"Noteus/service/note"
	"Noteus/service/system"
)

type ServiceGroup struct {
	NoteService   note.NousService
	SystemService system.UserService
}

var ServiceGroupApp = new(ServiceGroup)

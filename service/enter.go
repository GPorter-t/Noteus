package service

import (
	"Noteus/service/note"
	"Noteus/service/system"
)

type ServiceGroup struct {
	NoteServiceGroup   note.ServiceGroup
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

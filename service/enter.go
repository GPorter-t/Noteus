package service

import "Noteus/service/note"

type ServiceGroup struct {
	NoteService note.NousService
}

var ServiceGroupApp = new(ServiceGroup)

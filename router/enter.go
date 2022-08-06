package router

import (
	"Noteus/router/note"
	"Noteus/router/system"
)

type RouterGroup struct {
	Nous   note.NoteGroup
	System system.SystemGroup
}

var RouterGroupApp = new(RouterGroup)

package router

import "Noteus/router/note"

type RouterGroup struct {
	Nous note.NousRouter
}

var RouterGroupApp = new(RouterGroup)

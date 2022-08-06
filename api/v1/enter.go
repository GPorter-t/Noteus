package v1

import (
	"Noteus/api/v1/note"
	"Noteus/api/v1/system"
)

type ApiGroup struct {
	NoteApiGroup   note.ApiGroups
	SystemApiGroup system.ApiGroups
}

var ApiGroupApp = new(ApiGroup)

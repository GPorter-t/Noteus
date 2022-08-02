package v1

import "Noteus/api/v1/note"

type ApiGroup struct {
	NoteApiGroup note.ApiGroups
}

var ApiGroupApp = new(ApiGroup)

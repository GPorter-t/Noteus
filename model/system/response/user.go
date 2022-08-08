package response

import (
	"Noteus/model/system"
)

type UserRsp struct {
	User      *system.User
	SessionId string `json:"session_id"`
}

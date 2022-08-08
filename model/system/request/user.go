package request

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Mode     int    `json:"mode"`
}

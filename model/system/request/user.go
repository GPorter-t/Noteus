package request

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Mode     int    `json:"mode"` // mode=0 账号密码；model=1 邮箱验证码
}

type UserRegisterReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Captcha  string `json:"captcha"`
}

type Email struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

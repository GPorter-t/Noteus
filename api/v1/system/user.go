package system

import (
	"Noteus/model/common/response"
	systemReq "Noteus/model/system/request"
	systemRsp "Noteus/model/system/response"
	"github.com/gin-gonic/gin"
)

type SystemApi struct{}

func (s *SystemApi) GetCaptcha(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		response.FailWithMessage("username is empty", c)
		return
	}
	captcha := systemService.GetCaptcha(username)
	response.OkWithData(captcha, c)
}

func (s *SystemApi) Login(c *gin.Context) {
	var req systemReq.UserLoginReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user, _ := systemService.SelectLoginMode(req.Mode, req.Username, req.Password)
	user, sessionId, err := systemService.Login(user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(systemRsp.UserRsp{
		User:      user,
		SessionId: sessionId,
	}, c)
}

package system

import (
	"Noteus/model/common/response"
	systemReq "Noteus/model/system/request"
	"github.com/gin-gonic/gin"
)

type SystemApi struct{}

func (s *SystemApi) GetCaptcha(c *gin.Context) {
	captcha := systemService.GetCaptcha()
	response.OkWithData(captcha, c)
}

func (s *SystemApi) Login(c *gin.Context) {
	var req systemReq.UserLoginReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ok, err := systemService.VerifyCaptcha(req.Password)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(ok, c)
}

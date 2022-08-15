package api

import (
	"Noteus/global"
	"Noteus/model/common/response"
	email_response "Noteus/plugin/email/model/response"
	"Noteus/plugin/email/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EmailApi struct{}

// EmailTest
// request: POST https://host:port/email/emailTest
// request Body:
// response Body: {"code":0,"data":{},"msg":"发送成功"}
func (s *EmailApi) EmailTest(c *gin.Context) {
	if err := service.ServiceGroupApp.EmailTest(); err != nil {
		global.GVA_LOG.Error("发送失败", zap.Error(err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithMessage("发送成功", c)
	}
}

// SendEmail
// request: POST https://host:port/email/sendEmail
// request Body: {"to": "374744710@qq.com", "subject":"test", "body": "1234"}
// response Body: {"code":0,"data":{},"msg":"发送成功"}
func (s *EmailApi) SendEmail(c *gin.Context) {
	var email email_response.Email
	_ = c.ShouldBindJSON(&email)
	if err := service.ServiceGroupApp.SendEmail(email.To, email.Subject, email.Body); err != nil {
		global.GVA_LOG.Error("发送失败", zap.Error(err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithMessage("发送成功", c)
	}
}

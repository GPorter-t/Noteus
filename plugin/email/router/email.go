package router

import (
	"Noteus/plugin/email/api"
	"github.com/gin-gonic/gin"
)

type EmailRouter struct{}

func (r *EmailRouter) InitEmailRouter(Router *gin.RouterGroup) {
	//emailRouter := Router.Use(middleware.OperationRecord())
	emailRouter := Router
	EmailApi := api.ApiGroupApp.EmailApi.EmailTest
	SendEmail := api.ApiGroupApp.EmailApi.SendEmail
	{
		emailRouter.POST("emailTest", EmailApi)
		emailRouter.POST("sendEmail", SendEmail)
	}
}

package note

import (
	v1 "Noteus/api/v1"
	"github.com/gin-gonic/gin"
)

type NousRouter struct{}

func (r *NousRouter) InitNousRouter(router *gin.RouterGroup) {
	nousRouter := router.Group("nous")
	nousApi := v1.ApiGroupApp.NoteApiGroup.NousApi
	{
		nousRouter.GET("/all", nousApi.GetNousKeyList) // 获取所有keys
		nousRouter.GET("", nousApi.GetNousRandom)      // 随机获取一个常识
	}
}

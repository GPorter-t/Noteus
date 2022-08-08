package initialize

import (
	"Noteus/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	nousRouter := router.RouterGroupApp.Nous
	systemRouter := router.RouterGroupApp.System

	PublicGroup := Router.Group("")
	{
		// 健康检测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关

	}

	nousRouter.InitNousRouter(PublicGroup)
	systemRouter.InitUserRouter(PublicGroup)

	return Router
}

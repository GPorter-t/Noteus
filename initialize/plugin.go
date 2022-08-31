package initialize

import (
	"Noteus/global"
	"Noteus/plugin/email"
	"Noteus/utils/plugin"
	"fmt"

	chatroom "github.com/GPorter-t/gin-plugin-chatroom"
	"github.com/gin-gonic/gin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for _, item := range Plugin {
		PluginGroup := group.Group(item.RouterPath())
		item.Register(PluginGroup)
	}
}

func InstallPlugin(Router *gin.Engine) {
	PublicGroup := Router.Group("")
	fmt.Println("无鉴权插件安装==>", PublicGroup)
	PrivateGroup := Router.Group("")
	fmt.Println("鉴权插件安装==>", PrivateGroup)
	// PrivateGroup.User()
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.GVA_CONFIG.Email.To,
		global.GVA_CONFIG.Email.From,
		global.GVA_CONFIG.Email.Host,
		global.GVA_CONFIG.Email.Secret,
		global.GVA_CONFIG.Email.Nickname,
		global.GVA_CONFIG.Email.Port,
		global.GVA_CONFIG.Email.IsSSL,
	))

	PluginInit(PrivateGroup, chatroom.CreateChatRoomPlugin(
		global.GVA_CONFIG.Redis.Addr,
		global.GVA_CONFIG.Redis.Password,
		global.GVA_CONFIG.Redis.DB,
	))
}

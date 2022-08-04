package note

import (
	"Noteus/global"
	"Noteus/model/common/response"
	"github.com/gin-gonic/gin"
	"math/rand"
)

type NousApi struct{}

func (n *NousApi) GetNousKeyList(c *gin.Context) {
	keys, err := nousService.GetKeyList()
	if err != nil {
		global.GVA_LOG.Error("获取 keys 失败:" + err.Error())
		response.FailWithMessage("获取 keys 失败:"+err.Error(), c)
		return
	}
	response.OkWithData(keys, c)
	return
}

func (n *NousApi) GetNousRandom(c *gin.Context) {
	uuid := c.Query("uuid")
	if uuid == "" {
		keys, err := nousService.GetKeyList()
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			response.FailWithMessage("获取数据失败"+err.Error(), c)
			return
		}
		index := rand.Intn(len(keys))
		uuid = keys[index]
	}
	item, err := nousService.GetItem(uuid)
	if err != nil {
		global.GVA_LOG.Error("获取 key 失败:" + err.Error())
		response.FailWithMessage("获取 key 失败:"+err.Error(), c)
		return
	}
	response.OkWithData(item, c)
	return
}

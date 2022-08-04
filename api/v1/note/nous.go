package note

import (
	"Noteus/global"
	"Noteus/model/common/response"
	"Noteus/model/note"
	noteReq "Noteus/model/note/request"
	noteRsp "Noteus/model/note/response"
	"github.com/gin-gonic/gin"
	goUuid "github.com/satori/go.uuid"
	"math/rand"
)

type NousApi struct{}

func (n *NousApi) GetNousKeyList(c *gin.Context) {
	keys, err := nousService.GetAll()
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
		keys, err := nousService.GetAll()
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			response.FailWithMessage("获取数据失败"+err.Error(), c)
			return
		}
		index := rand.Intn(len(keys))
		uuid = keys[index]
	}
	value, err := nousService.GetItem(uuid)
	if err != nil {
		global.GVA_LOG.Error("获取 key 失败:" + err.Error())
		response.FailWithMessage("获取 key 失败:"+err.Error(), c)
		return
	}
	rsp := noteRsp.NousRsp{
		Uuid: uuid,
		Desc: value,
	}
	response.OkWithData(rsp, c)
	return
}

func (n *NousApi) PostNousItem(c *gin.Context) {
	var req noteReq.NousReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("提交 item 失败:" + err.Error())
		response.FailWithMessage("提交 item 失败:"+err.Error(), c)
		return
	}
	item := note.Nous{
		Key:   goUuid.NewV4().String(),
		Value: req.Desc,
	}
	_, err = nousService.PostItem(item)
	if err != nil {
		global.GVA_LOG.Error("提交 item 失败:" + err.Error())
		response.FailWithMessage("提交 item 失败:"+err.Error(), c)
		return
	}
	rsp := noteRsp.NousRsp{
		Uuid: item.Key,
		Desc: item.Value,
	}
	response.OkWithData(rsp, c)
	return
}

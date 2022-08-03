package note

import (
	"Noteus/global"
	"Noteus/model/common/response"
	"github.com/gin-gonic/gin"
)

type NousApi struct{}

func (n *NousApi) GetNousKeyList(c *gin.Context) {
	keys, err := nousService.GetKeyList()
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}
	response.OkWithData(keys, c)
}

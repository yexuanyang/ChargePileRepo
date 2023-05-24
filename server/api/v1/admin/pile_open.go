package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
)

type IsOpenRequest struct {
	request.IdsReq
	IsOpen bool `json:"is_open"`
}

// UpdateChargePileByIds 根据ID更新数据库开关
func (PileManageApi) UpdateChargePileByIds(c *gin.Context) {
	var cr IsOpenRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var chargePileList []system.ChargePile
	err = global.GVA_DB.Model(system.ChargePile{}).Find(&chargePileList, cr.Ids).Update("is_open", cr.IsOpen).Error
	if err != nil {
		response.FailWithMessage("更新数据库失败", c)
		return
	}
	response.OkWithMessage("更新开关成功", c)
}

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
	// todo 在这个地方增加逻辑，实现充电桩关闭时，对应的线程充电桩线程中断;充电桩开启时，对应的充电桩线程启动
	var cr IsOpenRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var chargePileList []system.ChargePile
	err = global.GVA_DB.Model(system.ChargePile{}).Find(&chargePileList, cr.Ids).Update("is_open", cr.IsOpen).Error
	if !cr.IsOpen {
		// todo 关闭充电桩，需要关闭对应的充电桩线程

	} else {
		// todo 开启充电桩，需要开启对应的充电桩线程

	}
	if err != nil {
		response.FailWithMessage("更新数据库失败", c)
		return
	}
	response.OkWithMessage("更新开关成功", c)
}

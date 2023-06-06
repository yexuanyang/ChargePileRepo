package admin

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/user"
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
	if !cr.IsOpen {
		// 遍历关闭的每一个充电桩
		for _, pile := range chargePileList {
			// 在这个充电桩所在的站点寻找这个充电桩
			for _, pile1 := range user.ChargeStations[pile.StationID-1].ChargePiles {
				if pile1.PileId == int(pile.ID) {
					// 中断这个充电桩的充电过程
					user.IsInterrupt[pile.StationID-1][pile1.PileId] = true
					// 关闭这个充电桩
					user.IsOpen[pile.StationID-1][pile1.PileId] = false
					break
				}
			}
		}
	} else {
		fmt.Println("enter open")
		// 遍历选中的充电桩
		for _, pile := range chargePileList {
			// 在这个充电桩所在的站点寻找这个充电桩,在数据库中的id一致
			for _, pile1 := range user.ChargeStations[pile.StationID-1].ChargePiles {
				if pile1.PileId == int(pile.ID) {
					// 打开这个充电桩
					user.IsOpen[pile.StationID-1][pile1.PileId] = true
					fmt.Println("open")
				}
			}
		}
	}
	if err != nil {
		response.FailWithMessage("更新数据库失败", c)
		return
	}
	response.OkWithMessage("更新开关成功", c)
}

package PileManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
)

func (PileRouter) PileOpen(c *gin.Context) {
	var stationSet StationSet
	var stationResponse Station
	var chargeStationModel []system.ChargeStationModel
	err := global.GVA_DB.Debug().Model(&system.ChargeStationModel{}).Preload("ChargePileModel").Find(&chargeStationModel).Error
	if err != nil {
		response.FailWithMessage("查询失败", c)
		return
	}
	// 查每个充电站信息
	for _, chargeStation := range chargeStationModel {
		stationResponse.ID = chargeStation.ID
		stationResponse.Name = chargeStation.Position
		global.GVA_DB.Debug().Model(&system.ChargePileModel{}).Where("station_id = ?", chargeStation.ID).
			Select("id").Scan(&stationResponse.ChargeArray)
		stationSet.StationArray = append(stationSet.StationArray, stationResponse)
	}
	response.OkWithData(stationSet, c)
}

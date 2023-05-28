package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
)

// StationSet 充电站集合
type StationSet struct {
	StationArray []Station `json:"station_array"` // 充电站数组
}

// Station 充电站
type Station struct {
	ID          uint   `json:"id"` // id标识
	Name        string `json:"name"`
	ChargeArray []int  `json:"charge_array"` // 充电桩id数组，充电站下的充电桩
}

func (PileManageApi) GetPileList(c *gin.Context) {
	var stationSet StationSet
	var stationResponse Station
	var chargeStationModel []system.ChargeStation
	err := global.GVA_DB.Model(&system.ChargeStation{}).Preload("ChargePile").Find(&chargeStationModel).Error
	if err != nil {
		response.FailWithMessage("查询失败", c)
		return
	}
	// 查每个充电站信息
	for _, chargeStation := range chargeStationModel {
		stationResponse.ID = chargeStation.ID
		stationResponse.Name = chargeStation.Position
		global.GVA_DB.Debug().Model(&system.ChargePile{}).Where("station_id = ?", chargeStation.ID).
			Select("id").Scan(&stationResponse.ChargeArray)
		stationSet.StationArray = append(stationSet.StationArray, stationResponse)
	}
	response.OkWithData(stationSet, c)
}

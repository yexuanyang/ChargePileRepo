package user

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	sysres "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type ChargeApi struct{}

func (chargePileApi *ChargeApi) GetChargePileList2(c *gin.Context) {
	auth := utils.GetUserAuthorityId(c)
	if auth != 888 {
		response.FailWithMessage("权限不足", c)
		return
	}
	stationId := 1
	if station, err := GetChargePileInfoList2(stationId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(station, c)
	}
}

func GetChargePileInfoList2(stationId int) (station sysres.StationResponse, err error) {
	var chargePile []system.ChargePile
	var order []system.Order
	stationIndex := stationId - 1
	global.GVA_DB.Where("station_id = ?", stationId).Find(&chargePile)

	station.ChargeArray = make([]sysres.ChargePileResponse, 0)
	for _, item := range chargePile {
		t := sysres.ChargePileResponse{
			ID:                  strconv.Itoa(int(item.ID)),
			Name:                "",
			Mode:                item.PileType,
			TotalCharge:         item.Electricity,
			TotalChargeTimes:    float64(item.ChargeCount),
			TotalChargeDuration: item.ChargeTime,
			TotalChargeFee:      item.ChargeCost,
			TotalChargeService:  item.ServiceCost,
			TotalFee:            item.TotalCost,
		}
		if item.IsOpen {
			t.State = "WORK"
		} else {
			t.State = "OFF"
		}
		global.GVA_DB.Where("stationId = ? AND pile_id = ? AND state <> 'WAITING'", stationId, item.ID).Find(&order)
		for _, item1 := range order {
			c := sysres.CarBlockResponse{
				CarID:  item1.CarId,
				UserID: strconv.Itoa(item1.UserId),
			}
			var car system.Car
			global.GVA_DB.Model(&system.Car{}).First(&car, "car_id = ?", item1.CarId)
			c.PowerCapacity = car.BatteryCapacity
			c.PowerCurrent = 0
			c.State = item1.State
			c.ApplyKwh = item1.ApplyKwh
			c.WaitTime = 0
			// 在调度队列中获取给汽车分配的number

			// 指定充电站下的充电桩的充电队列
			for _, item2 := range ChargeStations[stationIndex].ChargePiles[int(item.ID)].Cars {
				// 找到队列里的车辆
				if item2.CarId == item1.CarId {
					c.Number = item2.QueueId
					break
				}
			}
			t.CarBlocks = append(t.CarBlocks, c)
		}

		station.ChargeArray = append(station.ChargeArray, t)
	}
	station.WaitArray = make([]sysres.CarBlockResponse, 0)

	// 获取数据库中对应充电站id下状态为等待的订单
	global.GVA_DB.Model(&system.Order{}).Find(&order, "stationId = ? AND state = 'WAITING'", stationId)
	for _, item3 := range order {
		wc := sysres.CarBlockResponse{
			CarID:  item3.CarId,
			UserID: strconv.Itoa(item3.UserId),
		}
		var car system.Car
		global.GVA_DB.Model(&system.Car{}).First(&car, "car_id = ?", item3.CarId)
		wc.PowerCapacity = car.BatteryCapacity
		wc.PowerCurrent = 0
		wc.ApplyKwh = item3.ApplyKwh
		wc.WaitTime = 0
		wc.State = item3.State
		// 获取调度程序等待区里面的车辆信息
		for _, item4 := range ChargeStations[stationIndex].Waiting.Cars {
			if item4.CarId == item3.CarId {
				wc.Number = item4.QueueId
				break
			}
		}
		station.WaitArray = append(station.WaitArray, wc)
	}

	return station, nil
}

func (chargePileApi *ChargeApi) ChangeChargePile(c *gin.Context) {
	var changeReq request.ChangeChargePileRequest
	err := c.ShouldBindJSON(&changeReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var chargePileList []system.ChargePile
	if changeReq.Mode == "SWITCHON" {
		// 打开充电桩
		err = global.GVA_DB.Model(&system.ChargePile{}).Find(&chargePileList, changeReq.SwitchArray).Update("is_open", 1).Error
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		}
		for _, pile := range chargePileList {
			// 在这个充电桩所在的站点寻找这个充电桩,在数据库中的id一致
			for _, pile1 := range ChargeStations[pile.StationID-1].ChargePiles {
				if pile1.PileId == int(pile.ID) {
					// 打开这个充电桩
					IsOpen[pile.StationID-1][pile1.PileId] = true
					fmt.Println("open")
				}
			}
		}
	} else if changeReq.Mode == "SWITCHOFF" {
		// 关闭充电桩
		err = global.GVA_DB.Model(&system.ChargePile{}).Find(&chargePileList, changeReq.SwitchArray).Update("is_open", 0).Error
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		}
		for _, pile := range chargePileList {
			// 在这个充电桩所在的站点寻找这个充电桩
			for _, pile1 := range ChargeStations[pile.StationID-1].ChargePiles {
				if pile1.PileId == int(pile.ID) {
					// 中断这个充电桩的充电过程
					IsInterrupt[pile.StationID-1][pile1.PileId] = true
					// 关闭这个充电桩
					IsOpen[pile.StationID-1][pile1.PileId] = false
					break
				}
			}
		}
	}
	response.OkWithMessage("开关成功", c)
}

func (chargePileApi *ChargeApi) GetChargePileReport(c *gin.Context) {
	var pileId request.NormalChargeRequest
	err := c.ShouldBindQuery(&pileId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var res sysres.ChargePileReport
	var chargePile system.ChargePile
	global.GVA_DB.Model(&system.ChargePile{}).First(&chargePile)
	res.ID = strconv.Itoa(int(chargePile.ID))
	res.Name = ""
	res.Mode = chargePile.PileType
	if chargePile.IsOpen {
		res.State = "WORK"
	} else {
		res.State = "OFF"
	}
	res.TotalCharge = chargePile.Electricity
	res.TotalChargeTimes = float64(chargePile.ChargeCount)
	res.TotalChargeDuration = chargePile.ChargeTime
	res.TotalFee = chargePile.TotalCost
	res.TotalChargeService = chargePile.ServiceCost
	res.TotalChargeFee = chargePile.ChargeCost
	res.Time = time.Now().String()
	response.OkWithData(res, c)
}

func (chargePileApi *ChargeApi) GetChargePileInfo(c *gin.Context) {
	var pileId request.NormalChargeRequest
	err := c.ShouldBindQuery(&pileId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	db := global.GVA_DB
	var pile system.ChargePile
	var res sysres.ChargePileCarInfoResponse
	db.Model(&system.ChargePile{}).First(&pile, "id = ?", pileId.ChargeId).Debug()
	res.ID = strconv.Itoa(int(pile.ID))
	res.Name = ""
	if pile.IsOpen {
		res.State = "WORK"
	} else {
		res.State = "OFF"
	}
	res.Mode = pile.PileType
	station := int(pile.StationID) - 1

	var order []system.Order
	db.Model(&system.Order{}).Where("pile_id = ?", pileId.ChargeId).Find(&order)
	for _, item := range order {
		var car sysres.CarBlockResponse
		car.ApplyKwh = item.ApplyKwh
		car.WaitTime = 0
		car.CarID = item.CarId
		car.UserID = strconv.Itoa(item.UserId)
		car.PowerCurrent = 0
		var dbcar system.Car
		db.Model(&system.Car{}).Where("car_id = ?", item.CarId).First(&dbcar)
		car.PowerCapacity = dbcar.BatteryCapacity
		car.State = item.State
		for _, item := range ChargeStations[station].ChargePiles[int(pile.ID)].Cars {
			if item.CarId == car.CarID {
				car.Number = item.QueueId
				break
			}
		}
		res.CarBlocks = append(res.CarBlocks, car)
	}
	response.OkWithData(res, c)
}

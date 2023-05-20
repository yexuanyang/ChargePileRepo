package system

import "github.com/flipped-aurora/gin-vue-admin/server/model/system/ctype"

// ChargePileModel 充电桩表
type ChargePileModel struct {
	MODEL
	IsOpen             bool               `gorm:"default:false" json:"is_open"` // 是否正常工作
	PileType           ctype.ChargingType `json:"pile_type"`                    // 充电桩类型
	ChargeTime         float64            `json:"charge_time"`                  // 充电总时长
	Count              int                `json:"count"`                        // 系统启动后累计充电次数
	Electricity        float64            `json:"electricity"`                  // 充电总电量
	StationID          uint               `json:"station_id"`
	ChargeStationModel ChargeStationModel `gorm:"foreignKey:StationID" json:"-"`
	CarModel           []CarModel         `gorm:"foreignKey:PileID"` // 最多有两辆
	OrderModel         []OrderModel       `gorm:"foreignKey:PileID"`
}

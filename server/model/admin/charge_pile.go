// 自动生成模板ChargePile
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// ChargePile 结构体
type ChargePile struct {
	global.GVA_MODEL
	IsOpen      bool      `json:"isOpen" form:"isOpen" gorm:"column:is_open;comment:表述充电桩是否开启;"`
	PileType    string    `json:"pileType" form:"pileType" gorm:"column:pile_type;type:enum('快充','慢充','其他');comment:充电桩类型;"`
	ChargeCount int       `json:"chargeCount" form:"chargeCount" gorm:"column:charge_count;comment:充电桩累计充电的次数;"`
	StationId   int       `json:"stationId" form:"stationId" gorm:"column:station_id;comment:充电桩所在充电站的ID;"`
	Electricity float64   `json:"electricity" form:"electricity" gorm:"column:electricity;comment:充电桩累计充电度数;"`
	ChargeTime  time.Time `json:"chargeTime" form:"chargeTime" gorm:"column:charge_time;comment:充电桩开启之后的总充电时间;"`
}

// TableName ChargePile 表名
func (ChargePile) TableName() string {
	return "charge_piles"
}

// 自动生成模板Order
package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Order 结构体
type Order struct {
	global.GVA_MODEL
	UserId      int       `json:"userId,omitempty" form:"userId" gorm:"column:user_id;comment:订单所属的用户id;"`
	CarId       string    `json:"carId" form:"carId" gorm:"column:car_id;comment:充电的车牌号;size:20;"`
	ChargeType  string    `json:"chargeType" form:"chargeType" gorm:"column:charge_type;type:enum('快充','慢充','其他');comment:充电类型;"`
	ChargeCost  float64   `json:"chargeCost" form:"chargeCost" gorm:"column:charge_cost;comment:;size:22;"`
	Kwh         float64   `json:"kwh" form:"kwh" gorm:"column:kwh;comment:;size:22;"`
	Time        float64   `json:"time" form:"time" gorm:"column:time;comment:;"`
	PileId      int       `json:"pileId" form:"pileId" gorm:"column:pile_id;comment:;size:20;"`
	ServiceCost float64   `json:"serviceCost" form:"serviceCost" gorm:"column:service_cost;comment:;size:22;"`
	StartedAt   time.Time `json:"startedAt" form:"startedAt" gorm:"column:started_at;comment:;"`
	StopAt      time.Time `json:"stopAt" form:"stopAt" gorm:"column:stop_at;comment:;"`
	TotalCost   float64   `json:"totalCost" form:"totalCost" gorm:"column:total_cost;comment:;size:22;"`
}

// TableName Order 表名
func (Order) TableName() string {
	return "order"
}

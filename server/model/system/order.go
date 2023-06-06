// 自动生成模板Order
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Order 结构体
type Order struct {
	global.GVA_MODEL
	UserId      int           `json:"user_id" form:"user_id" gorm:"column:user_id;comment:订单所属的用户id;"`
	UserModel   SysUser       `json:"-" gorm:"foreignKey:UserId"`
	CarId       string        `json:"car_id" form:"car_id" gorm:"column:car_id;comment:充电的车牌号;size:20;"`
	Car         Car           `json:"-" gorm:"foreignKey:CarId"`
	ChargeType  string        `json:"mode" form:"mode" gorm:"column:charge_type;type:enum('F','T');comment:充电类型,F快充,T慢充;"`
	ChargeCost  float64       `json:"chargeCost" form:"chargeCost" gorm:"column:charge_cost;comment:;size:22;"`
	Kwh         float64       `json:"kwh" form:"kwh" gorm:"column:kwh;comment:实际充电度数;size:22;"`
	ApplyKwh    float64       `json:"apply_kwh" form:"apply_kwh" gorm:"column:apply_kwh;comment:申请充电度数;size:22;"`
	Time        float64       `json:"time" form:"time" gorm:"column:time;comment:;"`
	PileId      int           `json:"pileId" form:"pileId" gorm:"column:pile_id;comment:;size:20;"`
	ChargePile  ChargePile    `json:"-" gorm:"foreignKey:PileId"`
	ServiceCost float64       `json:"serviceCost" form:"serviceCost" gorm:"column:service_cost;comment:;size:22;"`
	StartedAt   time.Time     `json:"startedAt" form:"startedAt" gorm:"column:started_at;comment:;"`
	StopAt      time.Time     `json:"stopAt" form:"stopAt" gorm:"column:stop_at;comment:;"`
	TotalCost   float64       `json:"totalCost" form:"totalCost" gorm:"column:total_cost;comment:;size:22;"`
	State       string        `json:"state" gorm:"column:state;comment:根据车辆所在位置划分订单状态：WAITING,DISPATCHED,CHARGING,FINISHED"`
	StationId   int           `json:"stationId" gorm:"column:stationId;comment:充电站的id"`
	Station     ChargeStation `json:"-" gorm:"foreignKey:StationId"`
}

// TableName Order 表名
func (Order) TableName() string {
	return "order"
}

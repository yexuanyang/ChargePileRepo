package system

import "github.com/flipped-aurora/gin-vue-admin/server/model/system/ctype"

// CarModel 充电区汽车表
type CarModel struct {
	MODEL
	State              ctype.State        `gorm:"default:3" json:"state"` // 车辆状态，1：充电区充电中，2：充电区等待中 3：排队中
	ChargingType       ctype.ChargingType `json:"charging_type"`          // 车辆的充电类型
	QueueNum           uint               `json:"queue_num"`              // 排队号
	BatteryCapacity    float64            `json:"battery_capacity"`       // 车辆电池总容量(度)
	ChargeRequestCount float64            `json:"charge_request_count"`   // 请求充电量(度)
	WaitingTime        float64            `json:"waiting_time"`           // 排队时长
	ChargingTime       float64            `json:"charging_time"`          // 充电时长
	CarQueueID         uint               `json:"car_queue_id"`
	CarQueueModel      CarQueueModel      `gorm:"foreignKey:CarQueueID" json:"-"`
	UserID             uint               `json:"user_id"`
	UserModel          UserModel          `gorm:"foreignKey:UserID" json:"-"`
	PileID             uint               `json:"pile_id"`
	ChargePileModel    ChargePileModel    `gorm:"foreignKey:PileID" json:"-"`
}

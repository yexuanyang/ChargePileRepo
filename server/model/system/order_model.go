package system

import "time"

// OrderModel 订单表
type OrderModel struct {
	MODEL
	Date             time.Time        `json:"date"`         // 日期
	Kwh              float64          `json:"kwh"`          // 充电度数
	Time             time.Time        `json:"time"`         // 充电时长
	StartedAt        time.Time        `json:"started_at"`   // 启动时间
	StopAt           time.Time        `json:"stop_at"`      // 停止时间
	ChargeCost       float64          `json:"charge_cost"`  // 充电费用
	ServiceCost      float64          `json:"service_cost"` // 服务费用
	TotalCost        float64          `json:"total_cost"`   // 总费用
	ReportID         uint             `json:"report_id"`
	ReportFormsModel ReportFormsModel `gorm:"foreignKey:ReportID" json:"-"`
	UserID           uint             `json:"user_id"`
	UserModel        UserModel        `gorm:"foreignKey:UserID" json:"-"`
	PileID           uint             `json:"pile_id"`
	ChargePileModel  ChargePileModel  `gorm:"foreignKey:PileID" json:"-"`
}

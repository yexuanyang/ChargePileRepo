package system

import "time"

// ReportFormsModel 报表
type ReportFormsModel struct {
	MODEL
	Date           time.Time    `json:"date"`
	ChargeCountSum float64      `json:"charge_count_sum"` // 累计充电次数
	TimeSum        float64      `json:"time_sum"`         // 累计充电时长
	ElectricitySum float64      `json:"electricity_sum"`  // 累计充电量
	ChargeCostSum  float64      `json:"charge_cost_sum"`  // 累计充电费用
	ServiceCostSum float64      `json:"service_cost_sum"` // 累计服务费用
	TotalCostSum   float64      `json:"total_cost_sum"`   // 累计总费用
	OrderModel     []OrderModel `gorm:"foreignKey:ReportID" json:"-"`
}

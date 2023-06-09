package response

type ChargePileInfoResponse struct {
	PileId            int     `json:"pileId" gorm:"column:pileId"`
	ChargeCount       int     `json:"chargeCount" gorm:"column:chargeCount"`
	ChargeTime        float64 `json:"chargeTime" gorm:"column:chargeTime"`
	ChargeElectricity float64 `json:"chargeElectricity" gorm:"column:chargeElectricity"`
	ChargeCost        float64 `json:"chargeCost" form:"chargeCost" gorm:"column:chargeCost"`
	ServiceCost       float64 `json:"serviceCost" form:"serviceCost" gorm:"column:serviceCost"`
	TotalCost         float64 ` json:"totalCost" form:"totalCost" gorm:"column:totalCost"`
}

// OrderReportResponse 订单报表的返回信息
type OrderReportResponse struct {
	Date        string  `json:"date" gorm:"column:date"`
	TotalKwh    float64 `json:"totalKwh" gorm:"column:total_kwh"`
	TotalCost   float64 `json:"totalCost" gorm:"column:total_cost"`
	ServiceCost float64 `json:"serviceCost" gorm:"column:total_service_cost"`
	ChargeCost  float64 `json:"chargeCost" gorm:"column:total_charge_cost"`
}

package response

type ChargePileInfoResponse struct {
	ChargeCount       int     `json:"chargeCount" gorm:"column:chargeCount"`
	ChargeTime        float64 `json:"chargeTime" gorm:"column:chargeTime"`
	ChargeElectricity float64 `json:"chargeElectricity" gorm:"column:chargeElectricity"`
	//ChargeElectricityFee float64 `json:"chargeFee"`
	//ChargeServiceFee     float64 `json:"chargeServiceFee"`
	//ChargeFee            float64 `json:"ChargeFee"`
}

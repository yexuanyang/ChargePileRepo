package system

// ChargeStationModel 充电站表
type ChargeStationModel struct {
	MODEL
	ChargePileModel []ChargePileModel `gorm:"foreignKey:StationID"`
}

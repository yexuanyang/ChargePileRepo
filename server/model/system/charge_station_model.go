package system

// ChargeStationModel 充电站表
type ChargeStationModel struct {
	MODEL
	Position        string            `json:"position"`
	ChargePileModel []ChargePileModel `gorm:"foreignKey:StationID" json:"-"`
}

package system

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// ChargeStation 结构体
type ChargeStation struct {
	global.GVA_MODEL
	Position   string       `json:"position" form:"position" gorm:"column:position;comment:充电站位置;"`
	ChargePile []ChargePile `gorm:"foreignKey:StationID" json:"-"`
}

// TableName ChargeStation 表名
func (ChargeStation) TableName() string {
	return "charge_station"
}

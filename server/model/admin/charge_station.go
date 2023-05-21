// 自动生成模板ChargeStation
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ChargeStation 结构体
type ChargeStation struct {
	global.GVA_MODEL
	Positon string `json:"positon" form:"positon" gorm:"column:positon;comment:充电站位置;size:255;"`
}

// TableName ChargeStation 表名
func (ChargeStation) TableName() string {
	return "charge_station"
}

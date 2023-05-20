package system

import(
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type SysChargePile struct {
	global.GVA_MODEL
	WaitingNum  *int    `gorm:"column:waiting_num" json:"waiting_num,omitempty"`
	Type        string  `gorm:"column:type; not null; default:'slow'" json:"type"`
	State       string  `gorm:"column:state; not null; default:'wait'" json:"state"`
	Power       int64   `gorm:"column:power; not null" json:"power"`
	Location    *string `gorm:"column:location" json:"location,omitempty"`
}

func (SysChargePile) TableName() string {
	return "sys_charge_pile"
}

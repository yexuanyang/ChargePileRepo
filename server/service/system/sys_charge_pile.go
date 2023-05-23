package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type ChargePileService struct{}

func (chargePileService *ChargePileService) GetChargePileList() (list interface{}, total int64, err error) {
	db := global.GVA_DB.Model(&system.SysChargePile{})
	var chargePileList []system.SysChargePile
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&chargePileList).Error
	if err != nil {
		return
	}
	return chargePileList, total, err
}

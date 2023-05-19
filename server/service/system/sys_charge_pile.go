package system

import (
	"fmt"
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
	var chargePile system.SysChargePile
	global.GVA_DB.Take(&chargePile)
	fmt.Print(chargePile)
	return chargePileList, total, err
}

package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ChargeStationService struct {
}

// CreateChargeStation 创建ChargeStation记录
// Author [piexlmax](https://github.com/piexlmax)
func (chargeStationService *ChargeStationService) CreateChargeStation(chargeStation *admin.ChargeStation) (err error) {
	err = global.GVA_DB.Create(chargeStation).Error
	return err
}

// DeleteChargeStation 删除ChargeStation记录
// Author [piexlmax](https://github.com/piexlmax)
func (chargeStationService *ChargeStationService) DeleteChargeStation(chargeStation admin.ChargeStation) (err error) {
	err = global.GVA_DB.Delete(&chargeStation).Error
	return err
}

// DeleteChargeStationByIds 批量删除ChargeStation记录
// Author [piexlmax](https://github.com/piexlmax)
func (chargeStationService *ChargeStationService) DeleteChargeStationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]admin.ChargeStation{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateChargeStation 更新ChargeStation记录
// Author [piexlmax](https://github.com/piexlmax)
func (chargeStationService *ChargeStationService) UpdateChargeStation(chargeStation admin.ChargeStation) (err error) {
	err = global.GVA_DB.Save(&chargeStation).Error
	return err
}

// GetChargeStation 根据id获取ChargeStation记录
// Author [piexlmax](https://github.com/piexlmax)
func (chargeStationService *ChargeStationService) GetChargeStation(id uint) (chargeStation admin.ChargeStation, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&chargeStation).Error
	return
}

// GetChargeStationInfoList 分页获取ChargeStation记录
// Author [piexlmax](https://github.com/piexlmax)
func (chargeStationService *ChargeStationService) GetChargeStationInfoList(info adminReq.ChargeStationSearch) (list []admin.ChargeStation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&admin.ChargeStation{})
	var chargeStations []admin.ChargeStation
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Positon != "" {
		db = db.Where("positon LIKE ?", "%"+info.Positon+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&chargeStations).Error
	return chargeStations, total, err
}

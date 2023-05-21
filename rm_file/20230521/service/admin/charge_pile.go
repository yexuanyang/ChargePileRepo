package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ChargePileService struct {
}

// CreateChargePile 创建ChargePile记录
// Author [piexlmax](https://github.com/piexlmax)
func (chargePileService *ChargePileService) CreateChargePile(chargePile *admin.ChargePile) (err error) {
	err = global.GVA_DB.Create(chargePile).Error
	return err
}

// DeleteChargePile 删除ChargePile记录
// Author [piexlmax](https://github.com/piexlmax)
func (chargePileService *ChargePileService) DeleteChargePile(chargePile admin.ChargePile) (err error) {
	err = global.GVA_DB.Delete(&chargePile).Error
	return err
}

// DeleteChargePileByIds 批量删除ChargePile记录
// Author [piexlmax](https://github.com/piexlmax)
func (chargePileService *ChargePileService) DeleteChargePileByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]admin.ChargePile{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateChargePile 更新ChargePile记录
// Author [piexlmax](https://github.com/piexlmax)
func (chargePileService *ChargePileService) UpdateChargePile(chargePile admin.ChargePile) (err error) {
	err = global.GVA_DB.Save(&chargePile).Error
	return err
}

// GetChargePile 根据id获取ChargePile记录
// Author [piexlmax](https://github.com/piexlmax)
func (chargePileService *ChargePileService) GetChargePile(id uint) (chargePile admin.ChargePile, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&chargePile).Error
	return
}

// GetChargePileInfoList 分页获取ChargePile记录
// Author [piexlmax](https://github.com/piexlmax)
func (chargePileService *ChargePileService) GetChargePileInfoList(info adminReq.ChargePileSearch) (list []admin.ChargePile, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&admin.ChargePile{})
	var chargePiles []admin.ChargePile
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.PileType != "" {
		db = db.Where("pile_type = ?", info.PileType)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&chargePiles).Error
	return chargePiles, total, err
}

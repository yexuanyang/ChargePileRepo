package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
)

type CarService struct {
}

// CreateCar 创建Car记录
// Author [piexlmax](https://github.com/piexlmax)
func (carInfoService *CarService) CreateCar(carInfo *user.Car) (err error) {
	err = global.GVA_DB.Create(carInfo).Error
	return err
}

// DeleteCar 删除Car记录
// Author [piexlmax](https://github.com/piexlmax)
func (carInfoService *CarService) DeleteCar(carInfo user.Car) (err error) {
	err = global.GVA_DB.Delete(&carInfo).Error
	return err
}

// DeleteCarByIds 批量删除Car记录
// Author [piexlmax](https://github.com/piexlmax)
func (carInfoService *CarService) DeleteCarByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]user.Car{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCar 更新Car记录
// Author [piexlmax](https://github.com/piexlmax)
func (carInfoService *CarService) UpdateCar(carInfo user.Car) (err error) {
	err = global.GVA_DB.Save(&carInfo).Error
	return err
}

// GetCar 根据id获取Car记录
// Author [piexlmax](https://github.com/piexlmax)
func (carInfoService *CarService) GetCar(id uint) (carInfo user.Car, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&carInfo).Error
	return
}

// GetCarInfoList 分页获取Car记录
// Author [piexlmax](https://github.com/piexlmax)
func (carInfoService *CarService) GetCarInfoList(info userReq.CarSearch) (list []user.Car, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&user.Car{})
	var carInfos []user.Car
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.UserId != 0 {
		db = db.Where("user_id = ?", info.UserId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&carInfos).Error
	return carInfos, total, err
}

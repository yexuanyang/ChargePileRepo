package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"strconv"
)

type CarService struct {
}

// CreateCar 创建Car记录
// Author [piexlmax](https://github.com/piexlmax)
func (carInfoService *CarService) CreateCar(carInfo *system.Car) (err error) {
	err = global.GVA_DB.Create(carInfo).Error
	return err
}

// DeleteCar 删除Car记录
// Author [piexlmax](https://github.com/piexlmax)
func (carInfoService *CarService) DeleteCar(carInfo system.Car) (err error) {
	err = global.GVA_DB.Delete(&carInfo).Error
	return err
}

// DeleteCarByIds 批量删除Car记录
// Author [piexlmax](https://github.com/piexlmax)
func (carInfoService *CarService) DeleteCarByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]system.Car{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCar 更新Car记录
// Author [piexlmax](https://github.com/piexlmax)
func (carInfoService *CarService) UpdateCar(carInfo system.Car) (err error) {
	err = global.GVA_DB.Save(&carInfo).Error
	return err
}

func (carInfoService *CarService) UpdateCar2(carInfo userReq.CarUpdate) (err error) {
	var car system.Car
	global.GVA_DB.Model(&system.Car{}).Where("id = ?", carInfo.CarID).First(&car)
	car.CarName = carInfo.Name
	car.BatteryCapacity = carInfo.PowerCapacity
	global.GVA_DB.Save(&car)
	return err
}

// GetCar 根据id获取Car记录
// Author [piexlmax](https://github.com/piexlmax)
func (carInfoService *CarService) GetCar(id uint) (carInfo system.Car, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&carInfo).Error
	return
}

// GetCarInfoList 分页获取Car记录
// Author [piexlmax](https://github.com/piexlmax)
func (carInfoService *CarService) GetCarInfoList(info userReq.CarSearch) (list []system.Car, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.Car{})
	var carInfos []system.Car
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

func (carInfoService *CarService) GetCarInfoList2(id uint) (list []systemRes.CarResponse, total int64, err error) {
	db := global.GVA_DB.Model(&system.Car{})
	var carInfos []system.Car
	err = db.Find(&carInfos, "user_id = ?", id).Error
	res := make([]systemRes.CarResponse, 0, len(carInfos))
	for _, car := range carInfos {
		var t systemRes.CarResponse
		t.CarID = strconv.Itoa(int(car.ID))
		t.Name = &car.CarName
		t.PowerCapacity = car.BatteryCapacity
		t.PowerCurrent = 0
		res = append(res, t)
	}
	return res, int64(len(carInfos)), err
}

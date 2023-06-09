package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type UsersService struct {
}

// CreateUsers 创建Users记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UsersService) CreateUsers(userInfo *system.SysUser) (err error) {
	err = global.GVA_DB.Create(userInfo).Error
	return err
}

// DeleteUsers 删除Users记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UsersService) DeleteUsers(userInfo system.SysUser) (err error) {
	err = global.GVA_DB.Delete(&userInfo).Error
	return err
}

// DeleteUsersByIds 批量删除Users记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UsersService) DeleteUsersByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]system.SysUser{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUsers 更新Users记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UsersService) UpdateUsers(userInfo system.SysUser) (err error) {
	err = global.GVA_DB.Save(&userInfo).Error
	return err
}

// GetUsers 根据id获取Users记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UsersService) GetUsers(id uint) (userInfo system.SysUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userInfo).Error
	return
}

// GetUsersInfoList 分页获取Users记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UsersService) GetUsersInfoList(Info userReq.UsersSearch) (list []system.SysUser, total int64, err error) {
	limit := Info.PageSize
	offset := Info.PageSize * (Info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.SysUser{})
	var userInfos []system.SysUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if Info.StartCreatedAt != nil && Info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", Info.StartCreatedAt, Info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&userInfos).Error
	return userInfos, total, err
}

package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
)

type UsersService struct {
}

// CreateUsers 创建Users记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UsersService) CreateUsers(userInfo *user.Users) (err error) {
	err = global.GVA_DB.Create(userInfo).Error
	return err
}

// DeleteUsers 删除Users记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UsersService) DeleteUsers(userInfo user.Users) (err error) {
	err = global.GVA_DB.Delete(&userInfo).Error
	return err
}

// DeleteUsersByIds 批量删除Users记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UsersService) DeleteUsersByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]user.Users{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUsers 更新Users记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UsersService) UpdateUsers(userInfo user.Users) (err error) {
	err = global.GVA_DB.Save(&userInfo).Error
	return err
}

// GetUsers 根据id获取Users记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UsersService) GetUsers(id uint) (userInfo user.Users, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userInfo).Error
	return
}

// GetUsersInfoList 分页获取Users记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UsersService) GetUsersInfoList(info userReq.UsersSearch) (list []user.Users, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&user.Users{})
	var userInfos []user.Users
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&userInfos).Error
	return userInfos, total, err
}

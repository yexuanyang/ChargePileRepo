package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
)

type TestService struct {
}

// CreateTest 创建Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (Test22Service *TestService) CreateTest(Test22 *user.Test) (err error) {
	err = global.GVA_DB.Create(Test22).Error
	return err
}

// DeleteTest 删除Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (Test22Service *TestService) DeleteTest(Test22 user.Test) (err error) {
	err = global.GVA_DB.Delete(&Test22).Error
	return err
}

// DeleteTestByIds 批量删除Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (Test22Service *TestService) DeleteTestByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]user.Test{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTest 更新Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (Test22Service *TestService) UpdateTest(Test22 user.Test) (err error) {
	err = global.GVA_DB.Save(&Test22).Error
	return err
}

// GetTest 根据id获取Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (Test22Service *TestService) GetTest(id uint) (Test22 user.Test, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&Test22).Error
	return
}

// GetTestInfoList 分页获取Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (Test22Service *TestService) GetTestInfoList(info userReq.TestSearch) (list []user.Test, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&user.Test{})
	var Test22s []user.Test
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&Test22s).Error
	return Test22s, total, err
}

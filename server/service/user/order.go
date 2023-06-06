package user

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type OrderService struct {
}

// CreateOrder 创建Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) CreateOrder(order *system.Order) (err error) {
	err = global.GVA_DB.Create(order).Error
	return err
}

// DeleteOrder 删除Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) DeleteOrder(order system.Order) (err error) {
	err = global.GVA_DB.Delete(&order).Error
	return err
}

// DeleteOrderByIds 批量删除Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) DeleteOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]system.Order{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOrder 更新Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) UpdateOrder(order system.Order) (err error) {
	//err = global.GVA_DB.Save(&order).Error
	var oldOrder system.Order
	err = global.GVA_DB.First(&oldOrder, "id = ?", order.ID).Error
	if err != nil {
		return err
	}
	newOrder := oldOrder
	newOrder.ApplyKwh = order.ApplyKwh
	newOrder.ChargeType = order.ChargeType
	err = global.GVA_DB.Save(&newOrder).Error
	if err != nil {
		return err
	}
	return err
}

// GetOrder 根据id获取Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) GetOrder(id uint) (order system.Order, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&order).Error
	return
}

// GetOrderInfoList 分页获取Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) GetOrderInfoList(Info userReq.OrderSearch) (list []system.Order, total int64, err error) {
	limit := Info.PageSize
	offset := Info.PageSize * (Info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.Order{})
	var orders []system.Order
	// 如果有条件搜索 下方会自动创建搜索语句
	if Info.StartCreatedAt != nil && Info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", Info.StartCreatedAt, Info.EndCreatedAt)
	}
	if Info.CarId != "" {
		db = db.Where("car_id LIKE ?", "%"+Info.CarId+"%")
	}
	if Info.ChargeType != "" {
		db = db.Where("charge_type = ?", Info.ChargeType)
	}
	if Info.UserId != 0 {
		db = db.Where("user_id = ?", Info.UserId)
	}
	if Info.State != "" {
		db = db.Where("state = ?", Info.State)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&orders).Error
	return orders, total, err
}

func (orderService *OrderService) GetOrderInfoList2(Info userReq.OrderSearch2) (list []system.Order, total int64, err error) {
	db := global.GVA_DB.Model(&system.Order{}).Debug()
	if Info.Mode == "HISTORY" {
		tx := db.Where("user_id = ? AND state = 'FINISHED'", Info.Order.UserId).Find(&list)
		if tx.Error != nil {
			return nil, 0, err
		}
		err = db.Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
	} else if Info.Mode == "CURRENT" {
		tx := db.Where("user_id=? AND state <> 'FINISHED'", Info.Order.UserId).Find(&list)
		if tx.Error != nil {
			return nil, 0, err
		}
		err = db.Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
	} else {
		return nil, 0, errors.New("请求参数不符合规范")
	}
	return list, total, err
}

func (orderService *OrderService) GetFrontCars(order system.Order) (number int64, err error) {
	db := global.GVA_DB.Model(&system.Order{})
	err = db.Where("id < ? AND charge_type = ? AND state = 'WAITING'", order.ID, order.ChargeType).Count(&number).Error
	return number, err
}

// GetOrderByUserId 根据user_id获取Order记录
func (orderService *OrderService) GetOrderByUserId(id uint) (order system.Order, err error) {
	err = global.GVA_DB.Where("user_id = ?", id).Find(&order).Error
	return
}

package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"time"
)

type CarSearch struct {
	system.Car
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type CarSearch2 struct {
	UserId int `json:"user_id"`
}

type CarAdd struct {
	Name          *string  `json:"name"`           // 车辆名称
	PowerCapacity *float64 `json:"power_capacity"` // 总电量
	PowerCurrent  *float64 `json:"power_current"`  // 当前电量
}

type CarUpdate struct {
	CarID         string  `json:"car_id"`         // 车辆在数据库里的ID
	Name          string  `json:"name"`           // 车辆名称
	PowerCapacity float64 `json:"power_capacity"` // 总电量
	PowerCurrent  float64 `json:"power_current"`  // 当前电量
}

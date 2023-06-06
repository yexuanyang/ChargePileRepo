package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"time"
)

type OrderSearch struct {
	system.Order
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type OrderSearch2 struct {
	system.Order
	Mode string `json:"mode" form:"mode"`
}

type OrderUpdate struct {
	Id       string  `json:"id"`
	Mode     string  `json:"mode"`
	ApplyKwh float64 `json:"apply_kwh"`
}

type OrderDelete struct {
	Id string `json:"id"`
}

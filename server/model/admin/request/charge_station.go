package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ChargeStationSearch struct {
	admin.ChargeStation
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

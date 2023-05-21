package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ChargePileSearch struct {
	admin.ChargePile
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

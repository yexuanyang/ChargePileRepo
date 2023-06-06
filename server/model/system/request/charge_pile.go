package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"time"
)

type ChargePileSearch struct {
	system.ChargePile
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type ChangeChargePileRequest struct {
	Mode        string   `json:"mode"`         // 开关模式，SWITCHON - 打开; SWITCHOFF - 关闭
	SwitchArray []string `json:"switch_array"` // 开关的充电站数组
}

type NormalChargeRequest struct {
	ChargeId string `json:"charge_id" form:"charge_id"`
}

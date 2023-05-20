package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ChargeRouter struct{}

func (s *ChargeRouter) InitChargeRouter(Router *gin.RouterGroup) {
	chargeRouter := Router.Group("chargePile").Use(middleware.OperationRecord())
	chargeRouterWithoutRecord := Router.Group("chargePile")
	chargeRouterApi := v1.ApiGroupApp.SystemApiGroup.ChargeApi
	{
		chargeRouter.POST("createChargePile", chargeRouterApi.CreateChargePile) // 创建充电桩
	}
	{
		chargeRouterWithoutRecord.POST("getChargePileList", chargeRouterApi.GetChargePileList) // 获取所有充电桩
	}
}

package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ChargeStationRouter struct {
}

// InitChargeStationRouter 初始化 ChargeStation 路由信息
func (s *ChargeStationRouter) InitChargeStationRouter(Router *gin.RouterGroup) {
	chargeStationRouter := Router.Group("chargeStation").Use(middleware.OperationRecord())
	chargeStationRouterWithoutRecord := Router.Group("chargeStation")
	var chargeStationApi = v1.ApiGroupApp.AdminApiGroup.ChargeStationApi
	{
		chargeStationRouter.POST("createChargeStation", chargeStationApi.CreateChargeStation)             // 新建ChargeStation
		chargeStationRouter.DELETE("deleteChargeStation", chargeStationApi.DeleteChargeStation)           // 删除ChargeStation
		chargeStationRouter.DELETE("deleteChargeStationByIds", chargeStationApi.DeleteChargeStationByIds) // 批量删除ChargeStation
		chargeStationRouter.PUT("updateChargeStation", chargeStationApi.UpdateChargeStation)              // 更新ChargeStation
	}
	{
		chargeStationRouterWithoutRecord.GET("findChargeStation", chargeStationApi.FindChargeStation)       // 根据ID获取ChargeStation
		chargeStationRouterWithoutRecord.GET("getChargeStationList", chargeStationApi.GetChargeStationList) // 获取ChargeStation列表
	}
}

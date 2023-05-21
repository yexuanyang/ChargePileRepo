package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CarRouter struct {
}

// InitCarRouter 初始化 Car 路由信息
func (s *CarRouter) InitCarRouter(Router *gin.RouterGroup) {
	carInfoRouter := Router.Group("carInfo").Use(middleware.OperationRecord())
	carInfoRouterWithoutRecord := Router.Group("carInfo")
	var carInfoApi = v1.ApiGroupApp.UserApiGroup.CarApi
	{
		carInfoRouter.POST("createCar", carInfoApi.CreateCar)             // 新建Car
		carInfoRouter.DELETE("deleteCar", carInfoApi.DeleteCar)           // 删除Car
		carInfoRouter.DELETE("deleteCarByIds", carInfoApi.DeleteCarByIds) // 批量删除Car
		carInfoRouter.PUT("updateCar", carInfoApi.UpdateCar)              // 更新Car
	}
	{
		carInfoRouterWithoutRecord.GET("findCar", carInfoApi.FindCar)       // 根据ID获取Car
		carInfoRouterWithoutRecord.GET("getCarList", carInfoApi.GetCarList) // 获取Car列表
	}
}

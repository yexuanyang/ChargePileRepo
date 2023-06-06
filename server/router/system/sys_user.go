package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("client").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("client")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	orderApi := v1.ApiGroupApp.UserApiGroup.OrderApi
	carApi := v1.ApiGroupApp.UserApiGroup.CarApi
	{
		userRouterWithoutRecord.POST("register", baseApi.Register)        // 管理员注册账号
		userRouter.POST("changePassword", baseApi.ChangePassword)         // 用户修改密码
		userRouter.POST("setUserAuthority", baseApi.SetUserAuthority)     // 设置用户权限
		userRouter.DELETE("deleteUser", baseApi.DeleteUser)               // 删除用户
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo)                // 设置用户信息
		userRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)                // 设置自身信息
		userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
		userRouter.POST("resetPassword", baseApi.ResetPassword)           // 设置用户权限组
	}
	{
		userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList) // 分页获取用户列表
		userRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo)  // 获取自身信息

		userRouterWithoutRecord.GET("index/information", baseApi.GetUserInfo2)     // 获取自身信息
		userRouterWithoutRecord.POST("index/information", baseApi.UpdateUserInfo2) // 更新自身信息
		userRouterWithoutRecord.GET("index/order", orderApi.GetOrderListByUserId2) // 获取订单信息
		userRouterWithoutRecord.POST("index/order", orderApi.CreateOrder)          // 创建订单信息
		userRouterWithoutRecord.PATCH("index/order", orderApi.UpdateOrder)         // 更新订单信息
		userRouterWithoutRecord.DELETE("index/order", orderApi.DeleteOrder)        // 删除订单信息
		userRouterWithoutRecord.GET("index/car", carApi.GetCarListByUserId2)       // 查询用户车辆信息
		userRouterWithoutRecord.POST("index/car", carApi.CreateCar2)               //新增用户车辆信息
		userRouterWithoutRecord.PATCH("index/car", carApi.UpdateCar2)              //更新用户车辆信息

	}
}

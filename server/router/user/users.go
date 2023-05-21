package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UsersRouter struct {
}

// InitUsersRouter 初始化 Users 路由信息
func (s *UsersRouter) InitUsersRouter(Router *gin.RouterGroup) {
	userInfoRouter := Router.Group("userInfo").Use(middleware.OperationRecord())
	userInfoRouterWithoutRecord := Router.Group("userInfo")
	var userInfoApi = v1.ApiGroupApp.UserApiGroup.UsersApi
	{
		userInfoRouter.POST("createUsers", userInfoApi.CreateUsers)             // 新建Users
		userInfoRouter.DELETE("deleteUsers", userInfoApi.DeleteUsers)           // 删除Users
		userInfoRouter.DELETE("deleteUsersByIds", userInfoApi.DeleteUsersByIds) // 批量删除Users
		userInfoRouter.PUT("updateUsers", userInfoApi.UpdateUsers)              // 更新Users
	}
	{
		userInfoRouterWithoutRecord.GET("findUsers", userInfoApi.FindUsers)       // 根据ID获取Users
		userInfoRouterWithoutRecord.GET("getUsersList", userInfoApi.GetUsersList) // 获取Users列表
	}
}

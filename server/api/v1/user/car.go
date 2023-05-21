package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CarApi struct {
}

var carInfoService = service.ServiceGroupApp.UserServiceGroup.CarService

// CreateCar 创建Car
// @Tags Car
// @Summary 创建Car
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.Car true "创建Car"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /carInfo/createCar [post]
func (carInfoApi *CarApi) CreateCar(c *gin.Context) {
	var carInfo user.Car
	err := c.ShouldBindJSON(&carInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"UserId": {utils.NotEmpty()},
		"CarId":  {utils.NotEmpty()},
	}
	if err := utils.Verify(carInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := carInfoService.CreateCar(&carInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCar 删除Car
// @Tags Car
// @Summary 删除Car
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.Car true "删除Car"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /carInfo/deleteCar [delete]
func (carInfoApi *CarApi) DeleteCar(c *gin.Context) {
	var carInfo user.Car
	err := c.ShouldBindJSON(&carInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := carInfoService.DeleteCar(carInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCarByIds 批量删除Car
// @Tags Car
// @Summary 批量删除Car
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Car"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /carInfo/deleteCarByIds [delete]
func (carInfoApi *CarApi) DeleteCarByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := carInfoService.DeleteCarByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCar 更新Car
// @Tags Car
// @Summary 更新Car
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.Car true "更新Car"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /carInfo/updateCar [put]
func (carInfoApi *CarApi) UpdateCar(c *gin.Context) {
	var carInfo user.Car
	err := c.ShouldBindJSON(&carInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"UserId": {utils.NotEmpty()},
		"CarId":  {utils.NotEmpty()},
	}
	if err := utils.Verify(carInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := carInfoService.UpdateCar(carInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCar 用id查询Car
// @Tags Car
// @Summary 用id查询Car
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query user.Car true "用id查询Car"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /carInfo/findCar [get]
func (carInfoApi *CarApi) FindCar(c *gin.Context) {
	var carInfo user.Car
	err := c.ShouldBindQuery(&carInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recarInfo, err := carInfoService.GetCar(carInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recarInfo": recarInfo}, c)
	}
}

// GetCarList 分页获取Car列表
// @Tags Car
// @Summary 分页获取Car列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userReq.CarSearch true "分页获取Car列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /carInfo/getCarList [get]
func (carInfoApi *CarApi) GetCarList(c *gin.Context) {
	var pageInfo userReq.CarSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := carInfoService.GetCarInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

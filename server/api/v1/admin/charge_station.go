package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ChargeStationApi struct {
}

var chargeStationService = service.ServiceGroupApp.AdminServiceGroup.ChargeStationService

// CreateChargeStation 创建ChargeStation
// @Tags ChargeStation
// @Summary 创建ChargeStation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.ChargeStation true "创建ChargeStation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chargeStation/createChargeStation [post]
func (chargeStationApi *ChargeStationApi) CreateChargeStation(c *gin.Context) {
	var chargeStation admin.ChargeStation
	err := c.ShouldBindJSON(&chargeStation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := chargeStationService.CreateChargeStation(&chargeStation); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteChargeStation 删除ChargeStation
// @Tags ChargeStation
// @Summary 删除ChargeStation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.ChargeStation true "删除ChargeStation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chargeStation/deleteChargeStation [delete]
func (chargeStationApi *ChargeStationApi) DeleteChargeStation(c *gin.Context) {
	var chargeStation admin.ChargeStation
	err := c.ShouldBindJSON(&chargeStation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := chargeStationService.DeleteChargeStation(chargeStation); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteChargeStationByIds 批量删除ChargeStation
// @Tags ChargeStation
// @Summary 批量删除ChargeStation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ChargeStation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /chargeStation/deleteChargeStationByIds [delete]
func (chargeStationApi *ChargeStationApi) DeleteChargeStationByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := chargeStationService.DeleteChargeStationByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateChargeStation 更新ChargeStation
// @Tags ChargeStation
// @Summary 更新ChargeStation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.ChargeStation true "更新ChargeStation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /chargeStation/updateChargeStation [put]
func (chargeStationApi *ChargeStationApi) UpdateChargeStation(c *gin.Context) {
	var chargeStation admin.ChargeStation
	err := c.ShouldBindJSON(&chargeStation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := chargeStationService.UpdateChargeStation(chargeStation); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindChargeStation 用id查询ChargeStation
// @Tags ChargeStation
// @Summary 用id查询ChargeStation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query admin.ChargeStation true "用id查询ChargeStation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /chargeStation/findChargeStation [get]
func (chargeStationApi *ChargeStationApi) FindChargeStation(c *gin.Context) {
	var chargeStation admin.ChargeStation
	err := c.ShouldBindQuery(&chargeStation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rechargeStation, err := chargeStationService.GetChargeStation(chargeStation.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rechargeStation": rechargeStation}, c)
	}
}

// GetChargeStationList 分页获取ChargeStation列表
// @Tags ChargeStation
// @Summary 分页获取ChargeStation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.ChargeStationSearch true "分页获取ChargeStation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chargeStation/getChargeStationList [get]
func (chargeStationApi *ChargeStationApi) GetChargeStationList(c *gin.Context) {
	var pageInfo adminReq.ChargeStationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := chargeStationService.GetChargeStationInfoList(pageInfo); err != nil {
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

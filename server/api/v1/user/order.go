package user

import (
	"fmt"
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

type OrderApi struct {
}

var orderService = service.ServiceGroupApp.UserServiceGroup.OrderService

// CreateOrder 创建Order
// @Tags Order
// @Summary 创建Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.Order true "创建Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/createOrder [post]
func (orderApi *OrderApi) CreateOrder(c *gin.Context) {
	var order user.Order
	claims, err := utils.GetClaims(c)
	err = c.ShouldBindJSON(&order)
	if order.UserId == 0 {
		order.UserId = int(claims.BaseClaims.ID)
	}
	car := GetCarInfoByOrder(order)
	err = ChargeStations[order.StationId-1].Waiting.Enqueue(car)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	order.State = "等待区"

	verify := utils.Rules{
		"UserId":     {utils.NotEmpty()},
		"CarId":      {utils.NotEmpty()},
		"ChargeType": {utils.NotEmpty()},
		"Kwh":        {utils.NotEmpty()},
	}
	if err := utils.Verify(order, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	order.Kwh = 0
	if err := orderService.CreateOrder(&order); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrder 删除Order
// @Tags Order
// @Summary 删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.Order true "删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /order/deleteOrder [delete]
func (orderApi *OrderApi) DeleteOrder(c *gin.Context) {
	var order user.Order

	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	order, err = orderService.GetOrder(order.ID)
	fmt.Println(order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	car := GetCarInfoByOrder(order)
	CarNum[order.StationId-1][car.Mode] -= 1
	stationId := order.StationId - 1
	if order.State == "等待区" {
		// 将汽车从等待区中移除,订单在数据库中也删除
		err = ChargeStations[order.StationId].Waiting.Delete(car)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	} else if order.State == "队列区" {
		currentPile := ChargeStations[stationId].ChargePiles[order.PileId]
		if order.CarId == currentPile.Cars[0].CarId {
			// 该汽车正在充电
			currentPile.mu.Lock()
			defer currentPile.mu.Unlock()

			currentPile.Cars = append(currentPile.Cars[:0], currentPile.Cars[1:]...) // 将汽车移除

			FreePileMutex[stationId].Lock()
			defer FreePileMutex[stationId].Unlock()
			FreePile[stationId][order.PileId] += 1 // 空闲位置+1

			// 向指定的充电桩线程发送提前结束充电请求
			IsInterrupt[stationId][order.PileId] = true
			// 数据库中的订单不需要删除
			// 向前端返回成功即可
			response.OkWithMessage("删除成功", c)
			return

		} else {
			// 汽车不在充电，数据库中的订单也直接删除
			currentPile.mu.Lock() // 获取当前充电桩的充电队列的锁
			defer currentPile.mu.Unlock()

			currentPile.Cars = append(currentPile.Cars[:1], currentPile.Cars[2:]...) // 将汽车移除

			FreePileMutex[stationId].Lock()
			defer FreePileMutex[stationId].Unlock()
			FreePile[stationId][order.PileId] += 1 // 空闲位置+1
		}
	}

	if err := orderService.DeleteOrder(order); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrderByIds 批量删除Order
// @Tags Order
// @Summary 批量删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /order/deleteOrderByIds [delete]
func (orderApi *OrderApi) DeleteOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.DeleteOrderByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrder 更新Order
// @Tags Order
// @Summary 更新Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.Order true "更新Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /order/updateOrder [put]
func (orderApi *OrderApi) UpdateOrder(c *gin.Context) {
	var order user.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	claims, err := utils.GetClaims(c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if order.UserId == 0 {
		order.UserId = int(claims.BaseClaims.ID)
	}
	if order.ID == 0 {
		updateCar := GetCarInfoByOrder(order)
		err = ChargeStations[order.StationId-1].Waiting.Update(updateCar)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}
	verify := utils.Rules{
		"UserId":     {utils.NotEmpty()},
		"CarId":      {utils.NotEmpty()},
		"ChargeType": {utils.NotEmpty()},
		"Kwh":        {utils.NotEmpty()},
	}
	if err := utils.Verify(order, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.UpdateOrder(order); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrder 用id查询Order
// @Tags Order
// @Summary 用id查询Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query user.Order true "用id查询Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /order/findOrder [get]
func (orderApi *OrderApi) FindOrder(c *gin.Context) {
	var order user.Order
	err := c.ShouldBindQuery(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reorder, err := orderService.GetOrder(order.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reorder": reorder}, c)
	}
}

// GetOrderList 分页获取Order列表
// @Tags Order
// @Summary 分页获取Order列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userReq.OrderSearch true "分页获取Order列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/getOrderList [get]
func (orderApi *OrderApi) GetOrderList(c *gin.Context) {
	var pageInfo userReq.OrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := orderService.GetOrderInfoList(pageInfo); err != nil {
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

func (orderApi *OrderApi) GetOrderListByUserId(c *gin.Context) {
	claims, err := utils.GetClaims(c)
	var pageInfo userReq.OrderSearch
	err = c.ShouldBindQuery(&pageInfo)
	pageInfo.Order.UserId = int(claims.BaseClaims.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := orderService.GetOrderInfoList(pageInfo); err != nil {
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

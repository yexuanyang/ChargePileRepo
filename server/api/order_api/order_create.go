package order_api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type OrderRequest struct {
	ChargePileID int64   `json:"chargePileID"`        // 充电桩ID
	ChargeType   string  `json:"chargeType"`          // 充电类型
	Kwh          int64   `json:"kwh"`                 // 预计充电度数
	Price        *int64  `json:"price,omitempty"`     // 预计充电金额
	TimeToEnd    *string `json:"timeToEnd,omitempty"` // 预计充电结束时间
	Token        string  `json:"token"`               // 用户token
}

// 创建的订单，存储创建的订单的详细信息
type OrderResponse struct {
	OrderID       int64   `json:"orderID"`       // 订单ID
	Account       string  `json:"account"`       // 用户名
	ChargeStation string  `json:"chargeStation"` // 充电站名称
	CreateTime    string  `json:"createTime"`    // 创建时间
	Kwh           float64 `json:"kwh"`           // 度数
	Price         float64 `json:"price"`         // 金额
	Detail        Detail  `json:"detail"`
}

type Detail struct {
	ChargePileID int64   `json:"chargePileID"`         // 充电桩ID
	ChargeType   string  `json:"chargeType"`           // 充电类型
	FinishTime   *string `json:"finishTime,omitempty"` // 订单结束时间
	ServicePrice int64   `json:"servicePrice"`         // 服务费用
	State        string  `json:"state"`                // 订单状态，完成、进行中、已完成
}

func (OrderApi) OrderCreate(c *gin.Context) {
	var cr OrderRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
}

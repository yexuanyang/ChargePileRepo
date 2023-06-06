package response

// OrderResponse 查询订单回复信息
type OrderResponse struct {
	ApplyKwh     float64 `json:"apply_kwh"`        // 申请度数，可能不是最后充电的度数！
	CarID        string  `json:"car_id"`           // 车辆ID
	ChargeID     string  `json:"charge_id"`        // 调度到的充电桩id，无为空串
	ChargeKwh    float64 `json:"charge_kwh"`       // 实际充电度数，确切的充电度数，由时间算出
	ChargePrice  float64 `json:"charge_price"`     // 充电计价，平均充电计价
	CreateTime   string  `json:"create_time"`      // 创建时间
	DispatchTime string  `json:"dispatch_time"`    // 调度时间，无为xx-xx-xx xx:xx:xx
	Fee          float64 `json:"fee"`              // 订单最终花费
	FinishTime   string  `json:"finish_time"`      // 终止时间，无为xx-xx-xx xx:xx:xx
	FrontCars    int64   `json:"front_cars"`       // 查看前面有多少车，排号号码-当前已进入充电区号码-1
	ID           string  `json:"id"`               // 订单ID
	Mode         string  `json:"mode"`             // 充电模式，F快充，T慢充
	Number       *int64  `json:"number,omitempty"` // 排号号码，类似叫餐服务，订单当日递增
	ServicePrice float64 `json:"service_price"`    // 服务计价，服务计价
	StartTime    string  `json:"start_time"`       // 开始时间，无为xx-xx-xx xx:xx:xx
	State        string  `json:"state"`            // 订单状态，WAITING - 在等待区等待; DISPATCHED - 在充电桩内等待; CHARGING - 充电中; FINISHED - 结束
	UserID       string  `json:"user_id"`          // 用户ID
}

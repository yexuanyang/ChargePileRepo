package response

// ChargePileResponse
type ChargePileResponse struct {
	CarBlocks           []CarBlockResponse `json:"car_blocks,omitempty"`  // 充电车位数组，空格隔开，大小为2，第一个位置和第二个位置
	ID                  string             `json:"id"`                    // 充电桩id，唯一分配
	Mode                string             `json:"mode"`                  // 充电模式，F- 快充 T- 慢充
	Name                string             `json:"name"`                  // 名字
	State               string             `json:"state"`                 // 工作状态，OFF - 关机状态; WORK - 充电状态; REST - 空闲状态; FAULT - 故障状态
	TotalCharge         float64            `json:"total_charge"`          // 总充电总量
	TotalChargeDuration float64            `json:"total_charge_duration"` // 总充电时长，单位为小时
	TotalChargeFee      float64            `json:"total_charge_fee"`      // 总充电费用，单位为RMB
	TotalChargeService  float64            `json:"total_charge_service"`  // 总服务费用，单位为RMB
	TotalChargeTimes    float64            `json:"total_charge_times"`    // 总充电次数
	TotalFee            float64            `json:"total_fee"`             // 总费用，单位为RMB
}

// CarBlockResponse
type CarBlockResponse struct {
	ApplyKwh      float64 `json:"apply_kwh"` // 申请的充电量
	CarID         string  `json:"car_id"`    // 车辆ID
	Name          *string `json:"name,omitempty"`
	Number        string  `json:"number"`         // 此车位的车辆分配的number
	PowerCapacity float64 `json:"power_capacity"` // 总电量
	PowerCurrent  float64 `json:"power_current"`  // 当前电量
	State         string  `json:"state"`          // 车位内的车辆状态，WAITING - 在等待区等待DISPATCHED - 在充电桩内等待CHARGING - 充电中
	UserID        string  `json:"user_id"`        // 用户ID
	WaitTime      float64 `json:"wait_time"`      // 等待时间，单位为小时
}

// StationResponse
type StationResponse struct {
	ChargeArray []ChargePileResponse `json:"charge_array"` // 充电桩数组，充电站下的充电桩
	ID          string               `json:"id"`           // id标识
	Name        *string              `json:"name,omitempty"`
	WaitArray   []CarBlockResponse   `json:"wait_array"` // 等待车位数组，内容是车位
}

type ChargePileReport struct {
	ID                  string  `json:"id"`                    // 充电桩id，唯一分配
	Mode                string  `json:"mode"`                  // 充电模式，F- 快充 T- 慢充
	Name                string  `json:"name"`                  // 名字
	State               string  `json:"state"`                 // 工作状态，OFF - 关机状态; WORK - 充电状态; REST - 空闲状态; FAULT - 故障状态
	Time                string  `json:"time"`                  // 当前时间
	TotalCharge         float64 `json:"total_charge"`          // 总充电总量
	TotalChargeDuration float64 `json:"total_charge_duration"` // 总充电时长，单位为小时
	TotalChargeFee      float64 `json:"total_charge_fee"`      // 总充电费用，单位为RMB
	TotalChargeService  float64 `json:"total_charge_service"`  // 总服务费用，单位为RMB
	TotalChargeTimes    float64 `json:"total_charge_times"`    // 总充电次数
	TotalFee            float64 `json:"total_fee"`             // 总费用，单位为RMB
}

type ChargePileCarInfoResponse struct {
	CarBlocks []CarBlockResponse `json:"car_blocks"` // 充电车位数组，空格隔开，大小为2，第一个位置和第二个位置
	ID        string             `json:"id"`         // 充电桩id，唯一分配
	Mode      string             `json:"mode"`       // 充电模式，F- 快充 T- 慢充
	Name      string             `json:"name"`       // 名字
	State     string             `json:"state"`      // 工作状态，OFF - 关机状态; WORK - 充电状态; REST - 空闲状态; FAULT - 故障状态
}

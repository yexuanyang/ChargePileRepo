package response

// Car
type CarResponse struct {
	CarID         string  `json:"car_id"` // 车辆ID
	Name          *string `json:"name,omitempty"`
	PowerCapacity float64 `json:"power_capacity"`    // 总电量
	PowerCurrent  float64 `json:"power_current"`     // 当前电量
	UserID        *string `json:"user_id,omitempty"` // 用户ID
}

package system

// CarQueueModel 等待区排队队列表
type CarQueueModel struct {
	MODEL
	CarCount int        `json:"car_count"`
	CarModel []CarModel `gorm:"foreignKey:CarQueueID" json:"-"`
}

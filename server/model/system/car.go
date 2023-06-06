// 自动生成模板Car
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Car 结构体
type Car struct {
	global.GVA_MODEL
	UserId          uint    `json:"user_id" form:"user_id" gorm:"column:user_id;comment:;size:20;"`
	User            SysUser `gorm:"foreignKey:UserId;reference:ID" json:"-"`
	CarName         string  `json:"carId" form:"carId" gorm:"type:string;column:car_id;comment:汽车车牌号;"`
	BatteryCapacity float64 `json:"batteryCapacity" form:"batteryCapacity" gorm:"column:battery_capacity;comment:;size:22;"`
	CarBoard        string  `json:"carBoard" form:"carBoard" gorm:"column:car_board;comment:车的品牌型号;"`
}

// CreateCar 用户添加车辆时的请求结构体
type CreateCar struct {
	CarId           string  `json:"carId" form:"carId" gorm:"column:car_id;comment:汽车车牌号;"`
	BatteryCapacity float64 `json:"batteryCapacity" form:"batteryCapacity" gorm:"column:battery_capacity;comment:;size:22;"`
	CarBoard        string  `json:"carBoard" form:"carBoard" gorm:"column:car_board;comment:车的品牌型号;"`
}

// TableName Car 表名
func (Car) TableName() string {
	return "cars"
}

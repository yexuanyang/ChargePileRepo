// 自动生成模板Test
package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Test 结构体
type Test struct {
	global.GVA_MODEL
	TestF string `json:"testF" form:"testF" gorm:"column:test_f;comment:;"`
}

// TableName Test 表名
func (Test) TableName() string {
	return "test"
}

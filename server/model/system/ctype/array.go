package ctype

import (
	"database/sql/driver"
	"strings"
)

type Array []string

// Scan 从数据库中取
func (t *Array) Scan(value interface{}) error {
	data, _ := value.([]byte)
	if string(data) == "" {
		*t = []string{}
		return nil
	}
	*t = strings.Split(string(data), "\n")
	return nil
}

// Value 存入数据库
func (t Array) Value() (driver.Value, error) {
	// 将数字转换为值
	return strings.Join(t, "\n"), nil
}

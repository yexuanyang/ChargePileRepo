package system

import "time"

type MODEL struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"` // 查询的关键字
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

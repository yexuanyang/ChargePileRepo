package system

// UserModel 用户表
type UserModel struct {
	MODEL
	NickName   string      `gorm:"size:36" json:"nick_name"`   // 昵称
	UserName   string      `gorm:"size:36" json:"user_name"`   // 用户名
	Password   string      `gorm:"size:128" json:"-"`          // 密码
	Avatar     string      `gorm:"size:256" json:"avatar"`     // 头像
	Email      string      `gorm:"size:128" json:"email"`      // 邮箱
	Tel        string      `gorm:"size:18" json:"tel"`         // 手机号
	Token      string      `gorm:"size:256" json:"token"`      // token
	CarModel   *CarModel   `gorm:"foreignKey:UserID" json:"-"` // 用户的汽车信息
	OrderModel *OrderModel `gorm:"foreignKey:UserID" json:"-"`
}

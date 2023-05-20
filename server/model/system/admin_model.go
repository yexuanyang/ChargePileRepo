package system

// AdminModel 管理员表
type AdminModel struct {
	MODEL
	NickName string `gorm:"size:36" json:"nick_name"` // 昵称
	UserName string `gorm:"size:36" json:"user_name"` // 用户名
	Password string `gorm:"size:128" json:"-"`        // 密码
}

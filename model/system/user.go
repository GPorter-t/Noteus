package system

import (
	"Noteus/global"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	global.GVA_MODEL
	UUID     uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	Username string    `json:"username" gorm:"comment:用户登录名"`
	Password string    `json:"password" gorm:"comment:用户登录密码"`
	Nickname string    `json:"nickname" gorm:"comment:用户昵称"`
	Phone    string    `json:"phone" gorm:"comment:用户手机号码"`
	Email    string    `json:"email" gorm:"comment:用户邮箱"`
	Enable   int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	WechatId string    `json:"wechat_id" gorm:"comment: 微信来源"`
}

func (User) TableName() string {
	return "sys_users"
}

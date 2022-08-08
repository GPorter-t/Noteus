package system

import (
	"Noteus/global"
	"Noteus/model/system"
	"Noteus/utils"
	"context"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

var ctx = context.Background()

type UserService struct {
}

func (s *UserService) GetCaptcha(username string) string {
	captcha := utils.CreateCaptcha(global.GVA_CONFIG.Captcha.KeyLong)
	timeout := time.Duration(global.GVA_CONFIG.Captcha.TimeOut)
	global.GVA_REDIS.Set(ctx, "system:user:captcha::"+username, captcha, time.Second*timeout)
	return captcha
}

// Register @author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.User
//@return: userInter system.User, err error
func (s *UserService) Register(u system.User) (userInter system.User, err error) {
	var user system.User
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).Or("email = ?", u.Email).Or("wechat_id=?", u.WechatId).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户已注册")
	}
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	err = global.GVA_DB.Create(&u).Error
	return u, err
}

func (s *UserService) Login(u *system.User) (userInter *system.User, sessionId string, err error) {
	if nil == global.GVA_DB {
		return nil, "", fmt.Errorf("db not init")
	}

	var user system.User
	err = global.GVA_DB.Where("username = ?", u.Username).Or("email = ?", u.Email).Or("wechat_id=?", u.WechatId).Find(&user).Error
	if err == nil {
		if u.WechatId == "" {
			if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
				return nil, "", fmt.Errorf("密码错误")
			}
		} else {
			i, e := global.GVA_REDIS.Exists(ctx, "system:user:captcha::"+u.WechatId).Result()
			if e != nil {
				return nil, "", fmt.Errorf("请先获取验证码")
			}
			if i != int64(1) {
				return nil, "", fmt.Errorf("验证码错误")
			}
			global.GVA_REDIS.Del(ctx, "system:user:captcha::"+u.Username)
		}
	}

	sessionId = uuid.NewV4().String()
	return &user, sessionId, nil
}

func (s *UserService) SelectLoginMode(mode int, username, password string) (user *system.User, err error) {
	user = new(system.User)
	if mode == 0 {
		user.Username = username
	} else if mode == 1 {
		user.Email = username
	} else if mode == 2 {
		user.WechatId = username
	}
	user.Password = password
	return
}

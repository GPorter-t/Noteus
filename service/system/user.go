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
	_, err := global.GVA_REDIS.Set(ctx, "system:user:captcha::"+username, captcha, time.Second*timeout).Result()
	if err != nil {
		fmt.Printf(err.Error())
	}
	return captcha
}

func verifyCaptcha(username string, captcha string) (ok bool, err error) {
	c, e := global.GVA_REDIS.Get(ctx, "system:user:captcha::"+username).Result()
	if e != nil {
		global.GVA_LOG.Error("获取失败:" + e.Error())
		return false, fmt.Errorf("获取验证码失败: %v\n", e)
	}
	if c != captcha {
		return false, fmt.Errorf("验证码错误")
	} else {
		global.GVA_REDIS.Del(ctx, "system:user:captcha::"+username)
		return true, nil
	}
}

// Register @author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.User
//@return: userInter system.User, err error
func (s *UserService) Register(u system.User, captcha string) (userInter *system.User, err error) {
	var user system.User
	if !errors.Is(global.GVA_DB.Where("email = ?", u.Email).First(&user).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户已注册")
	}
	if ok, _ := verifyCaptcha(u.Email, captcha); !ok {
		return userInter, errors.New("验证码错误")
	}
	u.Username = u.Email
	if u.Password != "" {
		u.Password = utils.BcryptHash(u.Password)
	}
	u.UserId = uuid.NewV4()
	err = global.GVA_DB.Create(&u).Error
	return &u, err
}

type loginMode func(username, password string) (user *system.User, err error)

func (s *UserService) Login(username, password string, mode int) (userInter *system.User, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}
	var login loginMode
	switch mode {
	case 0:
		login = loginWithPassword
	case 1:
		login = loginWithCaptcha
	}
	user, err := login(username, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func loginWithCaptcha(username, captcha string) (user *system.User, err error) {
	var u *system.User
	err = global.GVA_DB.Where("email = ?", username).Or("username=?", username).Find(&u).Error
	if err == nil {
		ok, e := verifyCaptcha(username, captcha)
		if ok {
			return u, nil
		} else {
			return nil, e
		}
	} else {
		global.GVA_LOG.Error("数据库查询错误：" + err.Error())
	}
	return
}

func loginWithPassword(username, password string) (user *system.User, err error) {
	var u *system.User
	err = global.GVA_DB.Where("email = ?", username).Or("username=?", username).Find(&u).Error
	if err == nil {
		if ok := utils.BcryptCheck(password, u.Password); !ok {
			return nil, fmt.Errorf("密码错误")
		}
		return u, nil
	}
	global.GVA_LOG.Error("数据库查询错误：" + err.Error())
	return nil, err
}

func (s *UserService) GetItem(username string) (user *system.User, err error) {
	err = global.GVA_DB.Where("email = ?", username).Or("username=?", username).Find(&user).Error
	if err != nil {
		global.GVA_LOG.Error("查询数据错误：" + err.Error())
	}
	return
}

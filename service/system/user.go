package system

import (
	"Noteus/global"
	"Noteus/utils"
	"context"
	goUuid "github.com/satori/go.uuid"
	"time"
)

var ctx = context.Background()

type UserService struct {
}

func (s *UserService) GetCaptcha() string {
	captcha := utils.CreateCaptcha(global.GVA_CONFIG.Captcha.KeyLong)
	timeout := time.Duration(global.GVA_CONFIG.Captcha.TimeOut)
	global.GVA_REDIS.Set(ctx, "system:user:captcha::"+captcha, captcha, time.Second*timeout)
	return captcha
}

func (s *UserService) VerifyCaptcha(captcha string) (ok bool, err error) {
	i, e := global.GVA_REDIS.Exists(ctx, "system:user:captcha::"+captcha).Result()
	if e != nil {
		return
	}
	if i != int64(1) {
		return
	}
	global.GVA_REDIS.Del(ctx, "system:user:captcha::"+captcha)
	session_id := goUuid.NewV4().String()
	global.GVA_REDIS.Set(ctx, "system:user:session_id::"+session_id, session_id, time.Second*60*60*24)
	ok = true
	return
}

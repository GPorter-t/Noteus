package system

import (
	"Noteus/global"
	"Noteus/model/common/response"
	"Noteus/model/system"
	systemReq "Noteus/model/system/request"
	systemRsp "Noteus/model/system/response"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strings"
)

var ctx = context.Background()

type SystemApi struct{}

func (s *SystemApi) GetCaptcha(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		response.FailWithMessage("username is empty", c)
		return
	}
	captcha := systemService.GetCaptcha(username)

	targetUrl := fmt.Sprintf("http://%s:%d/email/sendEmail", global.GVA_CONFIG.System.Host, global.GVA_CONFIG.System.Port)

	email := &systemReq.Email{
		To:      username,
		Subject: "验证码",
		Body:    fmt.Sprintf("尊敬的%s:\n\t您好，您的验证码为：%s", username, captcha),
	}
	payload, err := json.Marshal(email)
	if err != nil {
		global.GVA_LOG.Error("验证码发送失败:" + err.Error())
		response.FailWithMessage("验证码发送失败，请稍后重试", c)
		return
	}
	body := strings.NewReader(string(payload))
	_, err = http.Post(targetUrl, "application/json", body)
	if err != nil {
		global.GVA_LOG.Error("验证码发送失败:" + err.Error())
		response.FailWithMessage("验证码发送失败，请稍后重试", c)
		return
	}

	response.OkWithData("验证码已发送至指定邮箱", c)
}

func (s *SystemApi) Login(c *gin.Context) {
	var req systemReq.UserLoginReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if req.Password == "" {
		response.FailWithMessage("暂未设置密码，请使用验证码进行登录", c)
		return
	}

	user, err := systemService.Login(req.Username, req.Password, req.Mode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sessionId := uuid.NewV4().String()
	global.GVA_REDIS.Set(ctx, "system:user:session_id::"+sessionId, user.Username, 0)
	response.OkWithData(systemRsp.UserRsp{
		User:      user,
		SessionId: sessionId,
	}, c)
}

func (s *SystemApi) Register(c *gin.Context) {
	var req systemReq.UserRegisterReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user, err := systemService.Register(system.User{
		Email:    req.Email,
		Password: req.Password,
	}, req.Captcha)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(user, c)
}

func (s *SystemApi) GetInfo(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		sessionId, _ := c.Cookie("session_id")
		var err error
		username, err = global.GVA_REDIS.Get(ctx, "system:user:session_id::"+sessionId).Result()
		if err != nil {
			global.GVA_LOG.Error("redis 查询失败" + err.Error())
			response.FailWithMessage("查询失败", c)
			return
		}
	}
	user, err := systemService.GetItem(username)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(user, c)
}

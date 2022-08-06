package utils

import (
	"math/rand"
	"time"
)

const (
	characterBytes  = "!@#$%^&*?"
	digitBytes      = "1234567890"
	lowLetterBytes  = "abcdefghijklmnopqrstuvwxyz"
	highLetterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	keyQuery        = characterBytes + digitBytes + lowLetterBytes + highLetterBytes
)

func CreateCaptcha(length int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	captcha := ""
	for i := 0; i < length; i++ {
		index := r.Intn(len(keyQuery))
		captcha += string(keyQuery[index])
	}
	return captcha
}

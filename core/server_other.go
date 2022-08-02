//go:build !windows
// +build !windows

package core

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func initServer(address string, rounter *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderSize = 1 << 20
}

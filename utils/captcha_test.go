package utils

import (
	"fmt"
	"testing"
)

func TestCreateCaptcha(t *testing.T) {
	r := CreateCaptcha(6)
	fmt.Println(r)
}

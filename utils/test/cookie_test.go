package test

import (
	"GoBlog/utils"
	"fmt"
	"testing"
)

func TestGenerateCookies(t *testing.T) {
	cookie := utils.GenerateCookies(20)
	fmt.Println(cookie)
}

package test

import (
	"GoBlog/utils"
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	ciphertext := utils.Md5("123456")
	if ciphertext != "e10adc3949ba59abbe56e057f20f883e" {
		fmt.Println("加密结果不一致")
		t.Fail()
	} else {
		fmt.Println(ciphertext)
	}
}

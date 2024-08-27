package test

import (
	"GoBlog/utils"
	"fmt"
	"testing"
	"time"
)

func TestGenerateJWT(t *testing.T) {
	payload := utils.JwtPayload{Sub: "auth", Name: "jwt", Iat: int(time.Now().Unix())}
	jwt, err := utils.GenerateJWT(utils.DefaultHeader, payload, utils.Secret)
	if err != nil {
		fmt.Println("生成JWT失败:", err)
		t.Fail()
	} else {
		fmt.Println(jwt)
	}
}

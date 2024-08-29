package test

import (
	"GoBlog/utils"
	"fmt"
	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	payload := utils.JwtPayload{Sub: "auth", Uid: 1, Iat: int(time.Now().Unix())}
	jwt, err := utils.GenerateJWT(utils.DefaultHeader, payload, utils.Secret)
	if err != nil {
		fmt.Println("生成JWT失败:", err)
		t.Fail()
	} else {
		fmt.Println(jwt)
		if _, p, e := utils.VerifyJWT(jwt, utils.Secret); e != nil { //验证
			fmt.Println("解析JWT失败", err)
		} else {
			fmt.Println("解析过后的uid:", p.Uid)
			if p.Uid != payload.Uid { //用解析过后的uid与解析之前的uid进行对比
				fmt.Println("两个uid不一致")
				t.Fail()
			}
		}
	}
}

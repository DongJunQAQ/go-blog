package test

import (
	"GoBlog/db"
	"fmt"
	"testing"
)

func TestSetCookie(t *testing.T) {
	cookieValue := "qweqweq1122"
	uid := 10
	db.SetCookieToRedis(cookieValue, uint(uid))
}
func TestGetCookie(t *testing.T) {
	cookieValue := "auth_cookie_qwe11"
	uid := db.GetCookieFromRedis(cookieValue)
	fmt.Println(uid)
}

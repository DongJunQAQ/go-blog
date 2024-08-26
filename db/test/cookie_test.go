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
	cookieValue := "iY8H5iF5F0s4NG3h6Y54Lp3K"
	uid := db.GetCookieFromRedis(cookieValue)
	fmt.Println(uid)
}

package test

import (
	"GoBlog/db"
	"GoBlog/utils"
	"testing"
)

func init() {
	utils.InitLog("log")
}
func TestGetBlogById(t *testing.T) {
	blog := db.GetBlogById(1)
	if blog == nil {
		t.Fail()
		return
	}
}
func TestBlogByUserIdList(t *testing.T) {
	blogs := db.GetBlogByUserIdList(1)
	if blogs == nil {
		t.Fail()
		return
	}
}

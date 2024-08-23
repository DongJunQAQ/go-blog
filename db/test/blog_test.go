package test

import (
	"GoBlog/db"
	"GoBlog/models"
	"GoBlog/utils"
	"fmt"
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
func TestUpdateBlog(t *testing.T) {
	blog := models.Blog{ID: 1, Title: "新标题123", Article: "新正文456"}
	if err := db.UpdateBlog(&blog); err != nil {
		fmt.Println(err)
		t.Fail()
	}
}

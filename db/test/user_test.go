package test

import (
	"GoBlog/db"
	"GoBlog/utils"
	"testing"
)

func init() {
	utils.InitLog("log")
}
func TestGetUserByName(t *testing.T) {
	user := db.GetUserByName("dongjun")
	if user == nil {
		t.Fail()
		return
	}
	if user.Password != "123456" {
		t.Fail()
		return
	}
	user = db.GetUserByName("none")
	if user != nil {
		t.Fail()
		return
	}
}
func TestCreateUser(t *testing.T) {
	db.CreateUser("董军", "7900802")
}
func TestDeleteUserByName(t *testing.T) {
	db.DeleteUserByName("董军")
}

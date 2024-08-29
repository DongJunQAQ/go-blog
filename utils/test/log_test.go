package test

import (
	"GoBlog/utils"
	"testing"
)

func TestLog(t *testing.T) {
	utils.InitLog("log")
	utils.LogRus.Debug("这是一条debug日志")
	utils.LogRus.Info("这是一条info日志")
	utils.LogRus.Warn("这是一条warn日志")
	utils.LogRus.Error("这是一条error日志")
	//utils.LogRus.Panic("这是一条panic日志")
}

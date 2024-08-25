package utils

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs" //日志的输出
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
	"time"
)

var LogRus *logrus.Logger

func InitLog(configFile string) {
	logConf := ReadConfig(configFile)
	LogRus = logrus.New()
	switch strings.ToLower(logConf.GetString("level")) { //根据配置文件获取当前日志级别
	case "debug":
		LogRus.SetLevel(logrus.DebugLevel) //设置日志的显示级别
	case "info":
		LogRus.SetLevel(logrus.InfoLevel)
	case "warn":
		LogRus.SetLevel(logrus.WarnLevel)
	case "error":
		LogRus.SetLevel(logrus.ErrorLevel)
	case "panic":
		LogRus.SetLevel(logrus.PanicLevel)
	default:
		panic(fmt.Errorf("无效的日志级别:%s", viper.GetString("level")))
	}
	LogRus.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05.000"}) //日志显示时间的格式，显示毫秒
	logPath := ProjectRootPath + logConf.GetString("logPath")                              //设置存放日志的路径
	logOut, err := rotatelogs.New(                                                         //设置输出日志的相关选项
		logPath+".%Y%m%d",                         //指定日志文件的路径和名称后缀
		rotatelogs.WithLinkName(logPath),          //为最新的一份日志创建软链接
		rotatelogs.WithRotationTime(24*time.Hour), //每隔1天生成一份新的日志
		rotatelogs.WithMaxAge(7*24*time.Hour))     //只保留最近7天的日志
	if err != nil {
		panic(err)
	}
	LogRus.SetOutput(logOut)     //设置日志文件的输出
	LogRus.SetReportCaller(true) //输出是从哪个文件的哪行代码生成的日志，快速找到问题点
}

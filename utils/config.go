package utils

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

var ProjectRootPath = path.Dir(getCurrentPath() + "/../") //获取当前函数路径后再往前推一级以获取项目根路径

func getCurrentPath() string { //获取当前函数的文件夹路径
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(filename) //通过文件名来获取文件夹路径
}
func ReadConfig(file string) *viper.Viper { //解析配置文件
	config := viper.New()
	configPath := ProjectRootPath + "/config/"
	config.AddConfigPath(configPath)              //配置文件所在目录
	config.SetConfigName(file)                    //文件名
	config.SetConfigType("yaml")                  //文件后缀
	configFile := configPath + file + ".yaml"     //配置文件的完整路径
	if err := config.ReadInConfig(); err != nil { //读取配置文件，先定义一个局部变量err，再来判断有无错误
		var configFileNotFoundError viper.ConfigFileNotFoundError //该类型用于捕获配置文件未找到的特定错误类型
		if errors.As(err, &configFileNotFoundError) {             //检查错误是否是viper.ConfigFileNotFoundError类型
			panic(fmt.Errorf("找不到配置文件: %s", configFile)) //如果发生错误则直接终止掉程序，使用fmt.Errorf()来返回错误并格式化报错信息，errors.New()无法格式化报错信息
		} else {
			panic(fmt.Errorf("解析配置文件%s时出错: %s", configFile, err))
		}
	}
	return config
}

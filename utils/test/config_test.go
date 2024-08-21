package test

import (
	"GoBlog/utils"
	"fmt"
	"path"
	"runtime"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	//1.通过viper的各种get方法来读取配置:
	dbViper := utils.ReadConfig("mysql")
	dbViper.WatchConfig()                   //在程序运行期间修改配置内容viper会自动获取到新配置
	host := dbViper.GetString("mysql.host") //如果要读取配置的值不存在则返回其应类型的空值
	fmt.Println("MySQL主机:", host)
	if !dbViper.IsSet("mysql.port") { //如果配置中没有设置mysql.port的值则单元测试视为失败
		t.Fatal()
	}
	port := dbViper.GetInt("mysql.port")
	fmt.Println("MySQL端口:", port)
	time.Sleep(10 * time.Second)
	port = dbViper.GetInt("mysql.port")
	fmt.Println("过了10秒后的MySQL端口:", port)
	//2.通过对结构体的反序列化来读取配置:
	logViper := utils.ReadConfig("log")
	type LogConfig struct {
		Level string `mapstructure:"level"`
		File  string `mapstructure:"file"`
	}
	var config LogConfig
	if err := logViper.Unmarshal(&config); err != nil { //反序列化，将配置内容转换为结构体，想要在函数内部修改变量的值就需要使用指针来修改
		fmt.Println(err)
		t.Fail()
	} else {
		fmt.Println("日志级别:", config.Level)
		fmt.Println("日志文件路径:", config.File)
	}
}
func af() (string, int) {
	_, filename, line, _ := runtime.Caller(0) //返回当前函数的调用信息，0表示当前函数本身，如果值为1则表示第2次调用该函数，以此类推
	return filename, line
}
func bf() (string, int) {
	return af() //如果上面的值的为1则返回此行行号
}
func TestCaller(t *testing.T) {
	filename, line := bf()                                //如果上面的值的为2则返回此行行号
	fmt.Println(filename, line)                           //打印该函数所在的文件名和行号
	fmt.Println(path.Dir(path.Dir(filename) + "/../../")) //获取项目根目录
}
func TestProjectRootPath(t *testing.T) {
	fmt.Println(utils.ProjectRootPath)
}

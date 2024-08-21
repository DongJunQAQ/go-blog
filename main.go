package main

import (
	"GoBlog/utils"
	"fmt"
)

func main() {
	conf := utils.CreateConfig("mysql")
	fmt.Println(conf.GetString("mysql.port"))
}

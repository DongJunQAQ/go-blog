package test

import (
	"GoBlog/db"
	"sync"
	"testing"
)

func TestConnectMySQL(t *testing.T) { //使用5个协程去连接MySQL，查看是否只创建了一个连接池
	const N = 5
	wg := sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() { //创建一个子协程来运行函数
			defer wg.Done()   //当函数内的所有代码执行完毕后再执行此行，表示该子协程运行完毕
			db.ConnectMySQL() //不管循环多少次我只会创建一个MySQL连接池
		}()
	}
	wg.Wait() //在主协程中等这几个子协程结束
}

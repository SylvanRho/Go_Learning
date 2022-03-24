package main

import (
	"fmt"
	"time"
)

func main() {
	//同步改成异步就需要一个go就可以了
	go mygo("协程1号")
	go mygo("协程2号")
	time.Sleep(time.Second)
}

func mygo(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("index： %d goruntinue %s \n", i, name)
		// 为了避免第一个协程执行过快，观察不到并发的效果，加个休眠
		time.Sleep(10 * time.Millisecond)
	}
}

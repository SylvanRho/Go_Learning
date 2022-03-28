package main

import (
	"fmt"
	"time"
)

func main() {

	//注意⚠️：recover，并不能任意使用，它有强制要求，必须得在 defer 下才能发挥用途。

	captionPanic()
	acrossGoroutineDefer()
}

//触发宕机
func triggerPanic() {
	//只要手动调用panic就可以了
	panic("carsh")
}

//捕获panic
func captionPanic() {
	fmt.Printf("**********captionPanic**********\n")
	set_data(20)

	// 如果能执行到这句，说明panic被捕获了
	// 后续的程序能继续运行
	fmt.Println("everything is ok")
}

//用于测试捕获panic的测试方法
func set_data(x int) {
	defer func() {
		// recover() 可以将捕获到的panic信息打印
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// 故意制造数组越界，触发 panic
	var arr [10]int
	// 即使 panic 会导致整个程序退出，但在退出前，若有 defer 延迟函数，还是得执行完 defer
	arr[x] = 15
}

//跨携程的延迟语句
func acrossGoroutineDefer() {
	fmt.Printf("**********acrossGoroutineDefer**********\n")

	defer fmt.Println("in acrossGoroutineDefer")

	go func() {
		fmt.Println("in gorouttine")
		//defer 在多个协程之间是没有效果，在子协程里触发 panic，只能触发自己协程内的 defer，而不能调用 main 协程里的 defer 函数的。
		//也就是说 只会输出 in gorouttine 而不会输出 in acrossGoroutineDefer
		panic("")
	}()

	time.Sleep(2 * time.Second)
}

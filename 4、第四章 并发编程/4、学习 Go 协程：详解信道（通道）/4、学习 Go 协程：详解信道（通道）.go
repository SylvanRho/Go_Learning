package main

import (
	"fmt"
	"time"
)

func main() {
	//几个注意事项⚠️
	// 关闭一个未初始化的 channel 会产生 panic
	// 重复关闭同一个 channel 会产生 panic
	// 向一个已关闭的 channel 发送消息会产生 panic
	// 从已关闭的 channel 读取消息不会产生 panic，且能读出 channel 中还未被读取的消息，若消息均已被读取，则会读取到该类型的零值。
	// 从已关闭的 channel 读取消息永远不会阻塞，并且会返回一个为 false 的值，用以判断该 channel 是否已关闭（x,ok := <- ch）
	// 关闭 channel 会产生一个广播机制，所有向 channel 读取消息的 goroutine 都会收到消息
	// channel 在 Golang 中是一等公民，它是线程安全的，面对并发问题，应首先想到 channel。

	// 值类型 ：String，Array，Int，Struct，Float，Bool
	// 引用类型：Slice，Map

	createChannelDemo()
	channelCapAndLenDemo()
	bufferOrNoBufferChannelDemo()
	twoWayChannelDemo()
	oneWayChannelDemo()
	throughTheChannelDemo()
	useChannelsForLocking()

}

//创建Channel 的Demo
func createChannelDemo() {
	fmt.Printf("**********createChannelDemo**********\n")
	/* 第一种 声明一个类型后创建
	var 信道实例 chan 信道类型
	声明后的信道，其零值是nil，无法直接使用，必须配合make函进行初始化。
	信道实例 = make(chan 信道类型)
	*/

	/* 第二种 直接创建
	信道实例 := make(chan 信道类型)
	*/

	pipline := make(chan int, 1)

	// 往信道中发送数据
	pipline <- 200

	// 从信道中取出数据，并赋值给mydata
	mydata := <-pipline
	fmt.Printf("mydata：%v \n", mydata)

	pipline <- 150

	//假如通道里没有数据，再去取要抛 panic
	//当从信道中读取数据时，可以有多个返回值，其中第二个可以表示 信道是否被关闭，如果已经被关闭，ok 为 false，若还没被关闭，ok 为true。
	x, ok := <-pipline
	if ok {
		fmt.Printf("x：%v \n", x)
	}
	//使用完后要关闭
	close(pipline)
}

//信道的容器以及长度Demo
func channelCapAndLenDemo() {
	fmt.Printf("**********channelCapAndLenDemo**********\n")
	//一般创建信道都是使用 make 函数，make 函数接收两个参数
	// 	第一个参数：必填，指定信道类型
	//  第二个参数：选填，不填默认为0，指定信道的容量（可缓存多少数据）
	/*
		对于信道的容量，很重要，这里要多说几点：
			当容量为0时，说明信道中不能存放数据，在发送数据时，必须要求立马有人接收，否则会报错。此时的信道称之为无缓冲信道。
			当容量为1时，说明信道只能缓存一个数据，若信道中已有一个数据，此时再往里发送数据，会造成程序阻塞。 利用这点可以利用信道来做锁。
			当容量大于1时，信道中可以存放多个数据，可以用于多个协程之间的通信管道，共享资源。
	*/

	pipline := make(chan int, 15)
	fmt.Printf("信道中可缓冲 %d 个数据 \n", cap(pipline))
	pipline <- 1
	fmt.Printf("信道中当前有 %d 个数据 \n", len(pipline))
}

//缓冲或无缓冲Demo
func bufferOrNoBufferChannelDemo() {
	fmt.Printf("**********bufferOrNoBufferChannelDemo**********\n")
	//缓冲信道
	//允许信道里存储一个或多个数据，这意味着，设置了缓冲区后，发送端和接收端可以处于异步的状态。
	bufferPipline := make(chan int, 10)
	bufferPipline <- 15

	//无缓冲信道
	//在信道里无法存储数据，这意味着，接收端必须先于发送端准备好，以确保你发送完数据后，有人立马接收数据，否则发送端就会造成阻塞\
	noBufferPipline := make(chan int)
	fmt.Println(noBufferPipline)

	// 或者
	// noBufferPipline := make(chan int, 0)
	// fmt.Println(noBufferPipline)
	close(bufferPipline)
}

//双向通道Demo
func twoWayChannelDemo() {
	fmt.Printf("**********twoWayChannelDemo**********\n")

	//默认就是双向的
	pipline := make(chan int)

	go func() {
		fmt.Println("准备发送数据: 100")
		pipline <- 100
	}()

	go func() {
		num := <-pipline
		fmt.Printf("接收到的数据是: %d \n", num)
	}()
	// 主函数sleep，使得上面两个goroutine有机会执行
	time.Sleep(time.Second)
}

//单向通道Demo
func oneWayChannelDemo() {
	fmt.Printf("**********oneWayChannelDemo**********\n")

	//接收
	type Recover = <-chan int

	//发送
	type Sender = chan<- int

	pipline := make(chan int)

	go func() {
		var sender Sender = pipline
		fmt.Println("准备发送数据：100")
		sender <- 100
	}()

	go func() {
		var recover Recover = pipline
		fmt.Println("准备接收数据")
		num := <-recover
		fmt.Println("接收到的数据num为：", num)
	}()

	// 主函数sleep，使得上面两个goroutine有机会执行
	time.Sleep(time.Second)
}

//遍历通道Demo
func throughTheChannelDemo() {
	fmt.Printf("**********throughTheChannelDemo**********\n")

	pipline := make(chan int, 10)
	fibonacci(pipline)
	for k := range pipline {
		fmt.Println(k)
	}
}

//斐波那契数列
func fibonacci(ch chan int) {
	n := cap(ch)

	x, y := 1, 1

	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}

	// 记得 close 信道
	// 不然主函数中遍历完并不会结束，而是会阻塞。
	close(ch)
}

//用信道来做锁
func useChannelsForLocking() {
	fmt.Printf("**********useChannelsForLocking**********\n")

	// 注意要设置容量为 1 的缓冲信道
	pipling := make(chan bool, 1)
	var num int
	for i := 0; i < 1000; i++ {
		go increment(pipling, &num)
	}

	// 确保所有的协程都已完成
	// 以后会介绍一种更合适的方法（Mutex），这里暂时使用sleep
	time.Sleep(time.Second)
	fmt.Println("num is ：", num)
}

// 由于 x=x+1 不是原子操作
// 所以应避免多个协程对x进行操作
// 使用容量为1的信道可以达到锁的效果
// 增量
func increment(ch chan bool, x *int) {
	ch <- true
	*x = *x + 1
	<-ch
}

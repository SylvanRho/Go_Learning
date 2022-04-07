package main

import (
	"fmt"
	"sync"
)

func main() {
	noBufferDeadLockDemo()
	outOfChanCapErrorDemo()
	notCloseChannelDeadLockDemo()
}

func printChanContent(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(<-ch)
}

//无缓冲死锁Demo
func noBufferDeadLockDemo() {
	fmt.Printf("**********noBufferDeadLockDemo**********\n")
	/*	错误示范
		pipline := make(chan string)
		pipline <- "hello world"
		fmt.Println(<-pipline)
		//对于无缓冲信道，在接收者未准备好之前，发送操作是阻塞的.
		//假如在同一协程中换成先接收再发送还是有问题👇
		fmt.Println(<-pipline)
		pipline <- "hello world"
		//由于前面接收者一直在等待数据 而处于阻塞状态，所以无法执行到后面的发送数据。
	*/
	//第一种办法
	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(1)
	go printChanContent(ch, &wg)
	ch <- "hello"
	wg.Wait()

	//第二种办法
	//改成缓冲通道
	//pipline := make(chan string, 1)
	close(ch)
}

//超出通道容器大小错误
func outOfChanCapErrorDemo() {
	fmt.Printf("**********outOfChanCapErrorDemo**********\n")
	ch := make(chan string, 1)

	ch <- "hello1"

	//信道容量为 1，但是往信道中写入两条数据，对于一个协程来说就会造成死锁。
	// ch <- "hello2"
	fmt.Println(<-ch)
	close(ch)
}

//没有关闭信道死锁
func notCloseChannelDeadLockDemo() {
	fmt.Printf("**********notCloseChannelDeadLockDemo**********\n")
	pipline := make(chan string)

	go func() {
		pipline <- "hello world"
		pipline <- "hello china"
		//假如这一句注释掉了 下面的for会一直等待 形成死锁
		close(pipline)
	}()

	for data := range pipline {
		fmt.Println(data)
	}
}

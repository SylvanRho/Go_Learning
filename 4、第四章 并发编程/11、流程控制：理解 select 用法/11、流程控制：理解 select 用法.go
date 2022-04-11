package main

import (
	"fmt"
	"time"
)

func main() {
	/*	select-case的定义结构
		注意⚠️：仅能用于 信道/通道 的相关操作
		select {
		case 表达式1:
			<code>
		case 表达式2:
			<code>
		default:
			<code>
		}
	*/
	simpleSelectDemo()

	timeOutSelectDemo()
	writeAndReadSelectDemo()
	checkCloseSelectDemo()
}

//简单打的Select使用Demo
func simpleSelectDemo() {
	fmt.Printf("**********simpleSelectDemo**********\n")
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	c2 <- "hello"

	select {
	case msg1 := <-c1:
		fmt.Println("msg1：", msg1)
	case msg2 := <-c2:
		fmt.Println("msg2：", msg2)
	default:
		fmt.Println("No data received.")
	}
}

func makeTimeout(ch chan bool, t int) {
	time.Sleep(time.Duration(t) * time.Second)
	ch <- true
}

//超时的SelectDemo
func timeOutSelectDemo() {
	fmt.Printf("**********timeOutSelectDemo**********\n")
	//当 case 里的信道始终没有接收到数据时，而且也没有 default 语句时，select 整体就会阻塞
	//但是有时我们并不希望 select 一直阻塞下去，这时候就可以手动设置一个超时时间。
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	timeOutCh := make(chan bool, 1)

	go makeTimeout(timeOutCh, 2)

	select {
	case msg1 := <-c1:
		fmt.Println("msg1：", msg1)
	case msg2 := <-c2:
		fmt.Println("msg2：", msg2)
	case <-timeOutCh:
		fmt.Println("Timeout, exit.")
	}
}

//写跟读的Select
func writeAndReadSelectDemo() {
	fmt.Printf("**********timeOutSelectDemo**********\n")
	//select 里的 case 表达式只要求你是对信道的操作即可，不管你是往信道写入数据，还是从信道读出数据。
	c1 := make(chan int, 2)

	c1 <- 2
	select {
	case c1 <- 4:
		fmt.Println("c1 received: ", <-c1)
		fmt.Println("c1 received: ", <-c1)
	default:
		fmt.Println("channel blocking")
	}
}

//检测关闭信道的Select
func checkCloseSelectDemo(){
	fmt.Printf("**********checkCloseSelectDemo**********\n")
	c1 := make(chan int, 1)
    c2 := make(chan int, 1)
    close(c1)
    for {
        select {
        case <-c1:
            fmt.Println("stop");
			return
        case <-c2:
            fmt.Println("hhh")
        }
    }
}

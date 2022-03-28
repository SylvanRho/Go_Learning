package main

import (
	"fmt"
	"sync"
)

func main() {
	usedChannelMarkCompleted()
	usedWaitGroup()
}

//使用通道来标记完成
func usedChannelMarkCompleted() {
	fmt.Printf("**********usedChannelMarkCompleted**********\n")

	done := make(chan bool, 1)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		done <- true
	}()
	//通道接收到消息 就是完成了
	<-done

	fmt.Println("协程完成啦")
}

//使用waitGroup
func usedWaitGroup() {
	fmt.Printf("**********usedWaitGroup**********\n")

	//实例化语法👇
	//var 实例名 sync.WaitGroup
	var wg sync.WaitGroup

	//实例化完成后，就可以使用它的几个方法：
	// Add：初始值为0，你传入的值会往计数器上加，这里直接传入你子协程的数量
	// Done：当某个子协程完成后，可调用此方法，会从计数器上减一，通常可以使用 defer 来调用。
	// Wait：阻塞当前协程，直到实例里的计数器归零。

	wg.Add(2)
	go worker(1, wg)
	go worker(2, wg)

	wg.Wait()
}

func worker(x int, wg sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("worker %d: %d\n", x, i)
	}
}

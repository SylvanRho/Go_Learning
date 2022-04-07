package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	/*	⚠️万能的通道编程模型--规范
		当只有一个发送者时，无论有多少接收者，业务通道都应由唯一发送者关闭。
		当有多个发送者，一个接收者时，应借助管理通道，由业务通道唯一接收者充当管理通道的发送者，其他业务通道的发送者充当接收者
		当有多个发送者，多个接收者时，这是最复杂的，不仅要管理通道，还要另起一个专门的媒介协程，新增一个媒介通道，但核心逻辑都是一样。
	*/
	multipleSenderSingleReceiver()
	multipleSenderMultipleReceiver()
}

//多个发送者 单个接收者
func multipleSenderSingleReceiver() {
	fmt.Printf("**********multipleSenderSingleReceiver**********\n")
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	const Max = 100000
	const NumSender = 1000

	wg := sync.WaitGroup{}
	wg.Add(1)

	// 业务通道
	dataCh := make(chan int)

	// 管理通道：必须是无缓冲通道
	// 其发送者是 业务通道的接收者。
	// 其接收者是 业务通道的发送者。
	stopCh := make(chan struct{})

	// 业务通道的发送者
	for i := 0; i < NumSender; i++ {
		go func() {
			for {
				// 提前检查管理通道是否关闭
				// 让业务通道发送者早尽量退出
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}

	// 业务通道的接收者，亦充当管理通道的发送者
	go func() {
		defer wg.Done()

		for value := range dataCh {
			if value == 6666 {
				fmt.Println("value : ", value, " I'm going to close")
				// 当达到某个条件时
				// 通过关闭管理通道来广播给所有业务通道的发送者
				close(stopCh)
				return
			}
			//打印出来接收到的数据
			// fmt.Println(value)
		}
	}()

	wg.Wait()
}

//多个发送者 多个接收者
func multipleSenderMultipleReceiver() {
	fmt.Printf("**********multipleSenderSingleReceiver**********\n")
	rand.Seed(time.Now().UnixNano())

	const Max = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	wg := sync.WaitGroup{}
	wg.Add(NumReceivers)

	// 1. 业务通道
	dataCh := make(chan int)

	// 2. 管理通道：必须是无缓冲通道
	// 其发送者是：额外启动的管理协程
	// 其接收者是：所有业务通道的发送者。
	stopCh := make(chan struct{})

	// 3. 媒介通道：必须是缓冲通道
	// 其发送者是：业务通道的所有发送者和接收者
	// 其接收者是：媒介协程（唯一）
	toStop := make(chan string, 1)

	var stoppedBy string

	// 媒介协程
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	// 业务通道发送者
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				// 提前检查管理通道是否关闭
				// 让业务通道发送者早尽量退出
				select {
				case <-stopCh:
					return
				default:
				}

				value := rand.Intn(Max)
				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// 业务通道的接收者
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wg.Done()

			for {
				// 提前检查管理通道是否关闭
				// 让业务通道接收者早尽量退出
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					// 一旦满足某个条件，就通过媒介通道发消息给媒介协程
					// 以关闭管理通道的形式，广播给所有业务通道的协程退出
					if value == 6666 {
						// 务必使用 select，两个目的：
						// 1、防止协程阻塞
						// 2、防止向已关闭的通道发送数据导致panic
						select {
						case toStop <- "接收者#" + id:
						default:
						}
						return
					}

				}
			}
		}(strconv.Itoa(i))
	}

	wg.Wait()
	fmt.Println("被" + stoppedBy + "终止了")
}

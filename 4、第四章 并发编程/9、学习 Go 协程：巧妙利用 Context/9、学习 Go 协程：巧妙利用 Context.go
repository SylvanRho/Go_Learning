package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	/*
		Context接口定义如下👇：

		type Context interface {
			Deadline() (deadline time.Time, ok bool)
			Done() <-chan struct{}
			Err() error
			Value(key interface{}) interface{}
		}

		Deadline：
			返回的第一个值是 截止时间，到了这个时间点，Context 会自动触发 Cancel 动作。
			返回的第二个值是 一个布尔值，true 表示设置了截止时间，false 表示没有设置截止时间，如果没有设置截止时间，就要手动调用 cancel 函数取消 Context。
		Done：
			返回一个只读的通道（只有在被cancel后才会返回），类型为 struct{}。当这个通道可读时，意味着parent context已经发起了取消请求，根据这个信号，开发者就可以做一些清理动作，退出goroutine。
		Err：
			返回 context 被 cancel 的原因。
		Value：
			返回被绑定到 Context 的值，是一个键值对，所以要通过一个Key才可以获取对应的值，这个值一般是线程安全的。
	*/

	simpleContextUsedDemo()
	contextWithDeadlineDemo()
	contextTimeoutDemo()
	contextWithValueDemo()
}

//简单上下文的使用Demo
func simpleContextUsedDemo() {
	fmt.Printf("**********simpleContextUsedDemo**********\n")
	//这是返回可以执行取消函数的方法
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go monitor(ctx, i)
	}

	time.Sleep(1 * time.Second)
	// 关闭所有 goroutine
	cancel()

	// 等待5s，若此时屏幕没有输出 <正在监控中> 就说明所有的goroutine都已经关闭
	time.Sleep(5 * time.Second)
	fmt.Println("主程序退出！")
}

//监控是否取消了
func monitor(ctx context.Context, number int) {
	for {
		select {
		// 其实可以写成 case <- ctx.Done()
		// 这里仅是为了让你看到 Done 返回的内容
		//就是一个空结构的数据
		case v := <-ctx.Done():
			fmt.Printf("监控器%v，接收到通道值为：%v，监控结束。\n", number, v)
			return
		default:
			fmt.Printf("监控器%v，正在监控中...\n", number)
			time.Sleep(2 * time.Second)
		}
	}
}

//WithDeadline使用Demo
func contextWithDeadlineDemo() {
	fmt.Printf("**********simpleContextUsedDemo**********\n")

	//创建一个子节点的context,当前时间1秒后自动取消
	//WithDeadline 传入的第二个参数是 time.Time 类型，它是一个绝对的时间，意思是在什么时间点超时取消。
	ctx01, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))

	defer cancel()

	for i := 0; i < 5; i++ {
		go monitor(ctx01, i)
	}

	time.Sleep(5 * time.Second)
	if ctx01.Err() != nil {
		fmt.Println("监控器取消的原因: ", ctx01.Err())
	}

	fmt.Println("主程序退出！")
}

//contextTimeout使用Demo
func contextTimeoutDemo() {
	fmt.Printf("**********contextTimeoutDemo**********\n")

	//WithTimeout 传入的第二个参数是 time.Duration 类型，它是一个相对的时间，意思是多长时间后超时取消。
	ctx01, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	for i := 0; i < 5; i++ {
		go monitor(ctx01, i)
	}

	time.Sleep(5 * time.Second)
	if ctx01.Err() != nil {
		fmt.Println("监控器取消的原因: ", ctx01.Err())
	}

	fmt.Println("主程序退出！")
}

func monitor02(ctx context.Context, number int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("监控器%v，监控结束。\n", number)
			return
		default:
			value := ctx.Value("item")
			fmt.Printf("监控器%v，正在监控 %v \n", number, value)
			time.Sleep(2 * time.Second)
		}
	}
}

//contextWithValue使用的Demo
//可以给context附加参数
func contextWithValueDemo() {
	fmt.Printf("**********contextWithValueDemo**********\n")

	ctx02, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	//通过Context我们也可以传递一些必须的元数据，这些数据会附加在Context上以供使用。
	//元数据以 Key-Value 的方式传入，Key 必须有可比性，Value 必须是线程安全的。
	ctx03 := context.WithValue(ctx02, "item", "CPU")

	defer cancel()

	for i := 0; i < 5; i++ {
		go monitor02(ctx03, i)
	}

	time.Sleep(5 * time.Second)
	if ctx03.Err() != nil {
		fmt.Println("监控器取消的原因: ", ctx03.Err())
	}
	fmt.Println("主程序退出！！")
}

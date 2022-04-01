package main

import (
	"context"
	"fmt"
	"sync"
)

func woker(ctx context.Context, ch chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("ctx value：%v \n", ctx.Value("key1"))

	for {
		select {
		case num := <-ch:
			fmt.Printf("read from ch value：%d \n", num)
		case <-ctx.Done():
			fmt.Println("work is cancelled")
			return
		}
	}

}

func main() {
	valueCtx := context.WithValue(context.Background(), "key1", "Value00001561")
	cancelCtx, cancelFunc := context.WithCancel(valueCtx)

	intCh := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go woker(cancelCtx, intCh, wg)

	for i := 0; i < 10; i++ {
		intCh <- i
	}

	cancelFunc()
	wg.Wait()
	close(intCh)
}

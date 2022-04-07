package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mutexDemo()
	rwMutexDemo()
}

func add(count *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*count = *count + 1
		lock.Unlock()
	}

	wg.Done()
}

//互斥锁Demo
func mutexDemo() {
	fmt.Printf("**********mutexDemo**********\n")

	/*	注意⚠️
		使用 Mutext 锁虽然很简单，但仍然有几点需要注意：
			同一协程里，不要在尚未解锁时再次使加锁
			同一协程里，不要对已解锁的锁再次解锁
			加了锁后，别忘了解锁，必要时使用 defer 语句
	*/

	//两种创建的方式
	// 第一种
	// var lock01 *sync.Mutex
	// lock01 = new(sync.Mutex)

	// 第二种
	lock02 := &sync.Mutex{}

	var wg sync.WaitGroup
	count := 0
	wg.Add(3)
	go add(&count, &wg, lock02)
	go add(&count, &wg, lock02)
	go add(&count, &wg, lock02)

	wg.Wait()
	fmt.Println("count：", count)
}

//读写锁Demo
func rwMutexDemo() {
	fmt.Printf("**********rwMutexDemo**********\n")

	/*	注意⚠️
		RWMutex，也是如此，它将程序对资源的访问分为读操作和写操作
		   为了保证数据的安全，它规定了当有人还在读取数据（即读锁占用）时，不允计有人更新这个数据（即写锁会阻塞）
		   为了保证程序的效率，多个人（线程）读取数据（拥有读锁）时，互不影响不会造成阻塞，它不会像 Mutex 那样只允许有一个人（线程）读取同一个数据。
	*/

	//创建示例
	//第一种👇
	// var lock *sync.RWMutex
	// lock = new(sync.RWMutex)

	//第二种👇
	lock := &sync.RWMutex{}

	//写锁关上
	lock.Lock()

	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Printf("第 %d 个协程准备开始... \n", i)
			lock.RLock()
			fmt.Printf("第 %d 个协程获得读锁, sleep 1s 后，释放锁\n", i)
			time.Sleep(time.Second)
			lock.RUnlock()
		}(i)
	}

	time.Sleep(time.Second * 2)

	fmt.Println("准备释放写锁，读锁不再阻塞")
	// 写锁一释放，读锁就自由了
	lock.Unlock()

	// 由于会等到读锁全部释放，才能获得写锁
	// 因为这里一定会在上面 4 个协程全部完成才能往下走
	lock.Lock()
	fmt.Println("程序退出...")
	lock.Unlock()

}

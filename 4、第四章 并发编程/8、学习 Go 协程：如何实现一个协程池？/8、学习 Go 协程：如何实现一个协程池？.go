package main

import (
	"fmt"
	"sync"
	"time"
)

type Pool struct {
	work chan func()
	sem  chan struct{}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	pool := New(2)
	for i := 0; i < 5; i++ {
		pool.NewTask(func() {
			defer wg.Done()
			time.Sleep(2 * time.Second)
			fmt.Printf("I'm In time: %s \n", time.Now().Format("2006-01-02 15:04:05"))
		})
	}

	// 保证所有的协程都执行完毕
	wg.Wait()
	fmt.Println("我结束啦")
}

//创建一个线程池对象
//size就是线程池的大小
func New(size int) *Pool {
	return &Pool{
		work: make(chan func()),
		sem:  make(chan struct{}, size),
	}
}

func (p *Pool) NewTask(task func()) {
	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		go p.worker(task)
	}
}

//执行任务
func (p *Pool) worker(task func()) {
	defer func() {
		//减少
		<-p.sem
	}()
	for {
		//去执行任务
		task()
		//把work清理掉
		<-p.work
	}
}

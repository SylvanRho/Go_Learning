package main

import (
	"fmt"
	"unsafe"
)

func main() {
	emptyStructUsedDemo()
	chanUsedEmptyStructDemo()
}

//该对象没有属性，它就是一个空结构体。
type lemp struct{}

func (l lemp) On() {
	fmt.Println("On")
}

func (l lemp) Off() {
	fmt.Println("Off")
}

//空结构体的使用Demo
func emptyStructUsedDemo() {
	fmt.Printf("**********emptyStructUsedDemo**********\n")

	l := lemp{}
	l.On()
	//空结构体是一个不占用空间的对象。
	//output:0
	fmt.Printf("lemp site：%d", unsafe.Sizeof(l))
}

//信道使用空结构体
func chanUsedEmptyStructDemo() {
	fmt.Printf("**********chanUsedEmptyStructDemo**********\n")

	ch := make(chan struct{}, 1)
	go func() {
		<-ch
		// do something
	}()
	ch <- struct{}{}
	// ...
}

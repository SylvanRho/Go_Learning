package main

import (
	"fmt"
)

func main() {
	newUsedDemo()
	makeUsedDemo()
}

type Student struct {
	name string
	age  int
}

//new使用的Demo
func newUsedDemo() {
	fmt.Printf("**********newUsedDemo**********\n")
	//new 只能传递一个参数，该参数为一个任意类型，可以是Go语言内建的类型，也可以是你自定义的类型
	num := new(int)
	fmt.Println(*num) //打印零值：0

	s := new(Student)
	s.name = "wangbm"
	s.age = 19
}

//make使用的Demo
func makeUsedDemo() {
	fmt.Printf("**********makeUsedDemo**********\n")
	//内建函数 make 用来为 slice，map 或 chan 类型（注意：也只能用在这三种类型上）分配内存和初始化一个对象
	//切片
	a := make([]int, 2, 10)
	fmt.Printf("a：%v cap：%d \n", a, cap(a))

	//字典
	b := make(map[string]int)
	fmt.Printf("b：%v \n", b)

	//通道
	c := make(chan int, 10)
	fmt.Printf("c：%v \n", c)

}

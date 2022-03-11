package main

import (
	"fmt"
)

func main() {
	PointerCreateDemo()

	PointerTypeDemo()

	PointerAddressEmptyDemo()

	PointerAndSliceDemo()
}

// 指针的创建
func PointerCreateDemo() {
	//三种方法

	//第一种
	// 定义普通变量
	aint := 1
	// 定义指针变量
	aintPtr := &aint
	fmt.Printf("aintPtr：%p \n", aintPtr)

	// 第二种
	// 创建指针
	astr := new(string)
	// 给指针赋值
	*astr = "Go_Learning"
	fmt.Printf("astr：%s \n", *astr)

	//第三种
	var bint *int // 声明一个指针
	bint = &aint
	fmt.Printf("bint：%p \n", bint)
	*bint = 15 //因为地址相同 所以转更改值的时候 aint的值也改变了
	fmt.Printf("change bint so aint value：%d \n", aint)

	// & ：从一个普通变量中取得内存地址
	// *：当*在赋值操作符（=）的右边，是从一个指针变量中取得变量值，当*在赋值操作符（=）的左边，是指该指针指向的变量

	//打印指针地址方法有两种
	// 第一种
	// fmt.Printf("%p", ptr)

	// 第二种
	// fmt.Println(ptr)
}

//指针类型Demo
func PointerTypeDemo() {
	aint := 1
	astr := "hello"
	abool := false
	arune := 'a'
	afloat := 1.2

	// 可以发现用 *+所指向变量值的数据类型，就是对应的指针类型
	fmt.Printf("aint PointerType：%T \n", &aint) //*int
	fmt.Printf("astr PointerType：%T \n", &astr)
	fmt.Printf("abool PointerType：%T \n", &abool)
	fmt.Printf("arune PointerType：%T \n", &arune)
	fmt.Printf("afloat PointerType：%T \n", &afloat)
	// 所以若我们定义一个只接收指针类型的参数的函数，可以这么写
	// func mytest(ptr *int)  {
	// 	fmt.Println(*ptr)
	// }
}

//空地址指针
func PointerAddressEmptyDemo() {
	a := 25
	var emptyPtr *int //声明一个空指针

	if emptyPtr == nil {
		fmt.Println(emptyPtr)
		emptyPtr = &a // 初始化：将a的内存地址给b
		fmt.Println(emptyPtr)
	}
}

//指针跟切片
func PointerAndSliceDemo() {
	// 如果我们想通过一个函数改变一个数组的值，有两种方法
	// 1、将这个数组的切片做为参数传给函数
	// 2、将这个数组的指针做为参数传给函数

	a := [...]int{89, 84, 15}
	modifyBySlice(a[:])
	fmt.Println(a)
	modifyByPointer(&a)
	fmt.Println(a)

}

//使用切片修改（推荐）
func modifyBySlice(sls []int) {
	sls[0] = 90
}

//使用指针修改
func modifyByPointer(arr *[3]int) {
	(*arr)[1] = 90
}

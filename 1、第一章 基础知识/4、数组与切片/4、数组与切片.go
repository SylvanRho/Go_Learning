package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	ArrDemo()
	SliceDemo()
}

//1、数组
func ArrDemo() {
	fmt.Printf("\n*******Array数组*******\n")
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	// 初始化数组的几种方法
	// 第一种方法
	var arr2 [3]int = [3]int{1, 2, 3}
	fmt.Printf("arr2 [0] value：%d\n", arr2[0])

	// 第二种方法
	arr3 := [3]int{1, 2, 3}
	fmt.Printf("arr3 [0] value：%d\n", arr3[0])

	//不定长数组
	arr4 := [...]int{4, 6, 5}
	fmt.Printf("arr4 [0] value：%d\n", arr4[0])

	//%T打印类型
	arr5 := [...]int{4, 6, 5, 7}
	fmt.Printf("%d 的类型是: %T\n", arr4, arr4)
	fmt.Printf("%d 的类型是: %T\n", arr5, arr5)

	//别名类型
	type arrLen3 [3]int
	arr6 := arrLen3{1, 2, 3}
	fmt.Printf("%d 的类型是: %T\n", arr6, arr6)

	//偷懒的写法
	arr7 := [4]int{2: 3}
	fmt.Printf("arr7 的结构体是 %v", arr7)
}

//2、切片
func SliceDemo() {
	fmt.Printf("\n*******Slice切片*******\n")
	myarr01 := [...]int{1, 2, 3}
	fmt.Printf("%d 的类型是: %T\n", myarr01[0:2], myarr01[0:2])

	// 切片的构造，有四种方式
	SliceDemo01()
	SliceDemo02()
	SliceDemo03()
	SliceDemo04()
	SliceAppendDemo()
}

//1、对数组进行片段截取 主要有如下两种写法
func SliceDemo01() {
	fmt.Printf("\n**********SliceDemo01**********\n")
	myarr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("myarr 的长度为：%d，容量为：%d\n", len(myarr), cap(myarr))

	mysli1 := myarr[1:3] //3的含义为 索引为2，内容为[2 3]（3减1，等于到下标3的位置但是下标三不包含）的元素
	fmt.Printf("mysli1 的长度为：%d，容量为：%d\n", len(mysli1), cap(mysli1))
	fmt.Println(mysli1)

	mysli2 := myarr[1:3:4] //4的含义为容量范围从下标1到下标四，长度为3（4减1，等于到下标4的位置 但是下标四的位置不包括）
	fmt.Printf("mysli2 的长度为：%d，容量为：%d\n", len(mysli2), cap(mysli2))
	fmt.Println(mysli2)
}

//2、从头声明赋值（例子如下）
func SliceDemo02() {
	fmt.Printf("\n**********SliceDemo02**********\n")
	//切片是引用类型，所以你不对它进行赋值的话，它的零值（默认值）是 nil
	var strList []string
	var numList []int
	var numListEmpty = []int{}
	fmt.Printf("strList 的长度为：%d，容量为：%d\n", len(strList), cap(strList))
	fmt.Printf("numList 的长度为：%d，容量为：%d\n", len(numList), cap(numList))
	fmt.Printf("numListEmpty 的长度为：%d，容量为：%d\n", len(numListEmpty), cap(numListEmpty))
}

//3、make函数构造
func SliceDemo03() {
	fmt.Printf("\n**********SliceDemo03**********\n")
	a := make([]int, 5)
	b := make([]int, 2, 10)
	fmt.Println(a, b)
	fmt.Printf("a 的长度为 %d，容量为：%d\n", len(a), cap(a))
	fmt.Printf("b 的长度为 %d，容量为：%d\n", len(b), cap(b))
}

//4、偷懒的写法
func SliceDemo04() {
	fmt.Printf("\n**********SliceDemo04**********\n")
	a := []int{4: 2}
	fmt.Println(a)
	fmt.Println(len(a), cap(a))
}

//切片追加Demo
func SliceAppendDemo() {
	fmt.Printf("\n**********SliceAppendDemo**********\n")
	myarr := []int{1}
	// 追加一个元素
	myarr = append(myarr, 2)
	// 追加多个元素
	myarr = append(myarr, 3, 4)
	// 追加一个切片, ... 表示解包，不能省略
	myarr = append(myarr, []int{7, 8}...)
	// 在第一个位置插入元素
	myarr = append([]int{0}, myarr...)
	// 在中间插入一个切片(两个元素)
	myarr = append(myarr[:5], append([]int{5, 6}, myarr[5:]...)...)
	fmt.Println(myarr)
}

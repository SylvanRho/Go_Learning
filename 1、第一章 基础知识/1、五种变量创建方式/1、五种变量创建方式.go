package main

import "fmt"

func main() {

	// 第一种 ：一行声明一个变量
	// var <name> <type>

	//第二种 多个变量一起声明
	var (
		name   string = "青"
		age    int    = 16
		gender string = "高一"
	)
	fmt.Println(name, age, gender)

	//第三种：声明和初始化一个变量
	name2 := "ceshi"
	//第三种：声明和初始化一个变量
	age2, gender2 := 18, "大一"
	fmt.Println(name2, age2, gender2)

	// 第五种：new 函数声明一个指针变量
	var agePtr = &age2 //&符号的意思是对变量取地址
	println(agePtr)

	ptr2 := new(int)
	println("Address：", ptr2)
	println("value：", *ptr2) // * 后面接指针变量，表示从内存地址中取出值

	println("funcNewInt：", newInt())

	//匿名变量，也称作占位符
	a, _ := GetData()
	_, b := GetData()
	println("a：", a)
	println("b：", b)
}

//使用表达式 new(Type) 将创建一个Type类型的匿名变量，初始化为Type类型的零值，然后返回变量地址，返回的指针类型为*Type
func newInt() *int {
	return new(int)
}

func GetData() (int, int) {
	return 100, 200
}

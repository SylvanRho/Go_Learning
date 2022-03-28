package main

import "fmt"

func main() {
	staticType()
	dynamicType()
	InterfaceComposition()
}

//静态类型
func staticType() {
	fmt.Printf("**********staticType**********\n")
	// 所谓的静态类型（即 static type），就是变量声明的时候的类型。
	var age int
	var name string
	age = 19
	name = "小明"
	fmt.Printf("姓名：%s 年龄：%d", name, age)
}

//动态类型
func dynamicType() {
	fmt.Printf("**********dynamicType**********\n")
	//动态类型是程序运行时系统才能看见的类型。
	var i interface{}
	i = 19
	i = "wang"

	fmt.Printf("Type：%T，value：%v", i, i)
}

//接口的组成
func InterfaceComposition() {
	fmt.Printf("**********InterfaceComposition**********\n")
	//每个接口变量，实际上都是由一 pair 对（type 和 data）组合而成

	//第一种普通创建
	//type:int		data:25
	var age01 int = 25

	//也可以用另一种方法创建
	age02 := (int)(25)
	//或者使用 age := (interface{})(25)

	fmt.Printf("age01  Type：%T，value：%v \n", age01, age02)
	fmt.Printf("age02  Type：%T，value：%v \n", age02, age02)

	//⚠️接口分两种
	//有方法的接口   叫做iface
	//没有方法的接口 叫做eface
}

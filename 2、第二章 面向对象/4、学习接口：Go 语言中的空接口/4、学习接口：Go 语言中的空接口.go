package main

import (
	"fmt"
)

func main() {

	//空接口：空接口是特殊形式的接口类型，普通的接口都有方法，而空接口没有定义任何方法口
	//因此，我们可以说所有类型都至少实现了空接口。
	//type empty_iface interface {}
	var i interface{}
	fmt.Printf("type:%T value:%v", i, i)

	emptyInterfaceCanUsedType()
	funcUsedEmptyInterface()
	StoreAnyTypeEmptyInterface()
	pitForEmptyInterface()
}

//空接口类型可以使用的类型
func emptyInterfaceCanUsedType() {
	fmt.Printf("**********emptyInterfaceCanUsedType**********\n")
	var i interface{}

	i = 10
	fmt.Println(i)

	i = "hello"
	fmt.Println(i)

	i = false
	fmt.Println(i)
}

//函数使用空接口
func funcUsedEmptyInterface() {
	fmt.Printf("**********funcUsedEmptyInterface**********\n")
	a := "a"
	b := 10
	c := false
	receiveAnyTypeFunc(a, b, c)
}

//接收任意类型的方法数组  如果是单个任意类型的话 可以改成ifaces interface{}
func receiveAnyTypeFunc(ifaces ...interface{}) {
	for _, iface := range ifaces {
		fmt.Println(iface)
	}
}

//储存任意类型的空接口集合类型
func StoreAnyTypeEmptyInterface() {
	fmt.Printf("**********StoreAnyTypeEmptyInterface**********\n")
	any := make([]interface{}, 5)
	any = append(any, 1)
	any[1] = "hello"
	any[2] = false
	for _, item := range any {
		fmt.Println(item)
	}
}

func pitForEmptyInterface() {
	fmt.Printf("**********pitForEmptyInterface**********\n")

	//pit1⚠️：空接口可以承载任意值，但不代表任意类型就可以承接空接口类型的值

	// 声明i变量, 类型为interface{}, 初始值为a, 此时i的值变为1
	// var a int = 1

	// 声明i变量, 类型为interface{}, 初始值为a, 此时i的值变为1
	// var i interface{} = a

	// 声明i变量, 类型为interface{}, 初始值为a, 此时i的值变为1
	// var b int = i//报错了

	//pit2⚠️:当空接口承载数组和切片后，该对象无法再进行切片

	// sil := []int{2, 3, 5, 7, 11, 13}
	// var i interface{}
	// i = sil

	//不能再切片了  报错了
	// g := i[1:3]

	// pit3⚠️:当你使用空接口来接收任意类型的参数时，它的静态类型是 interface{}，但动态类型（是 int，string 还是其他类型）我们并不知道，因此需要使用类型断言。
	a := 10
	b := "hello"
	dynamicTypeByEmptyInterface(a)
	dynamicTypeByEmptyInterface(b)
}

func dynamicTypeByEmptyInterface(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("您的参数是int类型")
	case string:
		fmt.Println("参数的类型是 string")
	default:
		fmt.Println("您的参数类型未知")
	}
}

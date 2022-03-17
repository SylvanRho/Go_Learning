package main

import (
	"fmt"
)

func main() {

	// Type Assertion（中文名叫：类型断言），通过它可以做到以下几件事情

	// 1、检查 i 是否为 nil
	// 2、检查 i 存储的值是否为某个类型

	typeAssertionDemo01()
	typeAssertionDemo02()

	findType(10)      // int
	findType("hello") // string

	var k interface{} // nil
	findType(k)

	findType(10.23) //float64
}

//断言demo01
func typeAssertionDemo01() {
	fmt.Printf("\n**********typeAssertionDemo01**********\n")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var i interface{} = 10
	t1 := i.(int)
	fmt.Println(t1)

	//在执行第二次断言的时候失败了，并且触发了 panic
	t2 := i.(string)
	fmt.Println(t2)
}

//断言demo02
func typeAssertionDemo02() {
	fmt.Printf("\n**********typeAssertionDemo02**********\n")

	var i interface{} = 10
	t1, ok := i.(int)
	fmt.Printf("%d-%t\n", t1, ok)

	fmt.Println("=====分隔线1=====")

	//失败了，但并没有触发了 panic
	t2, ok := i.(string)
	fmt.Printf("%s-%t\n", t2, ok)

	fmt.Println("=====分隔线2=====")

	var k interface{} // nil
	t3, ok := k.(interface{})
	fmt.Println(t3, "-", ok)

	fmt.Println("=====分隔线3=====")
	k = 10
	t4, ok := k.(interface{})
	fmt.Printf("%d-%t\n", t4, ok)

	fmt.Println("=====分隔线4=====")
	t5, ok := k.(int)
	fmt.Printf("%d-%t\n", t5, ok)
}

//查找类型
func findType(i interface{}) {
	fmt.Printf("\n**********findType**********\n")
	switch x := i.(type) {
	case int:
		fmt.Println(x, "is int")
	case string:
		fmt.Println(x, "is string")
	case nil:
		fmt.Println(x, "is nil")
	default:
		fmt.Println(x, "not type matched")
	}
}

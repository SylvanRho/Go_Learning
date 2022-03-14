package main

import (
	"fmt"
)

func main() {
	//延迟调用结构如下👇
	//defer xxx()
	deferSimpleDemo()
	deferVariableDemo01()
	deferVariableDemo02()
	multipleDefer()
	deferAndReturn()
}

//defer简单Demo
func deferSimpleDemo() {
	fmt.Printf("\n**********deferSimpleDemo**********\n")
	defer fmt.Println("A") //整个方法结束后执行的
	fmt.Println("B")
}

//defer变量的Demo
func deferVariableDemo01() {
	fmt.Printf("\n**********deferVariableDemo01**********\n")
	name := "go"
	//虽然是结束的时候执行 但是输出的是go 证明变量是👇
	//使用 defer 只是延时调用函数，此时传递给函数里的变量，不应该受到后续程序的影响。
	defer fmt.Println(name)

	name = "python"
	fmt.Println(name)
}

//defer变量的匿名函数Demo
func deferVariableDemo02() {
	fmt.Printf("\n**********deferVariableDemo02**********\n")
	name := "go"
	//但是如果是匿名函数的时候 取得值就是最后的值  -->  python
	defer func() {
		fmt.Println(name)
	}()
	name = "python"
	fmt.Println(name)
}

//多个defer
func multipleDefer() {
	fmt.Printf("\n**********multipleDefer**********\n")
	//多个defer 是反序调用的，有点类似栈一样，后进先出。
	name := "go"
	defer fmt.Println(name)

	name = "python"
	defer fmt.Println(name)

	name = "java"
	fmt.Println(name)
}

var name string = "go"

//defer跟return
func deferAndReturn() {
	fmt.Printf("\n**********deferAndReturn**********\n")
	myname := myfunc()
	fmt.Printf("main 函数里的name: %s\n", name)  //在这里拿到的name 是myfunc已经执行完成的执行完defer方法以后的name 也就是python
	fmt.Println("main 函数里的myname: ", myname) //这里拿到的myname 是myfunc在执行完成后 但是还没执行defer的方法的时候return的值
	//也就是说 defer 是return 后才调用的。所以在执行 defer 前，myname 已经被赋值成 go 了。
}

func myfunc() string {
	defer func() {
		name = "python"
	}()

	fmt.Printf("myfunc 函数里的name：%s\n", name)
	return name
}

package main

import (
	"fmt"
)

func main() {
	// for 循环的基本模型 👇
	// for [condition |  ( init; condition; increment ) | Range]
	// {
	//    statement(s);
	// }
	//可以看到 for 后面，可以接三种类型的表达式。
	// 1、接一个条件表达式
	// 2、接三个表达式
	// 3、接一个 range 表达式
	// 4、还可以不接
	forDemo01()
	forDemo02()
	forDemo03()
	forDemo04()
}

//for循环的第一种情况（条件表达式）
func forDemo01() {
	fmt.Printf("**********forDemo01**********\n")
	a := 1
	for a <= 5 {
		fmt.Printf("%d \n", a)
		a++
	}
}

//for循环的第二种情况(接三个表达式)
func forDemo02() {
	fmt.Printf("**********forDemo02**********\n")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d \n", i)
	}
}

//for循环的第三种情况(接 for-range 语句)
func forDemo03() {
	fmt.Printf("**********forDemo03**********\n")
	myarr := [...]string{"world", "python", "go"}
	//range会返回两个 索引和数据，若你后面的代码用不到索引，需要使用 _ 表示 。
	for _, item := range myarr {
		fmt.Println(item)
	}
}

//for循环的第四种情况（不接表达式：无限循环）
func forDemo04() {
	fmt.Printf("**********forDemo04**********\n")
	// for {
	// 	代码块
	// }

	// // 等价于
	// for ;; {
	// 	代码块
	// }
	var i int = 1
	for {
		if i <= 5 {
			fmt.Printf("Hello,%d \n", i)
			i++
		} else {
			break
		}
	}
}

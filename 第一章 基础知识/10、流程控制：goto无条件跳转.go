package main

import (
	"fmt"
)

func main() {
	// goto基本模型(跳转)
	// goto 标签;
	// ...
	// ...
	// 标签: 表达式;
	gotoSimpleDemo01()
	gotoModeledFor()
	gotoModeledBreak()
	gotoModeledContinue()
}

//一个简单的示例
func gotoSimpleDemo01() {
	fmt.Printf("\n**********gotoSimpleDemo01**********\n")
	// goto 可以打破原有代码执行顺序，直接跳转到某一行执行代码。
	goto flag
	fmt.Println("A") //不会执行
flag:
	fmt.Println("B")
}

//goto模仿For
func gotoModeledFor() {
	fmt.Printf("\n**********gotoModeledFor**********\n")
	//打印1~5的循环
	i := 1
flag:
	if i <= 5 {
		fmt.Println(i)
		i++
		goto flag
	}
}

//goto模仿Break
func gotoModeledBreak() {
	//使用 goto 实现 类型 break 的效果。
	fmt.Printf("\n**********gotoModeledBreak**********\n")
	i := 1
	for {
		if i > 5 {
			goto flag
		}
		fmt.Println(i)
		i++
	}
flag:
	fmt.Println("我结束了")
}

//goto模仿Continue
func gotoModeledContinue() {
	//使用 goto 实现 类型 continue的效果，打印 1到10 的所有偶数。
	fmt.Printf("\n**********gotoModeledContinue**********\n")
	i := 1
flag:
	for i <= 10 {
		if i%2 == 0 {
			fmt.Println(i)
		}
		i++
		goto flag
	}
}

//注意
func Attention() {
	// goto语句与标签之间不能有变量声明，否则编译错误。
	fmt.Println("start")
	goto flag
	//下面两句，放出来编译不过👇
	// var say = "hello oldboy"
	// fmt.Println(say)
flag:
	fmt.Println("end")
}

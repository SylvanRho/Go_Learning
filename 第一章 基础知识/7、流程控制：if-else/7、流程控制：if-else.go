package main

import (
	"fmt"
)

func main() {
	// 条件语句模型👇
	// if 条件 1 {
	// 	分支 1
	//   } else if 条件 2 {
	// 	分支 2
	//   } else if 条件 ... {
	// 	分支 ...
	//   } else {
	// 	分支 else
	//   }
	//important note⚠️：Go编译器，对于 { 和 } 的位置有严格的要求，它要求 else if （或 else）和 两边的花括号，必须在同一行。
	onlyIf_Else()
	moreIf_Else()
	advancedIf_Else()
}

//单分支if-else
func onlyIf_Else() {
	age := 20
	gender := "male"
	if age > 18 {
		fmt.Println("已经成年啦！")
	}

	// 如果条件里需要满足多个条件，可以使用 && 和 ||

	// 	1、&&：表示且，左右都需要为true，最终结果才能为 true，否则为 false

	// 	2、||：表示或，左右只要有一个为true，最终结果即为true，否则 为 false
	if age > 18 && gender == "male" {
		fmt.Println("是成年男性！")
	}
}

//多分支if-else
func moreIf_Else() {
	age := 20
	if age > 20 {
		fmt.Println("已经成年啦！")
	} else if age > 12 {
		fmt.Println("还是青少年哦")
	} else {
		fmt.Println("还不是青少年")
	}
}

//高级If-Else
func advancedIf_Else() {
	//声明一个变量以后 再判断
	if age := 20; age > 18 {
		fmt.Printf("年龄是：%d 已经是成年人啦！", age)
	}
}

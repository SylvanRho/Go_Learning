package main

import (
	"fmt"
)

func main() {
	mapDemo()
	fmt.Println()
}

//字典Demo
func mapDemo() {
	fmt.Println("**********mapDemo**********\n")
	// map[KEY_TYPE]VALUE_TYPE
	// 声明初始化字典 三种！
	// 第一种方法
	var scores01 map[string]int = map[string]int{"english": 80, "chinese": 85}
	fmt.Println("First map:", scores01)
	// 第二种方法
	scores02 := map[string]int{"english": 80, "chinese": 85}
	fmt.Println("Second map:", scores02)
	// 第三种方法
	scores03 := make(map[string]int)
	scores03["english"] = 80
	scores03["chinese"] = 85
	fmt.Println("Thirds map:", scores03)

	mapOperationDemo()
	mapContainsDemo()
	mapIterationDemo()
}

//字典操作
func mapOperationDemo() {
	fmt.Printf("**********mapOperationDemo**********\n")
	scores := make(map[string]int)
	// 添加元素
	scores["english"] = 80
	scores["chinese"] = 95
	fmt.Printf("增加元素后：%v\n", scores)
	//fmt.Println(scores)
	scores["chinese"] = 78
	fmt.Printf("更新元素后：%v\n", scores)
	// 读取元素，直接使用 [key] 即可 ，如果 key 不存在，也不报错，会返回其value-type 的零值。
	fmt.Printf("scores key[math] value: %d\n", scores["math"])
	// 删除元素，使用 delete 函数，如果 key 不存在，delete 函数会静默处理，不会报错
	delete(scores, "math")
}

//字典是否包含用例
func mapContainsDemo() {
	fmt.Printf("**********mapContainsDemo**********\n")
	scores := map[string]int{"English": 80, "Chinese": 75}
	//判断key在不在
	math, ok := scores["math"]
	if ok {
		fmt.Println(math)
	} else {
		fmt.Println("math不存在")
	}

	//上面代码优化后👇
	if math01, ok01 := scores["math"]; ok01 {
		fmt.Println(math01)
	} else {
		fmt.Println("math不存在")
	}
}

//字典迭代用例
func mapIterationDemo() {
	fmt.Printf("**********mapIterationDemos**********\n")
	scores := map[string]int{"english": 80, "chinese": 85}
	// 循环还分三种

	//第一种
	fmt.Println("First Iteration")
	for subject, score := range scores {
		fmt.Printf("score subject: %s，score：%d \n", subject, score)
	}

	//第二种 只取key 不需要占位符
	fmt.Println("Second Iteration")
	for subject := range scores {
		fmt.Printf("score subject: %s \n", subject)
	}

	//第三种 取value 需要占位符
	fmt.Println("Thirds Iteration")
	for _, score := range scores {
		fmt.Printf("score score： %d \n", score)
	}
}

func boolDemo() {
	fmt.Printf("**********boolDemo**********\n")
	// 在 Python 中使用 not 对逻辑值取反，而 Go 中使用 ! 符号
	male := true
	fmt.Println(!male == false)
	//或者
	fmt.Println(male != false)

	//在 Go 语言中，则使用 && 表示且，用 || 表示或，并且有短路行为（即左边表达式已经可以确认整个表达式的值，那么右边将不会再被求值。
	var age int = 15
	var gender string = "male"

	//  && 两边的表达式都会执行
	fmt.Println(age > 18 && gender == "male")
	// gender == "male" 并不会执行
	fmt.Println(age < 18 || gender == "male")
}

//bool转int
func bool2int(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

//int转bool
func int2bool(i int) bool {
	return i != 0
}

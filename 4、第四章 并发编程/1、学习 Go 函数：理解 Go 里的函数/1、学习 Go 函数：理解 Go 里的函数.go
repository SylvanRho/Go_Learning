package main

import "fmt"

func main() {
	funcCreateDemo()

	anonymousFunctionDemo()
}

//函数的创建
func funcCreateDemo() {
	fmt.Printf("\n**********funcCreateDemo**********\n")

	/*
		//函数创建的语法
		func 函数名(形式参数列表)(返回值列表){
			函数体
		}
	*/

	fmt.Println("sum：", sum(1, 2))

	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	differentTypeParamFunc(v1, v2, v3, v4)

	fmt.Println("Sum：", Sum(1, 88, 23, 648, 62))

	// 接收参数用逗号分隔
	a, b := returnMultipleParam(2)
	fmt.Println(a, b)

	fmt.Println(returnHaveNameParam(5))
}

func sum(a int, b int) int {
	fmt.Printf("\n**********sum**********\n")
	return a + b
}

// 多个可变参数函数传递参数
func Sum(args ...int) int {
	// 利用 ... 来解序列
	result := theSameTypeParamSumFunc(args...)
	return result
}

//相同类型相加的方法
// 使用 ...类型，表示一个元素为int类型的切片
func theSameTypeParamSumFunc(args ...int) int {
	fmt.Printf("\n**********theSameTypeParamSumFunc**********\n")
	var result int
	for _, v := range args {
		result += v
	}
	return result
}

//不同类型的多个参数方法
func differentTypeParamFunc(args ...interface{}) {
	fmt.Printf("\n**********differentTypeParamFunc**********\n")

	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

//返回多个返回值
func returnMultipleParam(a int) (int, int) {
	fmt.Printf("\n**********returnMultipleParam**********\n")
	b := a * 2
	return a, b
}

//返回拥有变量名的值
func returnHaveNameParam(a int) (b int, c int) {
	fmt.Printf("\n**********returnHaveNameParam**********\n")
	b = a * 2
	c = b * 2
	return
}

//匿名函数Demo
func anonymousFunctionDemo() {
	fmt.Printf("\n**********anonymousFunctionDemo**********\n")
	/*
		//定义
		func(参数列表)(返回参数列表){
			函数体
		}
	*/

	//定义后立马执行
	func(a int) {
		fmt.Println(a)
	}(100)

	//做为回调函数使用
	visit([]int{1, 2, 5, 3, 57, 5}, func(i int) {
		fmt.Println(i)
	})
}

//将方法做为参数传递进来
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

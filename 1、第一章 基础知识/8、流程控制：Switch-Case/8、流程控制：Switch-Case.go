package main

import (
	"fmt"
)

func main() {
	// 选择语句模型👇
	// switch 表达式 {
	// 	case 表达式1:
	// 		代码块
	// 	case 表达式2:
	// 		代码块
	// 	case 表达式3:
	// 		代码块
	// 	case 表达式4:
	// 		代码块
	// 	case 表达式5:
	// 		代码块
	// 	default:
	// 		代码块
	// }
	switchSimpleDemo()
	multipleConditionsSwitch()
	errorSwitch()
	switchAppendFunc()
	switchLikeIfelse()
	switchFallthrough()
}

//switch简单Demo
func switchSimpleDemo() {
	fmt.Printf("**********switchSimpleDemo**********\n")
	education := "本科"
	switch education {
	case "博士":
		fmt.Println("我是博士")
	case "研究生":
		fmt.Println("我是研究生")
	case "本科":
		fmt.Println("我是本科生")
	case "大专":
		fmt.Println("我是大专生")
	case "高中":
		fmt.Println("我是高中生")
	default:
		fmt.Println("学历未达标...")
	}
}

//多个条件的Switch
func multipleConditionsSwitch() {
	fmt.Printf("**********multipleConditionsSwitch**********\n")
	month := 2

	switch month {
	case 3, 4, 5:
		fmt.Println("春天来了~")
	case 6, 7, 8:
		fmt.Println("夏天来了~")
	case 9, 10, 11:
		fmt.Println("秋天来了~")
	case 12, 1, 2:
		fmt.Println("冬天来了~")
	default:
		fmt.Println("未知季节")
	}
}

//错误的Switch
func errorSwitch() {
	fmt.Printf("**********errorSwitch**********\n")
	gender := "male"
	//同时出现先两个一样的
	// switch gender {
	// case "male", "male":
	// 	fmt.Println("男性")
	// }

	switch gender {
	case "male":
		fmt.Println("男性")
	// 与上面重复👇是错误的
	// case "male":
	// 	fmt.Println("男性")
	case "female":
		fmt.Println("女性")
	}
}

// switch后面跟函数
func switchAppendFunc() {
	fmt.Printf("**********switchAppendFunc**********\n")
	chinese := 60
	english := 45
	math := 150

	switch getResult(chinese, english, math) {
	case true:
		fmt.Println("该同学所有成绩都合格~")
	case false:
		fmt.Println("该同学有挂科记录")
	}
}

//获取成绩是否都符合标准
func getResult(args ...int) bool {
	for _, i := range args {
		if i < 60 {
			return false
		}
	}
	return true
}

//switch不接表达式
func switchLikeIfelse() {
	fmt.Printf("**********switchLikeIfelse**********\n")
	score := 30
	// switch 后可以不接任何变量、表达式、函数。
	switch {
	case score >= 95 && score <= 100:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 60:
		fmt.Println("合格")
	case score > 60:
		fmt.Println("不合格")
	default:
		fmt.Println("输入有误~")
	}
}

//switch穿透能力
func switchFallthrough() {
	fmt.Printf("**********switchFallthrough**********\n")
	s := "hello"
	// fallthrough 只能穿透一层，意思是它让你直接执行下一个case的语句，而且不需要判断条件。
	//所以输出的是hello xxx
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s == "xxx":
		fmt.Println("xxx")
	case s != "hello":
		fmt.Println("world")
	}
}

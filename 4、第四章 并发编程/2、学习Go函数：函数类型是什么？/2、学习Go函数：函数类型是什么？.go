package main

import "fmt"

func main() {
	// 转换成 Greeting 类型的函数对象
	greet01 := Greeting(english)
	// 或者
	var greet02 Greeting = english

	fmt.Println(greet01("xm01"))

	fmt.Println(greet02("xm02"))

	greet01.say("World")
}

//定义了一个传参是string  返回是string的函数类型属性
type Greeting func(name string) string

//跟Greeting函数类型相同 所以可以当成函数的实现
func english(name string) string {
	fmt.Printf("**********english**********\n")
	return "Hello " + name
}

//greet 做为 Greeting 类型的对象，也拥有 Greeting 类型的所有方法，比如下面的 say 方法
func (g Greeting) say(n string) {
	fmt.Printf("**********say**********\n")
	fmt.Println(g(n))
}

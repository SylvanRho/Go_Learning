package main

import "fmt"

func main() {
	interfaceRule01()
	interfaceRule02()
	interfaceRule03()
}

//接口规则1
//接口是一组固定的方法集，由于静态类型的限制，接口变量有时仅能调用其中特定的一些方法。
func interfaceRule01() {
	fmt.Printf("**********interfaceRule01**********")

	//因为我们的phone对象显示声明为 Phone 接口类型，因此 phone调用的方法会受到此接口的限制。
	//因为接口中没有call_wechat方法 所以没法调用
	// var phone Phone
	// phone = IPhone{name: "苹果8plus"}

	//可以不显示的声明为 Phone接口类型 ，但要清楚 phone 对象实际上是隐式的实现了 Phone 接口，
	//如此一来，方法的调用就不会受到接口类型的约束。
	phone := IPhone{name: "苹果8plus"}
	phone.call()
	phone.call_wechat()

}

//定义一个接口
type Phone interface {
	call()
}

//定义一个手机结构体
type IPhone struct {
	name string
}

//实现打电话方法
func (phone IPhone) call() {
	fmt.Println(phone.name + "正在打电话")
}

//定义一个IPhone结构体的打微信方法
func (phone IPhone) call_wechat() {
	fmt.Println(phone.name + "拨打微信电话")
}

//接口规则02
//Go 语言中的函数调用都是值传递的，变量会在方法调用前进行类型转换。
func interfaceRule02() {
	fmt.Printf("**********interfaceRule02**********")
	a := "hello"
	i := 15

	printType(a)
	printType(i)

	//如果想在方法内部写 就需要显性转化一下
	switch interface{}(a).(type) {
	case string:
		fmt.Println("参数的类型是 string")
	default:
	}
}

//打印类型
func printType(i interface{}) {
	// 当一个函数接口 interface{} 空接口类型时，我们说它可以接收什么任意类型的参数
	// 当你使用这种写法时 Go 会默默地为我们做一件事，就是把传入函数的参数值（注意：Go 语言中的函数调用都是值传递的）的类型隐式的转换成 interface{} 类型。
	switch i.(type) {
	case int:
		fmt.Print("这是int类型的")
	case string:
		fmt.Println("这是string类型的")
	default:
		fmt.Println("这是未知类型的")
	}
}

//接口规则3：类型断言中的隐式转换
func interfaceRule03() {
	fmt.Printf("**********interfaceRule03**********")
	var a interface{} = 25

	//当类型断言完成后
	switch b := a.(type) {
	case int:
		//会返回一个静态类型为你断言的类型的对象，也就是说，当我们使用了类型断言，Go 实际上又会默认为我们进行了一次隐式的类型转换。
		fmt.Print(b) //output:25
		// b type is int
		// b.(type)//所以这句话要报错 因为b不是空接口类型
	}
}

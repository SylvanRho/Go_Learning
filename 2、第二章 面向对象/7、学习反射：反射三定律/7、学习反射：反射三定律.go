package main

import (
	"fmt"
	"reflect"
)

func main() {
	//反射三大定律
	// 1、反射可以将接口类型变量 转换为“反射类型对象”；
	// 2、反射可以将 “反射类型对象”转换为 接口类型变量；
	// 3、如果要修改 “反射类型对象” 其类型必须是 可写的；

	reflectDemo01()
	reflectDemo02()
	reflectDemo03()
}

//第一定律
//从接口变量到反射对象的转换。
func reflectDemo01() {
	fmt.Printf("\n**********reflectDemo01**********\n")

	// 为了实现从接口变量到反射对象的转换，需要提到 reflect 包里很重要的两个方法：
	// 1、reflect.TypeOf(i) ：获得接口值的类型
	// 2、reflect.ValueOf(i)：获得接口值的值
	// 这两个方法返回的对象，我们称之为反射对象：Type object 和 Value object。

	var age interface{} = 25

	fmt.Printf("age 的类型为 %T 值为 %d \n", age, age)

	t := reflect.TypeOf(age)
	v := reflect.ValueOf(age)

	fmt.Printf("从接口变量到反射对象：Type对象的类型为：%T \n", t)
	fmt.Printf("从接口变量到反射对象：Value对象的类型为：%T \n", v)
}

//第二定律
//从反射对象到接口变量的转换。
func reflectDemo02() {
	fmt.Printf("\n**********reflectDemo02**********\n")

	var age interface{} = 25

	fmt.Printf("原始接口变量的类型为 %T，值为 %v \n", age, age)

	t := reflect.TypeOf(age)
	v := reflect.ValueOf(age)

	// 从接口变量到反射对象
	fmt.Printf("从接口变量到反射对象：Type对象的类型为 %T \n", t)
	fmt.Printf("从接口变量到反射对象：Value对象的类型为 %T \n", v)

	//⚠️只有 Value 才能逆向转换，而 Type 则不行，这也很容易理解，如果 Type 能逆向，那么逆向成什么呢？
	// 从反射对象到接口变量
	// i := v.Interface()
	//如果使用下面这个 类型不对 要进panic
	i := v.Interface().(int)
	fmt.Printf("从反射对象到接口变量：新对象的类型为 %T 值为 %v \n", i, i)
}

//第三定律
//settable （可设置性，或可写性）
func reflectDemo03() {
	fmt.Printf("\n**********reflectDemo03**********\n")

	//Go 语言里的函数都是值传递，只要你传递的不是变量的指针，你在函数内部对变量的修改是不会影响到原始的变量的。
	//所以如果传递的不是接口变量的指针，反射世界里的变量值始终将只是真实世界里的一个拷贝，你对该反射对象进行修改，并不能反映到真实世界里。

	//因此在反射的规则里
	// 1、不是接收变量指针创建的反射对象，是不具备『可写性』的
	// 2、是否具备『可写性』，可使用 CanSet() 来获取得知
	// 3、对不具备『可写性』的对象进行修改，是没有意义的，也认为是不合法的，因此会报错。

	var name string = "Go编程"

	v := reflect.ValueOf(name)
	fmt.Println("可写性为:", v.CanSet())

	//要让反射对象具备可写性，需要注意两点
	// 1、创建反射对象时传入变量的指针
	// 2、使用 Elem()函数返回指针指向的数据

	v1 := reflect.ValueOf(&name)
	fmt.Println("v1 可写性为:", v1.CanSet())

	v2 := v1.Elem()
	fmt.Println("v2 可写性为:", v2.CanSet())

	v2.SetString("Python编程时光")
	fmt.Println("通过反射对象进行更新后，真实世界里 name 变为：", name)

}

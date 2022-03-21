package main

import (
	"fmt"
	"reflect"
)

func main() {
	kindDemo()

	typeConverter()

	sliceReflectDemo()

	fieldsReflectDemo()
	methodsReflectDemo()

	dynamicCallingReflectFunc()
}

type Person struct {
	name   string
	age    int
	gender string
}

type Person01 struct {
	name   string
	age    int
	gender string
}

//Kind函数
//用于查看更深层次的类型
func kindDemo() {
	fmt.Printf("\n**********kindDemo**********\n")

	// Type 对象 和 Value 对象都可以通过 Kind() 方法返回对应的接口变量的基础类型。
	// reflect.TypeOf(m).Kind()
	// reflect.ValueOf(m).Kind()
	m := Person{}

	// kind用法：第一种：传入值
	t := reflect.TypeOf(m)
	fmt.Println("Type：", t)
	fmt.Println("Kind：", t.Kind())

	// kind用法：第二种：传入指针

	fmt.Println("=====分隔线1=====")

	t1 := reflect.TypeOf(&m)
	fmt.Println("&m Type: ", t1)
	fmt.Println("&m Kind: ", t1.Kind())

	fmt.Println("m Type：", t1.Elem())
	fmt.Println("m Kind:", t1.Elem().Kind())

	//使用 ValueOf 方法
	fmt.Println("=====分隔线2=====")

	v := reflect.ValueOf(&m)

	fmt.Println("&m Type: ", v.Type())
	fmt.Println("&m Kind: ", v.Kind())

	fmt.Println("m Type: ", v.Elem().Type())
	fmt.Println("m Kind: ", v.Elem().Kind())
}

//类型转换
func typeConverter() {
	fmt.Printf("\n**********typeConverter**********\n")
	reflectValue2IntDemo()
	reflectValue2FloatDemo()
	reflectValue2StringDemo()
	reflectValue2BoolDemo()
	reflectValue2PointerDemo()
	reflectValue2InterfaceDemo()
}

//反射value转int
func reflectValue2IntDemo() {
	fmt.Printf("\n**********reflectValue2IntDemo**********\n")
	var age int = 25

	v1 := reflect.ValueOf(age)
	fmt.Printf("转换前， type: %T, value: %v \n", v1, v1)
	v2 := v1.Int()
	fmt.Printf("转换后，type：%T, value：%v \n", v2, v2)
}

//反射value转float
func reflectValue2FloatDemo() {
	fmt.Printf("\n**********reflectValue2FloatDemo**********\n")

	var score float64 = 99.5

	v1 := reflect.ValueOf(score)
	fmt.Printf("转换前， type: %T, value: %v \n", v1, v1)
	v2 := v1.Float()
	fmt.Printf("转换后，type：%T, value：%v \n", v2, v2)
}

//反射Value转string
func reflectValue2StringDemo() {
	fmt.Printf("\n**********reflectValue2StringDemo**********\n")

	var name string = "Hello"

	v1 := reflect.ValueOf(name)
	fmt.Printf("转换前， type: %T, value: %v \n", v1, v1)
	v2 := v1.String()
	fmt.Printf("转换后，type：%T, value：%v \n", v2, v2)
}

//反射Value转bool
func reflectValue2BoolDemo() {
	fmt.Printf("\n**********reflectreflectValue2BoolDemoValue2StringDemo**********\n")

	var isMale bool = false

	v1 := reflect.ValueOf(isMale)
	fmt.Printf("转换前， type: %T, value: %v \n", v1, v1)
	v2 := v1.Bool()
	fmt.Printf("转换后，type：%T, value：%v \n", v2, v2)
}

//反射Value转指针
func reflectValue2PointerDemo() {
	fmt.Printf("\n**********reflectValue2PointerDemo**********\n")

	var age int = 25

	v1 := reflect.ValueOf(&age)
	fmt.Printf("转换前， type: %T, value: %v \n", v1, v1)
	v2 := v1.Pointer()
	fmt.Printf("转换后，type：%T, value：%v \n", v2, v2)
}

//反射Value转接口
func reflectValue2InterfaceDemo() {
	fmt.Printf("\n**********reflectValue2InterfaceDemo**********\n")
	var age int = 25

	v1 := reflect.ValueOf(age)
	fmt.Printf("转换前， type: %T, value: %v \n", v1, v1)
	v2 := v1.Interface()
	fmt.Printf("转换后，type：%T, value：%v \n", v2, v2)
}

//切片的反射Demo
func sliceReflectDemo() {
	fmt.Printf("\n**********sliceReflectDemo**********\n")

	reflectValue2Slice()
	updateSlice()
}

//反射Value转切片
func reflectValue2Slice() {
	fmt.Printf("\n**********reflectValue2Slice**********\n")
	var numList []int = []int{1, 2, 3, 4}

	v1 := reflect.ValueOf(numList)
	fmt.Printf("转换前， type: %T, value: %v \n", v1, v1)

	// Slice 函数接收两个参数
	v2 := v1.Slice(0, 2)
	v3 := v1.Slice3(0, 2, 3)

	fmt.Printf("转换后，type：%T, value：%v \n", v2, v2)
	fmt.Printf("v2转换后的长度大小：%d，容器大小：%d \n", v2.Len(), v2.Cap())
	fmt.Printf("转换后，type：%T, value：%v \n", v3, v3)
	fmt.Printf("v3转换后的长度大小：%d，容器大小：%d \n", v3.Len(), v3.Cap())
}

//更新切片
func updateSlice() {
	fmt.Printf("\n**********updateSlice**********\n")

	arr := []int{1, 2}
	appendToSlice(&arr)
	fmt.Println(arr)
}

//切片追加
func appendToSlice(arrPtr interface{}) {
	valuePtr := reflect.ValueOf(arrPtr)
	value := valuePtr.Elem()

	value.Set(reflect.Append(value, reflect.ValueOf(3)))

	fmt.Println(value)
	fmt.Println(value.Len())
}

//属性反射Demo
func fieldsReflectDemo() {
	fmt.Printf("\n**********fieldsReflectDemo**********\n")
	person := Person{
		name:   "小明",
		age:    22,
		gender: "male",
	}

	v := reflect.ValueOf(person)
	fmt.Println("字段数:", v.NumField())
	fmt.Println("第 1 个字段：", v.Field(0))
	fmt.Println("第 2 个字段：", v.Field(1))
	fmt.Println("第 3 个字段：", v.Field(2))

	fmt.Println("==========================")
	// 也可以这样来遍历
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("第 %d 个字段：%v \n", i+1, v.Field(i))
	}
}

func (p Person) SayHello() {
	fmt.Println("hello")
}

func (p Person01) SayHelloReturnString() string {
	return "hello"
}

func (p Person) SayBye() {
	fmt.Println("Bye")
}

func (p Person01) SayByeReturnString() string {
	return "Bye"
}

//方法反射Demo
func methodsReflectDemo() {
	fmt.Printf("\n**********methodsReflectDemo**********\n")
	p := &Person{}

	t := reflect.TypeOf(p)

	fmt.Println("方法数（可导出的）:", t.NumMethod())
	fmt.Println("第 1 个方法：", t.Method(0).Name)
	fmt.Println("第 2 个方法：", t.Method(1).Name)

	fmt.Println("==========================")
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Printf("第 %d 个方法：%s \n", i+1, t.Method(i).Name)
	}
}

//动态调用反射的方法
func dynamicCallingReflectFunc() {
	fmt.Printf("\n**********dynamicCallingReflectFunc**********\n")
	dynamicCallingReflectFuncDemo01()
	dynamicCallingReflectFuncDemo02()
	dynamicCallingReflectFuncDemo03()
}

//动态调用函数（使用索引且无参数）
func dynamicCallingReflectFuncDemo01() {
	fmt.Printf("\n**********dynamicCallingReflectFuncDemo01**********\n")
	//要调用 Call，注意要使用 ValueOf

	p := Person01{}

	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	for i := 0; i < v.NumMethod(); i++ {
		fmt.Printf("调用第 %d 个方法，方法名为 %v，调用结果为：%v \n",
			i+1,
			t.Method(i).Name,
			v.Method(i).Call(nil),
		)
	}
}

//动态调用函数（使用函数名且无参数）
func dynamicCallingReflectFuncDemo02() {
	fmt.Printf("\n**********dynamicCallingReflectFuncDemo02**********\n")

	p := Person{}

	v := reflect.ValueOf(p)

	v.MethodByName("SayHello").Call(nil)
	v.MethodByName("SayBye").Call(nil)
}

func (p *Person) SelfIntroduction(name string, age int) {
	fmt.Printf("Hello, my name is %s and i'm %d years old. \n", name, age)
	p.name = name
	p.age = age
}

//动态调用函数（使用函数且有参数）
func dynamicCallingReflectFuncDemo03() {
	fmt.Printf("\n**********dynamicCallingReflectFuncDemo02**********\n")

	p := &Person{}

	v := reflect.ValueOf(p)
	name := reflect.ValueOf("wangming")
	age := reflect.ValueOf(16)
	input := []reflect.Value{name, age}
	v.MethodByName("SelfIntroduction").Call(input)
	fmt.Printf("SelfIntroduction After Person: %+v", v.Elem())
}

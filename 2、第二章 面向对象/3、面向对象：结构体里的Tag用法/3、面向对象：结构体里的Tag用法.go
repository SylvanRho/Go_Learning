package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name" label:"Name is: "`
	Age  int    `json:"age" label:"Age is: "`
	//只要发现对象里的 Addr 为 false， 0， 空指针，空接口，空数组，空切片，空映射，空字符串中的一种，就会被忽略。
	Addr string `json:"addr,omitempty" label:"Gender is: " default:"unknown"`
}

func main() {
	usedTag()
	p := Person{
		Name: "青",
		Age:  18,
	}
	getTag(p)
	tagDemo()
}

//使用Tag
func usedTag() {
	fmt.Printf("**********usedTag**********\n")
	p1 := Person{
		Name: "Jack",
		Age:  22,
	}

	//json.Marshal加密成json字符串
	data1, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("data1 : %s \n", data1)

	p2 := Person{
		Name: "Jack",
		Age:  22,
		Addr: "China",
	}

	data2, err := json.Marshal(p2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("data2 : %s \n", data2)
}

//获取Tag
func getTag(person Person) {
	fmt.Printf("**********getTag**********\n")
	i := 0
	// 三种获取 field
	field01, _ := reflect.TypeOf(person).FieldByName("Name")
	field02 := reflect.ValueOf(person).Type().Field(i + 1)
	field03 := reflect.ValueOf(&person).Elem().Type().Field(i + 2)

	// 获取 Tag
	tag01 := field01.Tag
	tag02 := field02.Tag
	tag03 := field03.Tag

	// 获取键值对
	// 其实 Get 只是对 Lookup 函数的简单封装而已，当没有获取到对应 tag 的内容，会返回空字符串。
	jsonValue := tag01.Get("json")
	fmt.Println("jsonValue01：", jsonValue)
	//lookup返回的代表是否查询到对应的内容
	jsonValue02, _ := tag02.Lookup("json")
	fmt.Println("jsonValue01：", jsonValue02)

	jsonvalue03 := tag03.Get("json")
	fmt.Println("jsonvalue03：", jsonvalue03)

}

//一个Tag使用的小Demo
func tagDemo() {
	fmt.Printf("**********tagDemo**********\n")
	xm := Person{
		Name: "xiaoming",
		Age:  22,
	}
	Print(xm)
}

func Print(obj interface{}) error {
	//取value
	v := reflect.ValueOf(obj)
	// 解析字段
	//NumField()拿到所有的属性数量
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag

		label := tag.Get("label")
		defaultValue := tag.Get("default")

		value := fmt.Sprintf("%v", v.Field(i))

		if value == "" {
			value = defaultValue
		}
		fmt.Println(label, value)
	}
	return nil
}

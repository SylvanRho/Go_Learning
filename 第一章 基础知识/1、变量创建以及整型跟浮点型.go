package main

import "fmt"

func main() {
	var (
		name   string = "青"
		age    int    = 16
		gender string = "高一"
	)
	fmt.Println(name, age, gender)

	name2 := "ceshi"
	age2, gender2 := 18, "大一"
	fmt.Println(name2, age2, gender2)

	var agePtr = &age2
	println(agePtr)

	ptr2 := new(int)
	println("Address：", ptr2)
	println("value：", *ptr2) // * 后面接指针变量，表示从内存地址中取出值

	println("funcNewInt：", newInt())

	a, _ := GetData()
	_, b := GetData()
	println("a：", a)
	println("b：", b)
}

func newInt() *int {
	return new(int)
}

func GetData() (int, int) {
	return 100, 200
}

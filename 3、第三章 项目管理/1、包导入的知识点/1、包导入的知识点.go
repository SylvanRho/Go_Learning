package main

import (
	//相对路径
	sp "./subpack"
	//使用点 在使用的时候可以不需要再加前缀fmt.了
	. "fmt"
	//假如只想使用包的init()方法，就可以使用 _ 来匿名导入
	_ "image/png"
)

var a int = 10

//init会比main先执行，且无法被调用
//第一个init
func init() {
	a = 20
	Println("init : ", a)
}

//第二个init
func init() {
	a = 30
	Println("init2 : ", a)
}

func main() {
	Println(a)
	//同级目录 可以直接使用
	PrintHello("小明")
	function()

	//不同路径下的需要引用路径
	sp.Sub()
}

//第三个init
func init() {
	a = 40
	Println("init3 : ", a)
}

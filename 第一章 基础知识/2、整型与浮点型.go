package main

import (
	"fmt"
)

func main() {
	// 1、整形
	var num int = 10
	// 2进制：以0b或0B为前缀
	var num01 int = 0b1100
	//8进制：以0o或者 0O为前缀
	var num02 int = 0o14
	// 16进制：以0x 为前缀
	var num03 int = 0xc

	fmt.Println("10进制：", num)
	fmt.Printf("2进制数 %b 表示的是: %d \n", num01, num01)
	fmt.Printf("8进制数 %o 表示的是: %d \n", num02, num02)
	fmt.Printf("16进制数 %X 表示的是: %d \n", num03, num03)

	// %b    表示为二进制
	// %c    该值对应的unicode码值
	// %d    表示为十进制
	// %o    表示为八进制
	// %q    该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	// %x    表示为十六进制，使用a-f
	// %X    表示为十六进制，使用A-F
	// %U    表示为Unicode格式：U+1234，等价于"U+%04X"
	// %E    用科学计数法表示
	// %f    用浮点数表示

	// 2、浮点型Float32 Float64

	// 常量 math.MaxFloat32 表示 float32 能取到的最大数值，大约是 3.4e38；
	// 常量 math.MaxFloat64 表示 float64 能取到的最大数值，大约是 1.8e308；
	// float32 和 float64 能表示的最小值分别为 1.4e-45 和 4.9e-324。

	// float32的精度只能提供大约6个十进制数（表示后科学计数法后，小数点后6位）的精度
	// float64的精度能提供大约15个十进制数（表示后科学计数法后，小数点后15位）的精度

	var myfloat float32 = 10000018
	fmt.Println("myfloat:", myfloat)
	fmt.Println("myfloat+1:", myfloat+1) //小数点后7位还可以比较出来准确性

	var myfloat01 float32 = 100000182
	var myfloat02 float32 = 100000183
	fmt.Println("myfloat01: ", myfloat01)
	fmt.Println("myfloat01+1: ", myfloat01+1)
	fmt.Println(myfloat02 == myfloat01+1) //小数点7位以后 将不准确 返回为True
	//但是只有+1会存在 +5就是False了
	//100000182 +5 跟 100000187比较就是False
}

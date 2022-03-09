package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// byte 1个字节 8个比特位
	// 在 ASCII 表中，由于字母 A 的ASCII 的编号为 65 ，字母 B 的ASCII 编号为 66
	var a byte = 65
	// 8进制写法: var a byte = '\101'     其中 \ 是固定前缀
	// 16进制写法: var a byte = '\x41'    其中 \x 是固定前缀

	var b uint8 = 66
	fmt.Printf("a 的值: %c \nb 的值: %c\n", a, b)

	// 或者使用 string 函数
	fmt.Println("a 的值: ", string(a), " \nb 的值: ", string(b))

	// rune 4个字节 32个比特位
	var a01 byte = 'A'
	var b01 rune = 'B'
	var name rune = '中'
	fmt.Printf("a01 占用 %d 个字节数\nb01 占用 %d 个字节数\n", unsafe.Sizeof(a01), unsafe.Sizeof(b01))
	fmt.Printf("中 10进制是：%d \n", name)

	//字符串
	//Go 语言的 string 是用 uft-8 进行编码的，英文字母占用一个字节，而中文字母占用 3个字节
	var mystr01 string = "hello"
	var mystr02 [5]byte = [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("mystr01：%s\n", mystr01)
	fmt.Printf("mystr02：%s\n", mystr02)
	//mystr01 和 mystr02 输出一样，说明了 string 的本质，其实是一个 byte数组
	var country string = "hello,中国"
	fmt.Printf("contry字节长度为：%d \n", len(country)) //5 + 1 + 3 * 2 = 12

	// 比如我想表示 \r\n 这个 字符串，使用双引号是这样写的，这种叫解释型表示法
	// var mystr01 string = "\\r\\n"
	// 而使用反引号，就方便多了，所见即所得，这种叫原生型表示法
	// var mystr02 string = `\r\n`
	// 你仍然想使用解释型的字符串，但是各种转义实在太麻烦了。你可以使用 fmt 的 %q 来还原一下。

	var mystr03 string = `\r\n`
	fmt.Print(`\r\n`)
	fmt.Printf("的解释型字符串是： %q", mystr03)
}

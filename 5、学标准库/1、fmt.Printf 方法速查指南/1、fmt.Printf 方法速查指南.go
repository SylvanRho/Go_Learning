package main

import "fmt"

func main() {
	/*	fmt打印三大函数对比
		1、fmt.Print：正常打印字符串和变量，不会进行格式化，不会自动换行，需要手动添加 \n 进行换行，多个变量值之间不会添加空格
		2、fmt.Println：正常打印字符串和变量，不会进行格式化，多个变量值之间会添加空格，并且在每个变量值后面会进行自动换行
		3、fmt.Printf：可以按照自己需求对变量进行格式化打印。需要手动添加 \n 进行换行
	*/

	fmt.Print("hello", "world\n") //output:helloworld
	fmt.Println("hello", "world") //output:hello world
	fmt.Printf("hello world\n")   //output:hello world

	fmtPrintPlaceholder()
	fmtPrintBool()
	fmtPrintStr()
	fmtPrintPointer()
	fmtPrintInteger()
	fmtPrintFloat()
	fmtPrintWidthIdentifier()
	fmtPrintIdentifier01()
	fmtPrintIdentifier02()
	alignTheCompletion()
	plusOrMinusAlign()
}

//人
type Profile struct {
	name   string
	gender string
	age    int
}

//打印占位符
func fmtPrintPlaceholder() {
	fmt.Printf("**********fmtPrintPlaceholder**********\n")
	var people = Profile{name: "xiaoming", gender: "male", age: 18}
	/*	通用占位符的语义
		%v：以值的默认格式打印
		%+v：类似%v，但输出结构体时会添加字段名
		%#v：值的Go语法表示
		%T：打印值的类型
		%%： 打印百分号本身
	*/
	fmt.Printf("%v \n", people) // output:{xiaoming male 18}
	fmt.Printf("%T \n", people) // output:main.Profile

	// 打印结构体名和类型
	fmt.Printf("%+v \n", people) // output:{name:xiaoming gender:male age:18}
	fmt.Printf("%#v \n", people) // output:main.Profile{name:"xiaoming", gender:"male", age:18}
	fmt.Printf("%% \n")          // output:%
}

//打印布尔值
func fmtPrintBool() {
	fmt.Printf("**********fmtPrintBool**********\n")
	fmt.Printf("%t \n", true)  //output: true
	fmt.Printf("%t \n", false) //output: false
}

//打印字符串
func fmtPrintStr() {
	fmt.Printf("**********fmtPrintStr**********\n")
	/*	打印字符串的定义
		%s：输出字符串表示（string类型或[]byte)
		%q：双引号围绕的字符串，由Go语法安全地转义
		%x：十六进制，小写字母，每字节两个字符
		%X：十六进制，大写字母，每字节两个字符
	*/

	fmt.Printf("%s \n", []byte("Hello, Golang")) // output:Hello, Golang
	fmt.Printf("%s \n", "Hello, Golang")         // output:Hello, Golang

	fmt.Printf("%q \n", []byte("Hello, Golang")) // output: "Hello, Golang"
	fmt.Printf("%q \n", "Hello, Golang")         // output: "Hello, Golang"
	fmt.Printf("%q \n", `Hello, \r\n Golang`)    // output:"Hello, \\r\\n Golang"

	fmt.Printf("%x \n", "Hello, Golang") // output:48656c6c6f2c20476f6c616e67
	fmt.Printf("%X \n", "Hello, Golang") // output:48656C6C6F2C20476F6C616E67
}

//打印指针
func fmtPrintPointer() {
	fmt.Printf("**********fmtPrintPointer**********\n")
	people := Profile{name: "xiaoming", gender: "male", age: 18}
	//%p	表示为十六进制，并加上前导的0x
	fmt.Printf("%p \n", &people)
}

//打印整型
func fmtPrintInteger() {
	fmt.Printf("**********fmtPrintInteger**********\n")
	/*	打印整型的定义
		%b：以二进制打印
		%d：以十进制打印
		%o：以八进制打印
		%x：以十六进制打印，使用小写：a-f
		%X：以十六进制打印，使用大写：A-F
		%c：打印对应的的unicode码值
		%q：该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
		%U：该值对应的 Unicode格式：U+1234，等价于”U+%04X”
	*/
	n := 1024
	fmt.Printf("%d 的 2 进制：%b \n", n, n)
	fmt.Printf("%d 的 8 进制：%o \n", n, n)
	fmt.Printf("%d 的 10 进制：%d \n", n, n)
	fmt.Printf("%d 的 16 进制：%x \n", n, n)

	// 将 10 进制的整型转成 16 进制打印： %x 为小写， %X 为大写
	fmt.Printf("%x \n", 1024)
	fmt.Printf("%X \n", 1024)

	// 根据 Unicode码值打印字符
	fmt.Printf("ASCII 编码为%d 表示的字符是： %c \n", 65, 65) // output: A

	// 根据 Unicode 编码打印字符
	fmt.Printf("%c \n", 0x4E2D) // output: 中

	// 打印 raw 字符时
	fmt.Printf("%q \n", 0x4E2D) // output: '中'

	// 打印 Unicode 编码
	fmt.Printf("%U \n", '中') // output: U+4E2D
}

//打印浮点型
func fmtPrintFloat() {
	fmt.Printf("**********fmtPrintFloat**********\n")
	/*	打印浮点型的定义
		%b：无小数部分、二进制指数的科学计数法，如-123456p-78
		%e：科学计数法，如-1234.456e+78
		%E：科学计数法，如-1234.456E+78
		%f：有小数部分但无指数部分，如123.456
		%F：等价于%f
		%g：根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
		%G：根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
	*/

	f := 12.34
	fmt.Printf("%b\n", f) // output: 6946802425218990p-49
	fmt.Printf("%e\n", f) // output: 1.234000e+01
	fmt.Printf("%E\n", f) // output: 1.234000E+01
	fmt.Printf("%f\n", f) // output: 12.340000
	fmt.Printf("%g\n", f) // output: 12.34
	fmt.Printf("%G\n", f) // output: 12.34
}

//打印宽度标识符
func fmtPrintWidthIdentifier() {
	fmt.Printf("**********fmtPrintWidthIdentifier**********\n")
	/*	打印宽度标识符的定义
		%f	默认宽度，默认精度
		%9f	宽度9，默认精度
		%.2f	默认宽度，精度2
		%9.2f	宽度9，精度2
		%9.f	宽度9，精度0
	*/

	n := 12.34
	fmt.Printf("%f \n", n)    //output: 12.340000
	fmt.Printf("%9f \n", n)   //output: 12.340000
	fmt.Printf("%.2f \n", n)  //output: 12.34
	fmt.Printf("%9.2f \n", n) //output:     12.34
	fmt.Printf("%9.f \n", n)  //output:        12
}

//打印占位符 %+
func fmtPrintIdentifier01() {
	fmt.Printf("**********fmtPrintIdentifier01**********\n")

	/*	%+占位符的定义
		%+v：若值为结构体，则输出将包括结构体的字段名。
		%+q：保证只输出ASCII编码的字符，非 ASCII 字符则以unicode编码表示
	*/

	// 若值为结构体，则输出将包括结构体的字段名。
	var people = Profile{name: "wangbm", gender: "male", age: 27}
	fmt.Printf("%v \n", people)  // output: {name:wangbm gender:male age:27}
	fmt.Printf("%+v \n", people) // output: {name:wangbm gender:male age:27}

	// 保证只输出ASCII编码的字符
	fmt.Printf("%q \n", "golang")  // output: "golang"
	fmt.Printf("%+q \n", "golang") // output: "golang"

	// 非 ASCII 字符则以unicode编码表示
	fmt.Printf("%q \n", "中文")  // output: "中文"
	fmt.Printf("%+q \n", "中文") // output: "\u4e2d\u6587"
}

//打印占位符 %#
func fmtPrintIdentifier02() {
	fmt.Printf("**********fmtPrintIdentifier01**********\n")
	/*	%#占位符的定义
		%#x：给打印出来的是 16 进制字符串加前缀 0x
		%#q：用反引号包含，打印原始字符串
		%#U：若是可打印的字符，则将其打印出来
		%#p：若是打印指针的内存地址，则去掉前缀 0x
	*/

	// 对于打印出来的是 16 进制，则加前缀 0x
	fmt.Printf("%x \n", "Hello, Golang")  // output: 48656c6c6f2c20476f6c616e67
	fmt.Printf("%#x \n", "Hello, Golang") // output: 0x48656c6c6f2c20476f6c616e67

	// 用反引号包含，打印原始字符串
	fmt.Printf("%q \n", "Hello, Golang")  // output: "Hello, Golang"
	fmt.Printf("%#q \n", "Hello, Golang") // output: `Hello, Golang`

	// 若是可打印的字符，则将其打印出来
	fmt.Printf("%U \n", '中')  // output: U+4E2D
	fmt.Printf("%#U \n", '中') // output: U+4E2D '中'

	// 若是打印指针的内存地址，则去掉前缀 0x
	a := 1024
	fmt.Printf("%p \n", &a)  // output: 0xc0000160e0
	fmt.Printf("%#p \n", &a) // output: c0000160e0
}

//对齐补全
func alignTheCompletion() {
	fmt.Printf("**********alignTheCompletion**********\n")
	alignTheCompletionString()

	alignTheCompletionFloat()
}

//对齐补全字符串
func alignTheCompletionString() {
	fmt.Printf("**********alignTheCompletionString**********\n")
	// 打印的值宽度为5，若不足5个字符，则在前面补空格凑足5个字符。
	fmt.Printf("a%5sc\n", "b") // output: a    bc
	// 打印的值宽度为5，若不足5个字符，则在后面补空格凑足5个字符。
	fmt.Printf("a%-5sc\n", "b") // output: ab    c

	// 不想用空格补全，还可以指定0，其他数值不可以，注意：只能在前边补全，后边补全无法指定字符
	fmt.Printf("a%05sc\n", "b") // output: a0000bc
	// 若超过5个字符，不会截断
	fmt.Printf("a%05sc\n", "bbbccc") // output: abbbcccc
}

//对齐补全浮点数
func alignTheCompletionFloat() {
	fmt.Printf("**********alignTheCompletionFloat**********\n")
	// 保证宽度为6（包含小数点)，2位小数，右对齐
	// 不足6位时，整数部分空格补全，小数部分补零，超过6位时，小数部分四舍五入
	fmt.Printf("%6.2f,%6.2f\n", 12.34, 123.456789)

	// 保证宽度为6（包含小数点)，2位小数，- 表示左对齐
	// 不足6位时，整数部分空格补全，小数部分补零，超过6位时，小数部分四舍五入
	fmt.Printf("%-6.2f,%-6.2f\n", 12.34, 123.456789)
}

//正负号占位
func plusOrMinusAlign() {
	fmt.Printf("**********plusOrMinusAlign**********\n")
	//如果是正数，则留一个空格，表示正数
	//如果是负数，则在此位置，用 - 表示
	fmt.Printf("1% d3\n", 222)
	fmt.Printf("1% d3\n", -222)
}

# Go_Learning

这是学习Go的代码库

## fmt Printf语法
+ %b    表示为二进制
+ %c    该值对应的unicode码值
+ %d    表示为十进制
+ %o    表示为八进制
+ %q    该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
+ %x    表示为十六进制，使用a-f
+ %X    表示为十六进制，使用A-F
+ %U    表示为Unicode格式：U+1234，等价于"U+%04X"
+ %E    用科学计数法表示
+ %f    用浮点数表示

## 代码示例
``` go
// 打印结构体
p := point{1, 2}
fmt.Printf("%v\n", p) //输出结果为 {1 2}

// 如果值是一个结构体，%+v 变体将包括结构体的字段名。
fmt.Printf("%+v\n", p) //输出结果为 {x:1 y:2}

// %#v 变体打印值的 Go 语法表示，即将生成该值的源代码片段。
fmt.Printf("%#v\n", p) //输出结果为 main.point{x:1, y:2}

// 打印类型
fmt.Printf("%T\n", p) //输出结果为 main.point

// 打印布尔值
fmt.Printf("%t\n", true) //输出结果为 true

// 打印整数。
fmt.Printf("%d\n", 123) //输出结果为 123

// 打印二进制
fmt.Printf("%b\n", 14) //输出结果为 1110

// 打印字符
fmt.Printf("%c\n", 33) //输出结果为 ！

// 打印 16 进制
fmt.Printf("%x\n", 456) //输出结果为 1c8

// 打印浮点数
fmt.Printf("%f\n", 78.9) //输出结果为 78.900000
// 指数型
fmt.Printf("%e\n", 123400000.0) //输出结果为 1.234000e+08
fmt.Printf("%E\n", 123400000.0) //输出结果为 1.234000E+08

// 字符串
fmt.Printf("%s\n", "\"string\"") //输出结果为 "string"
// 双引号字符串.
fmt.Printf("%q\n", "\"string\"") //输出结果为 "\"string\""

// 每个字符用两位 16 进制表示。
fmt.Printf("%x\n", "hex this") //输出结果为 6865782074686973

// 打印指针`.
fmt.Printf("%p\n", &p) //输出结果为 0xc00001c08

// 字符宽度
fmt.Printf("|%6d|%6d|\n", 12, 345) //输出结果为 |    12|   345|

//字符精度
fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45) //输出结果为 |  1.20|  3.45|

// 左对齐
fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) //输出结果为 |1.20  |3.45  |

//同样可以控制字符的宽度
fmt.Printf("|%6s|%6s|\n", "foo", "b") //输出结果为 |   foo|     b|

// 同样字符左对齐.
fmt.Printf("|%-6s|%-6s|\n", "foo", "b") //输出结果为 |foo   |b     |

// 合并
s := fmt.Sprintf("a %s", "string") //输出结果为 a string
fmt.Println(s) //输出结果为
// 同样的效果
fmt.Fprintf(os.Stderr, "an %s\n", "error") //输出结果为 an error
```

## 特有名词
``` go
var LintGonicMapper = GonicMapper{
    "API":   true,
    "ASCII": true,
    "CPU":   true,
    "CSS":   true,
    "DNS":   true,
    "EOF":   true,
    "GUID":  true,
    "HTML":  true,
    "HTTP":  true,
    "HTTPS": true,
    "ID":    true,
    "IP":    true,
    "JSON":  true,
    "LHS":   true,
    "QPS":   true,
    "RAM":   true,
    "RHS":   true,
    "RPC":   true,
    "SLA":   true,
    "SMTP":  true,
    "SSH":   true,
    "TLS":   true,
    "TTL":   true,
    "UI":    true,
    "UID":   true,
    "UUID":  true,
    "URI":   true,
    "URL":   true,
    "UTF8":  true,
    "VM":    true,
    "XML":   true,
    "XSRF":  true,
    "XSS":   true,
```

# 暂时不理解的问题
- [ ] 为什么sync.Mutex传递的时候要加*号
- [x] 为什么切片在循环里面 append数据不会报错 且不会改变数据大小 循环结束以后是改变后的数据
    - range 返回的是每个元素的副本，而不是直接返回对该元素的引用
- [ ] 为什么 for range 信道 只要不close 会一直等待发送方发消息
package main

import (
	"fmt"
)

func main() {
	initStruct()
	inheritanceByStruct()

	//实例化
	createStructDemo01()
	createStructDemo02()
	createStructDemo03()

	selectorDemo()
}

//初始化结构体
func initStruct() {
	fmt.Printf("**********initStruct**********\n")
	//声明结构体
	// type 结构体名 struct {
	// 	属性名   属性类型
	// 	属性名   属性类型
	// 	...
	// }
	xmFather := Profile{
		name:   "小明父亲",
		age:    35,
		gender: "male", //如果没有跟结尾的花括号在同一行 就要加逗号
	}
	// 初始化结构体，并不一定要所有字段都赋值，未被赋值的字段，会自动赋值为其类型的零值。
	xmMother := Profile{name: "小明母亲", age: 34, gender: "female"}
	xm := Profile{name: "小明", age: 18, gender: "male", mother: &xmMother, father: &xmFather}
	fmt.Printf("小明母亲的名称：%s \n", xm.mother.name)
	fmt.Println("小明母亲的母亲：", xmMother.mother)
	xm.FmtProfile()
	xm.increase_age()
	xm.FmtProfile()
}

//人的结构体
type Profile struct {
	name, gender string
	age          int
	mother       *Profile // 指针
	father       *Profile // 指针
}

//实体绑定方法
func (person Profile) FmtProfile() {
	fmt.Printf("**********FmtProfile**********\n")
	fmt.Printf("名字：%s\n", person.name)
	fmt.Printf("年龄：%d\n", person.age)
	fmt.Printf("性别：%s\n", person.gender)
}

//假如要在方法里更改对象的属性值  类型要加*号
func (person *Profile) increase_age() {
	fmt.Printf("**********increase_age**********\n")
	person.age += 1
}

//结构体的"继承"
func inheritanceByStruct() {
	fmt.Printf("**********inheritanceByStruct**********\n")

	myCom := company{
		companyName: "Tencent",
		companyAddr: "深圳市南山区",
	}
	staffInfo := staff{
		name:     "小明",
		age:      28,
		gender:   "男",
		position: "云计算开发工程师",
		company:  myCom,
	}
	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.companyName)
	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.company.companyName)
}

//电脑
type company struct {
	companyName string
	companyAddr string
}

//职位
type staff struct {
	name     string
	age      int
	gender   string
	position string
	//将 company 这个 结构体嵌入到 staff 中，做为 staff 的一个匿名字段，staff 就直接拥有了 company 的所有属性了。
	company // 匿名字段
}

//实例化
//第一种：正常实例化
func createStructDemo01() {
	fmt.Printf("**********createStructDemo01**********\n")
	xm := Profile{
		name:   "小明",
		age:    18,
		gender: "male",
	}
	fmt.Println(xm)
}

//第二种：使用 new
func createStructDemo02() {
	fmt.Printf("**********createStructDemo02**********\n")
	// 等价于: var xm *Profile = new(Profile)
	xm := new(Profile)
	// output: &{ 0 }
	fmt.Println(xm)

	(*xm).name = "iswbm"
	//与上面同等的
	xm.age = 18
	xm.gender = "male"
	fmt.Println(xm)
}

//第三种：使用 &
func createStructDemo03() {
	fmt.Printf("**********createStructDemo03**********\n")
	var xm *Profile = &Profile{}
	fmt.Println(xm)
	// output: &{ 0 }

	xm.name = "iswbm"  // 或者 (*xm).name = "iswbm"
	xm.age = 18        //  或者 (*xm).age = 18
	xm.gender = "male" // 或者 (*xm).gender = "male"
	fmt.Println(xm)
}

//选择器的Demo
func selectorDemo() {
	fmt.Printf("**********selectorDemo**********\n")
	p1 := &Profile{name: "iswbm"}
	//当你对象是结构体对象的指针时，你想要获取字段属性时，按照常规理解应该这么做
	fmt.Println((*p1).name)
	//两种写法一样
	//当对象是结构体对象的指针时，可以直接省去*取值的操作，选择器"."会直接解引用
	fmt.Println(p1.name)
}

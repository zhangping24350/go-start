// Package basicgrammar 文件夹和包名可以不一样
package basicgrammar

import (
	"errors"
	"fmt"
	"runtime"
	"unicode/utf8"

	//匿名引用
	_ "net/http"
)

//变量大小写控制访问权限 大写可以在包外访问 小写不行
// 定义一个名称为 “variableName”，类型为 "type" 的变量
//var variableName type
//变量声明方式一 var
//什么在包中的为 全局变量 跟js类似
var name = "李四"
var Name = "张三"
var Name1 string = "张三"

//var 多个变量写法
//var vname1, vname2, vname3 type= v1, v2, v3
var age1, phone1 = 17, "f1111"

var (
	age   = 18
	phone = "17578972"
)

// 函数大小写控制访问权限 大写可以在包外访问 小写不行
func variableDeclarations(age int16) {
	//变量声明方式二 局部变量
	//覆盖name
	name := "王五"
	//多个变量
	//vname1, vname2, vname3 := v1, v2, v3
	age, phone = 17, "111"
	//匿名变量 _ 符号
	_, email := "dd", "@qq.com"

	fmt.Printf("%s", email)
	fmt.Printf(name)
	//变量声明方式三
	fmt.Printf("ss")
	//常量申明
	const Sex = "常量"

	//Go 内置错误处理
	err := errors.New("emit macho dwarf: elf header corrupted")
	if nil != err {
		println(err.Error())
		println(err)
	}

	//数组 var arr [n]type
	//数组申明方式一
	var arr [2]int
	arr[0] = 12
	arr[1] = 13
	//数组申明方式二
	arr1 := [2]int{1, 2}
	fmt.Printf("%v", arr1)
	//二维数组
	var doubleArr [2][3]int
	doubleArr[0][0] = 22
	doubleArr[0][1] = 23
	doubleArr1 := [2][1]int{{2}, {3}}
	fmt.Printf("%v", doubleArr1)

	//切片slice java ArrayList
	var fslice []int
	var fslice2 []int
	fslice = append(fslice, 2)
	i := len(fslice)
	println("最大长度", i)
	i2 := cap(fslice)
	println("最大容量", i2)
	i3 := copy(fslice2, fslice)
	println("fslice2长度", i3)
	// 声明一个含有 10 个元素元素类型为 byte 的数组
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个含有 byte 的 slice
	var a, b []byte
	// a 指向数组的第 3 个元素开始，并到第五个元素结束，
	a = ar[2:5]
	// 现在 a 含有的元素: ar[2]、ar[3]和ar[4]
	// b 是数组 ar 的另一个 slice
	b = ar[3:5]
	// b 的元素是：ar[3] 和 ar[4]
	println(a)
	println(b)

	//map 格式为 map[keyType]valueType
	//map 申明方式一
	//var numbers map[string]int
	//map 申明方式二
	numbers := make(map[string]int)
	numbers["one"] = 1  // 赋值
	numbers["ten"] = 10 // 赋值
	// 初始化一个map map 也是一种引用类型
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	fmt.Printf("%s", rating["Python"])
	//删除key
	delete(rating, "Python")

	//make 用于内建类型（map、slice 和 channel）的内存分配
	//new 用于各种类型的内存分配

}

func VariableDeclarations(name string, age int16) {
	//变量声明方式一
	//变量声明方式二
	//变量声明方式三
	fmt.Printf("ss")
}

func main() {
	println(`
	反引号可以
	换行
	换行
`)

	//Go len计算的是字节数
	println(len("你好"))
	println(utf8.RuneCountInString("你好"))

	fmt.Printf("fdsjkjfd]%s")

	err := errors.New("emit macho dwarf: elf header corrupted")
	if nil != err {
		println(err.Error())
		println(err)
	}

	//切片slice java ArrayList
	var fslice []int
	var fslice2 []int
	fslice = append(fslice, 2)
	i := len(fslice)
	println("最大长度", i)
	i2 := cap(fslice)
	println("最大容量", i2)
	copy(fslice2, fslice)
	//println("fslice2长度", fslice2[0])
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个含有 byte 的 slice
	var a, b []byte
	// a 指向数组的第 3 个元素开始，并到第五个元素结束，
	a = ar[2:5]
	// 现在 a 含有的元素: ar[2]、ar[3]和ar[4]
	// b 是数组 ar 的另一个 slice
	b = ar[3:5]
	copy(a, b)
	fmt.Printf("%v ==== %v", a, b)
	//runtime.Gosched()
	num1 := runtime.NumGoroutine()
	println(num1)
	num := runtime.GOMAXPROCS(16)
	println(num)
	println(runtime.NumCPU())

}

//数据类型 放在 builtin.go 文件中
func DataType() {
	//string
	//int8 int16 int32 in64 int
	//无符号数
	//uint8 uint16 uint32 uint64 uint
	//float32 float64
	//bool
	//struct 结构体 相当于java bean
	//rune    type rune = int32
	//byte    type byte = uint8
	//uintptr uintptr:用于指针运算

	//数据类型的零值
	//int     0
	//int8    0
	//int32   0
	//int64   0
	//uint    0x0
	//rune    0 // rune 的实际类型是 int32
	//byte    0x0 // byte 的实际类型是 uint8
	//float32 0 // 长度为 4 byte
	//float64 0 // 长度为 8 byte
	//bool    false
	//string  ""

	//数组 var arr [n]type
	var arr [3]int
	arr1 := [3]int{1, 1, 1}
	fmt.Printf("%v", arr)
	fmt.Printf("%v", arr1)
}

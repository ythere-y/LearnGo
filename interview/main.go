package interview

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func Main() {
	//pointer()
	//multipleReturn()
	//errorType()
	//stringConcat()
	//runeType()
	//defaultParameters()
	//deferExecution()
	//valueSwap()
	//getTags()
	//slicesEqual()
	//differentVPrint()
	//enumsInGo()
	emptyStruct()
}

// region reflect

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
func Title(i interface{}) {
	fmt.Printf("[___  %+v  ___]\n", GetFunctionName(i))
}
func Tail() {
	fmt.Println()
}

// endregion

// region pointer
func pointer() {
	Title(pointer)
	var x = 5
	var p = &x
	fmt.Printf("x = %d\n", *p)
	Tail()
}

// endregion

// region multiple return
func swap(x, y string) (string, string) {
	return y, x
}
func multipleReturn() {
	Title(multipleReturn)
	a, b := swap("A", "B")
	fmt.Println(a, b)
	Tail()
}

// endregion

// region error type

// Go 没有异常类型，只有错误类型(Error)，通常用返回值来表示异常状态

// 也可以用errors.New()来定义自己的异常。
// errors.Error()会返回异常的字符串表示。
// 只要实现error接口就可以定义自己的异常。
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
func New(text string) error {
	return &errorString{text}
}

func errorType() {
	Title(errorType)
	_, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err.Error())
		//log.Fatal(err) avoid use this, because log.Fatal case os.Exit()
	}
	err = New("my error type")
	if err != nil {
		fmt.Println(err.Error())
	}
	Tail()
}

// endregion

// region Goroutine
/*
Goroutine 是与其他函数或方法同时运行的函数或方法。
Goroutines 可以被认为是轻量级的线程。
与线程相比，创建 Goroutine 的开销很小。
Go应用程序同时运行数千个 Goroutine 是非常常见的做法。
*/
// endregion

// region string concatenation
/*
"+"
使用+操作符进行拼接时，会对字符串进行遍历，计算并开辟一个新的空间来存储原来的两个字符串。
*/
func catWithPlus(iter int) {
	var str = "+"
	for i := 0; i < iter; i++ {
		str += "+"
	}
	//fmt.Println(str)
}

/*
fmt.Sprintf
由于采用了接口参数，必须要用反射获取值，因此有性能损耗。
*/

func catWithSprintf(iter int) {
	var str string
	for i := 0; i < iter; i++ {
		str = fmt.Sprintf("%vs", str)
	}
	//fmt.Println(str)
}

/*
strings.Builder
用WriteString()进行拼接，内部实现是指针+切片
同时String()返回拼接后的字符串，它是直接把[]byte转换为string，从而避免变量拷贝。

Builder 内部结构简洁，使用切片来存储byte
WriteString就是append
String就是利用强制转换(unsafe.Pointer)来避免内存拷贝
*/
func catWithBuilder(iter int) {
	var str strings.Builder
	for i := 0; i < iter; i++ {
		str.WriteString("b")
	}
	//fmt.Println(str)
}

/*
bytes.Buffer
Buffer 的结构也非常简洁，可以看源码
持续向Buffer尾部写入数据，从Buffer头部读取数据，所以off字段记录读取位置
利用切片cap特性知道写入位置
WriteString多次中，随着写入自动申请空间，追加时采用copy方式，
copy是内置的拷贝函数，可以减少内存分配
String中使用标准类型转为String，所以会发生内存分配，所以性能上不如Builder
*/
func catWithBuff(iter int) {
	var str bytes.Buffer
	for i := 0; i < iter; i++ {
		str.WriteString("f")
	}
	//fmt.Println(str.String())

}

/*
strings.json
是基于builder来实现的，并且可以自定义分隔符
*/

func catWithJson(list []string) {
	var _ = strings.Join(list, "")
	//fmt.Println(str)
}

/*
性能排名：
理论上：
builder = json > buffer > plus > sprintf
我的测试显示:
builder > buffer > json > plus > sprintf
*/
func stringConcat() {
	Title(stringConcat)
	iter := 100000
	clock := time.Now()
	catWithPlus(iter)
	fmt.Printf("plus use time : [%v]\n", time.Now().Sub(clock).String())
	clock = time.Now()
	catWithSprintf(iter)
	fmt.Printf("sprintf use time : [%v]\n", time.Now().Sub(clock).String())
	clock = time.Now()
	catWithBuilder(iter)
	fmt.Printf("builder use time : [%v]\n", time.Now().Sub(clock).String())
	clock = time.Now()
	catWithBuff(iter)
	fmt.Printf("buffer use time : [%v]\n", time.Now().Sub(clock).String())
	list := make([]string, iter)
	for i := range list {
		list[i] = "j"
	}
	clock = time.Now()
	catWithJson(list)
	fmt.Printf("json use time : [%v]\n", time.Now().Sub(clock).String())
	Tail()
}

// endregion string concatenation

// region rune type

/*
// part1
ASCII 码只需要 7 bit 就可以完整地表示，但只能表示英文字母在内的128个字符，
为了表示世界上大部分的文字系统，发明了 Unicode，
Unicode是ASCII的超集，包含世界上书写系统中存在的所有字符，
并为每个代码分配一个标准编号（称为Unicode CodePoint），
在 Go 语言中称之为 rune，是 int32 类型的别名。

Go 语言中，字符串的底层表示是 byte (8 bit) 序列，而非 rune (32 bit) 序列。

// part2
String不能更改，当给String赋其他值的时候，是改变了指针指向的位置
String和[]byte的转化，非常方便
如果发生类型转换之后给其他变量赋值，会发生copy，变量拿到的是一份copy值
slice与slice之间的赋值操作只修改指针指向不copy真实数据
slice的内容是mutable的，String的内容是immutable

// part3
String和[]byte之间强转换，高性能-不安全
`slicebytetostring`
`stringtoslicebyte`
这个两个函数在runtime/string.go中实现，且是私有的，如果要用，需要自己模仿之来实现
如果想直接调用这个两个私有函数，还要用到其他私有类，比较麻烦
*/

func runeType() {
	Title(runeType)
	// part1
	sample := "12我爱你GO"
	runeSample := []rune(sample)
	runeSample[0] = '你'
	fmt.Println(string(runeSample))
	fmt.Println(len(runeSample))

	// part2
	str := "asong"
	fmt.Printf("%p : %v : %v\n", []byte(str), []byte(str), string([]byte(str)))
	str = "song"
	fmt.Printf("%p : %v : %v\n", []byte(str), []byte(str), string([]byte(str)))
	// 观察到打印出来的指针值发生改变

	// part3
	//say := "hello"
	//by := stringtoslicebyte(say)
	//by[0] = 'H'
	//fmt.Printf("say : %v\n", say)
	Tail()
}

// endregion

// region map

func mapTest() {
	Title(mapTest)
	var sample map[int]int
	var idx = 10
	if _, ok := sample[idx]; ok {
		fmt.Printf("have index = %v\n", idx)
	} else {
		fmt.Printf("don't have index = %v\n", idx)
	}
	Tail()
}

// endregion

// region default parameters?

/*

Go不支持默认参数，也不支持可选参数

可以使用结构体参数来实现，或者传入参数切片slice

*/

// 可以由用户自定义的参数集合
type Option struct {
	retryNum int
}

//定义修改默认参数的钩子函数
type ModifyOptFunc func(opt *Option)

//实际修改默认参数的函数
func WithRetryNum(num int) ModifyOptFunc {
	return func(opt *Option) {
		opt.retryNum = num
	}
}

func OperateRedis(modOptions ...ModifyOptFunc) {
	//定义参数并添加默认项
	opt := Option{retryNum: 10}

	//调用钩子函数，并对默认值进行修改
	for _, optFunc := range modOptions {
		optFunc(&opt)
	}

	//使用参数做相关实际逻辑
	fmt.Println("current retry number:", opt.retryNum)
}

func defaultParameters() {
	Title(defaultParameters)
	OperateRedis()
	OperateRedis(WithRetryNum(1))
	OperateRedis(WithRetryNum(14), WithRetryNum(53))
	Tail()
}

// endregion

// region defer's execution order

/*
defer执行顺序和调用顺序相反，类似于栈后进先出(LIFO)。

defer在return之后执行，但在函数退出之前，defer可以修改返回值(仅限于有名返回)
*/

// 无名返回，返回值只说明类型没有说明变量名
// 执行return之后，go会创建一个临时变量来存储返回值，然后运行defer，
// 再之后将临时保存的返回值给到上一层，所以i避开了被修改
func unnamedReturn() int {
	i := 0
	defer func() { fmt.Println("defer 1") }()
	defer func() {
		i += 1
		fmt.Println("defer 2")
	}()
	return i
}

// 有名返回，返回值说明类型和变量名
// 在这种情况下，执行return时不会再创建临时变量保存
func namedReturn() (i int) {
	i = 0
	defer func() { fmt.Println("defer 1") }()
	defer func() {
		i += 1
		fmt.Println("defer 2")
	}()
	return i
}

func deferExecution() {
	Title(deferExecution)
	fmt.Printf("unnamedReturn = %v\n", unnamedReturn())
	fmt.Printf("namedReturn = %v\n", namedReturn())
	Tail()
}

// endregion

// region value swap

func valueSwap() {
	Title(valueSwap)
	b := 'b'
	a := 'a'
	fmt.Printf("at the begining: a = %v, b = %v\n", a, b)
	a, b = b, a
	fmt.Printf("value swap : a = %v, b = %v\n", a, b)
	pA := &a
	pB := &b
	*pA, *pB = *pB, *pA
	fmt.Printf("pointer swap: a = %v, b = %v\n", *pA, *pB)
	Tail()
}

// endregion

// region what's about tag?

/*
Go语言Tag的用处？

tag可以为结构体成员提供属性。常见的：

1. json序列化或反序列化时字段的名称
2. db: sqlx模块中对应的数据库字段名
3. form: gin框架中对应的前端的数据字段名
4. binding: 搭配 form 使用, 默认如果没查找到结构体中的某个字段则不报错值为空, binding为 required 代表没找到返回错误给前端
*/

// endregion

// region how to get tags? reflect

type Author struct {
	Name         int      `json:"name"`
	Publications []string `json:"publications,omitempty"`
}

func getTags() {
	Title(getTags)
	t := reflect.TypeOf(Author{})
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		s, _ := t.FieldByName(name)
		fmt.Printf("name = %v, tag = %v\n", name, s.Tag)
	}

	Tail()
}

// endregion

// region if two slices are equal?

func slicesEqual() {
	Title(slicesEqual)
	a := []string{"h", "e"}
	b := []string{"h", "e"}
	fmt.Printf("compare between [%v] and [%v]\n", a, b)
	if reflect.DeepEqual(a, b) {
		fmt.Printf("equal!\n")
	} else {
		fmt.Printf("not equal!\n")
	}
	Tail()
}

// endregion

// region diff between %v and %+v

/*
%v输出结构体各成员的值；

%+v输出结构体各成员的名称和值；

%#v输出结构体名称和结构体各成员的名称和值

占位符特殊用法美化输出：
%5d 整形长度为5，右对齐，左边留白（使用空格填充）
%-5d 左对齐右留白
%05d 数字前面补零
还有很多其他技巧可见:
https://www.cnblogs.com/forever521Lee/p/10700549.html
*/

func differentVPrint() {
	Title(differentVPrint)
	t := Author{
		Name:         123,
		Publications: []string{"hello", "world"},
	}
	fmt.Printf("the v for %v\n", t)
	fmt.Printf("the +v for %+v\n", t)
	fmt.Printf("the #v for %#v\n", t)
	fmt.Printf("|%v|\n", 12)
	fmt.Printf("|%5v|\n", 12)
	fmt.Printf("|%-5v|\n", 12)
	fmt.Printf("|%05v|\n", 12)
	fmt.Printf("|%015v|\n", 12)
	Tail()
}

// endregion

// region enums in go

/*
参考链接：https://youwu.today/skill/backend/using-enum-in-golang/
参考链接: https://go.dev/ref/spec#Iota
使用`iuto`

iota的用法还可以融入表达式，具体参考链接很有收获

发现：
加上String的实现之后，枚举值自动有2中打印，使用%v会打印出String，使用%d会打印出int


*/

type State int

const (
	Running State = iota
	Pending
	Stopped
)

func (s State) String() string {
	switch s {
	case Running:
		return "Running"
	case Pending:
		return "Pending"
	case Stopped:
		return "Stopped"
	default:
		return "Unknown"
	}
}

func enumsInGo() {
	Title(enumsInGo)
	fmt.Printf("State running: %d, %v, %v\n", Running, Running, Running.String())
	fmt.Printf("State pending: %d, %v, %v\n", Pending, Pending, Pending.String())
	fmt.Printf("State stoped: %d, %v, %v\n", Stopped, Pending, Stopped.String())
	Tail()
}

// endregion

// region what's the use of struct{}?

/*
1. 用map模拟一个set，那么就要把值置为struct{}，struct{}本身不占任何空间，可以避免任何多余的内存分配。
2. 有时候给通道发送一个空结构体,channel<-struct{}{}，也是节省了空间。
3. 仅有方法的结构体
*/

type Set map[string]struct{}

func emptyStruct() {
	// part 1
	set := make(Set)
	for _, item := range []string{"A", "B", "C", "A"} {
		set[item] = struct{}{}
	}
	fmt.Printf("len of set = %v\n", len(set))
	if _, ok := set["A"]; ok {
		fmt.Printf("A exists\n")
	}

	// part 2
	ch := make(chan struct{}, 1)
	go func() {
		<-ch
		// do something
	}()
	ch <- struct{}{}

	// part 3
	type Lamp struct {
	}
}

// endregion

// region does int equals to int32 ?

/*

不是一个概念！
千万不能混淆。
go语言中的int的大小是和操作系统位数相关的，
如果是32位操作系统，int类型的大小就是4字节。
如果是64位操作系统，int类型的大小就是8个字节。

除此之外uint也与操作系统有关。

int8占1个字节，int16占2个字节，int32占4个字节，int64占8个字节。

一个字节是8比特

*/

// endregion

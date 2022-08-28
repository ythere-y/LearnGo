package typeAssertion

import (
	"fmt"
	"reflect"
)

func Main() {

	assertionFail()
	//safeAssertion()
	//differentInt()
	//switchAssertion()
}

type NoMethods interface {
}
type OneMethod interface {
	method()
}
type Son struct {
	age int
}

func (son Son) method() {

}

type Brother struct {
	name string
}

func (brother Brother) method() {

}
func Judge(c interface{}) bool {
	_, ok := c.(OneMethod)
	return ok
}
func assertionFail() {
	// type必须是相符合的
	var str interface{} = "hello world"
	fmt.Printf("before assertion, [value = %v] [type = %v]\n", str, reflect.TypeOf(str))
	v := str.(string)
	fmt.Printf("assertion correct, [value =  %v] [type = %v]\n", v, reflect.TypeOf(v))
	// 下面这种情况运行会panic中断程序
	//{
	//	u := str.(int)
	//	fmt.Printf("can't reach here because of panic %v\n", u)
	//}

	// 只有接口类可以用,下面这种用法直接报错
	//{
	//	var x int = 10
	//	v, ok := x.(int)
	//	fmt.Printf("get x type = %v\n", v)
	//}

	// 方法必须全都实现

	var noM NoMethods = 8
	noMTp, noMOk := noM.(int)
	fmt.Printf("assertion of no methods result : %v,%v\n", noMTp, noMOk)

	// 如果不实现接口，是无法通过编译的，int的类型初始化为0
	var oneM OneMethod
	oneMSonTp, oneMSonOk := oneM.(Son)
	fmt.Printf("assertion of one method : %v,%v\n", oneMSonTp, oneMSonOk)

	//可以重新作为其他东西，string类型的初始值为空
	oneMBroTp, oneMBroOk := oneM.(Brother)
	fmt.Printf("assertion of one method : %v,%v\n", oneMBroTp, oneMBroOk)

	rejudge := new(Son)
	if Judge(rejudge) {
		fmt.Printf("Son is an OneMethod")
	} else {
		fmt.Printf("Son is NOT an OneMethod")
	}

}

func safeAssertion() {
	var str interface{} = "hello"

	s := str.(string)
	fmt.Println(s)

	s, ok := str.(string)
	fmt.Println(s, ok)

	f, ok := str.(float64)
	fmt.Println(f, ok)

	it, ok := str.(int)
	fmt.Println(it, ok)
	if v, ok := str.(int); ok {
		fmt.Printf("str type is : %v\n", reflect.TypeOf(v))
	} else {
		fmt.Printf("str not type of int\n")
	}
}

func differentInt() {
	var x64 interface{} = int64(7)
	j := x64.(int64)
	fmt.Printf("assertion twice = %v\n", reflect.TypeOf(j))
	if i, ok := x64.(int); ok {
		fmt.Printf("assertion once = %v\n", reflect.TypeOf(i))
	} else {
		fmt.Printf("int64 not equal to int\n")
	}

}

func switchAssertion() {
	var i interface{} = "hello"

	switch v := i.(type) {
	case string:
		fmt.Println("x is type of string", v)
	case int:
		fmt.Println("x is type of int", v)
	default:
		fmt.Printf("Unknown type %T!\n", v)
	}
}

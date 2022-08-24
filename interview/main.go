package interview

import (
	"fmt"
	"log"
	"os"
)

func Main() {
	pointer()
	multiple_return()
	errorType()
}

// region pointer
func pointer() {
	var x = 5
	var p = &x
	fmt.Printf("x = %d", *p)
}

// endregion

// region multiple return
func swap(x, y string) (string, string) {
	return y, x
}
func multiple_return() {
	a, b := swap("A", "B")
	fmt.Println(a, b)
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

func errorType() {
	_, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
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

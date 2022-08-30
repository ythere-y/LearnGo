package pointer

import "fmt"

func Main() {
	pointerInitiation()
}

type Student struct {
	Id   int
	Name string
}

var StudentDefault *Student

func pointerInitiation() {
	fmt.Printf("default pointer = %+v\n", StudentDefault)
}

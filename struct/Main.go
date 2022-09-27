package _struct

import "fmt"

type StructTest struct {
	Name  string
	Score int
}

func Main() {
	testP := &StructTest{
		Name:  "joe",
		Score: 123,
	}
	fmt.Printf("test_p = %+v\n", testP)
	testP = &StructTest{
		Name: "hello",
	}
	fmt.Printf("test_p = %+v\n", testP)

}

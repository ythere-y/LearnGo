package concurrency

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
	Age  int
}

func FetchData() {
	var (
		testPerson *Person
		//err        error
	)

	testPerson = &Person{
		Name: "joe",
		Age:  123,
	}
	var waitGroup sync.WaitGroup
	//var m *sync.Mutex
	//var tupleErr *errs.Error
	waitGroup.Add(1)
	fmt.Printf("testPerson = %+v\n", testPerson)

	go func() {
		defer waitGroup.Done()
		testPerson.Age = 312
	}()
	waitGroup.Wait()
	fmt.Printf("testPerson = %+v\n", testPerson)

}

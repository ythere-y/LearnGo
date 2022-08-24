package interview

import "fmt"

func Main() {
	pointer()
	multiple_return()
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

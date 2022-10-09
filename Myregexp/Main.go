package Myregexp

import (
	"fmt"
	"regexp"
)

func Main() {
	testMatchStringStrange()
}

func testMatchStringStrange() {
	{
		// test1
		pattern := "General_Processing Status successful:("
		s := "General_Processing Status successful:(0: )"
		var isOk bool
		if pattern == s {
			isOk = true
		} else {
			isOk, _ = regexp.MatchString(pattern, s)
		}
		fmt.Printf("isOk = %v\n", isOk)
	}
	{
		// test2
		pattern := "j***"
		s := "j***"
		var isOk bool
		isOk, _ = regexp.MatchString(pattern, s)

		fmt.Printf("isOk = %v\n", isOk)
	}
}

package ioTest

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Main() {
	AnyRows()
}

// AnyRows
// 多行数据（不知道有几行），每行固定个数，空格隔开
func AnyRows() {
	var a, b int
	for {
		_, err := fmt.Scan(&a, &b)
		if err == io.EOF {
			break
		}
		fmt.Println(a + b)
	}
}

// NRows
// N 行数据,第一行只有一个数字n，表示后面的行数，其他每行固定个数
func NRows() {
	var n, a, b int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}

//SpecialEnding
// 输入多行数据，每行固定个数，读取到特殊数据（如0，0）时停止
func SpecialEnding() {
	var a, b int
	for {
		fmt.Scan(&a, &b)
		if a == 0 && b == 0 {
			break
		}
		fmt.Println(a + b)
	}
}

//RandomLists
// 第一行一个正整数t，表示数据组数
// 接下来t行，每行一组数据
// 每行第一个整数为整数的个数n ( 1 <= n <= 100)
// 接下来是n个正整数
func RandomLists() {
	var t, n, a int
	fmt.Scan(&t)
	var array [][]int
	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		var cur []int
		for i := 0; i < n; i++ {
			fmt.Scan(&a)
			cur = append(cur, a)
		}
		array = append(array, cur)
	}
	fmt.Printf("array = >\n%v", array)
}

//RandomLineRandomList
// 输入数据有多组，每行表示一组输入数据
// 每行不定有n个整数，空格隔开。( 1 <= n <= 100)。
func RandomLineRandomList() {
	inputs := bufio.NewScanner(os.Stdin)
	var array [][]int
	// 每次读入一行
	for inputs.Scan() {
		// 通过空格将他们分割，并存入一个字符串切片
		data := strings.Split(inputs.Text(), " ")
		var curLine []int
		for i := range data {
			val, _ := strconv.Atoi(data[i])
			curLine = append(curLine, val)
		}
		array = append(array, curLine)
	}
}

// NWords
// 输入有两行，第一行一个正整数n
// 第二行是n个字符串，字符串之间用空格隔开
func NWords() {
	inputs := bufio.NewScanner(os.Stdin)
	n := inputs.Scan()
	fmt.Println("n = ", n)
	for inputs.Scan() {
		str := inputs.Text()
		strList := strings.Split(str, " ")
		sort.Strings(strList)
		fmt.Println(strings.Join(strList, " "))
	}
}

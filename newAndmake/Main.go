package newAndmake

import "fmt"

func Main() {
	testNew()
}

func testNew() {
	{
		var p int
		var v *int
		v = &p
		*v = 11
		fmt.Println(*v)
	}
	// 错误的使用方法，没有空间直接赋值是不能通过编译的
	// 会产生panic错误
	//{
	//	var u *int
	//	*u = 8
	//	fmt.Println(*u)
	//}
	{
		var v *int
		v = new(int)
		*v = 8
		fmt.Println(*v)
	}
	{
		type Name struct {
			P string
		}
		var av *[5]int
		var iv *int
		var sv *string
		var tv *Name

		av = new([5]int)
		fmt.Println(*av) //[0 0 0 0 0 0]
		iv = new(int)
		fmt.Println(*iv) // 0
		sv = new(string)
		fmt.Println(*sv) //
		tv = new(Name)
		fmt.Println(*tv) //{}
	}
	{
		// 对于数组，声明本身会赋初值，new之后会换一个区域
		var a [5]int
		fmt.Printf("a: %p %#v \n", &a, a) //a: 0xc04200a180 [5]int{0, 0, 0, 0, 0}
		av := new([5]int)
		fmt.Printf("av: %p %#v \n", &av, av) //av: 0xc000074018 &[5]int{0, 0, 0, 0, 0}
		// 使用过程中可以不加星号
		(*av)[1] = 8
		av[2] = 10
		fmt.Printf("av: %p %#v \n", &av, av) //av: 0xc000006028 &[5]int{0, 8, 0, 0, 0}
	}
	{
		// 对于slice,使用的时候需要加星号，不带星号就用[]取值会编译不通过
		var a *[]int
		fmt.Printf("a: %p %#v \n", &a, a) //a: 0xc042004028 (*[]int)(nil)
		av := new([]int)
		fmt.Printf("av: %p %#v \n", &av, av) //av: 0xc000074018 &[]int(nil)
		*av = append(*av, 2, 3)
		(*av)[0] = 8
		//av[3] = 10
		fmt.Printf("av: %p %#v \n", &av, av) //panic: runtime error: index out of range
	}
	{
		// map这里，使用new之后，也只是产生一个相当于刚刚声明的map，在make之前无法使用
		var m map[string]string
		fmt.Printf("m: %p %#v \n", &m, m) //m: 0xc042068018 map[string]string(nil)
		mv := new(map[string]string)
		fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc000006028 &map[string]string(nil)
		//(*mv)["a"] = "a"
		//fmt.Printf("mv: %p %#v \n", &mv, mv) //这里会报错panic: assignment to entry in nil map
	}
	{
		// channel也不能直接使用
		cv := new(chan string)
		fmt.Printf("cv: %p %#v \n", &cv, cv) //cv: 0xc000074018 (*chan string)(0xc000074020)
		//cv <- "good" //会报 invalid operation: cv <- "good" (send to non-chan type *chan string)
	}
	fmt.Printf("[[make part]]\n")
	{
		// 对于这4类（slice，map，channel），make之后都能直接使用
		av := make([]int, 5)
		fmt.Printf("av: %p %#v \n", &av, av) //av: 0xc000046400 []int{0, 0, 0, 0, 0}
		av[0] = 1
		fmt.Printf("av: %p %#v \n", &av, av) //av: 0xc000046400 []int{1, 0, 0, 0, 0}
		mv := make(map[string]string)
		fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc000074020 map[string]string{}
		mv["m"] = "m"
		fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc000074020 map[string]string{"m":"m"}
		chv := make(chan string)
		fmt.Printf("chv: %p %#v \n", &chv, chv) //chv: 0xc000074028 (chan string)(0xc00003e060)
		go func(message string) {
			chv <- message // 存消息
		}("Ping!")
		fmt.Println(<-chv) // 取消息 //"Ping!"
		close(chv)
	}
	fmt.Printf("[[combine part]]\n")
	{
		// 先new，让指针有值，然后对于指针指向的内容进行make填充
		var mv *map[string]string
		fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc042004028 (*map[string]string)(nil)
		mv = new(map[string]string)
		fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc000006028 &map[string]string(nil)
		*mv = make(map[string]string)
		(*mv)["a"] = "a"
		fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc042004028 &map[string]string{"a":"a"}
	}
}

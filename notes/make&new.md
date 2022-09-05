# Golang的make和new
> [golang new和make的区别](https://cloud.tencent.com/developer/article/1706196#:~:text=make%E5%92%8Cnew%E9%83%BD%E6%98%AF,%E6%98%AF%E6%8C%87%E5%90%91%E7%B1%BB%E5%9E%8B%E7%9A%84%E6%8C%87%E9%92%88%E3%80%82)

## 总结：
- make和new都是golang用来分配内存的內建函数，且在堆上分配内存，make 即分配内存，也初始化内存。new只是将内存清零，并没有初始化内存。对于slice，map，channel，new操作得到的结果无法直接使用。
- make返回的还是引用类型本身；而new返回的是指向类型的指针。
- make只能用来分配及初始化类型为slice，map，channel的数据；new可以分配任意类型的数据。
- 可以先new得到指针，然后对指向的地址进行make，来初始化一个slice，map，channel的对象的指针。

## new操作
  指针声明之后是nil，是不能直接操作，也不能直接给其赋值的。可以通过已有变量取地址来赋值。
  通过new可以给指针赋值，让其指向一块地址，并且内容会被设置为零值，不同的指针类型的零值不同。
  总结：
  数组虽然是复合类型，但不是引用类型，其他silce、map、channel类型也属于引用类型，go会给引用类型初始化为nil，nil是不能直接赋值的（还需要make）。并且不能用new分配内存。无法直接赋值。
### 普通类型
```go
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
```
### 数组
```go
  // 对于数组，声明本身会赋初值，new之后会换一个区域
  var a [5]int
  fmt.Printf("a: %p %#v \n", &a, a) //a: 0xc04200a180 [5]int{0, 0, 0, 0, 0}
  av := new([5]int)
  fmt.Printf("av: %p %#v \n", &av, av) //av: 0xc000074018 &[5]int{0, 0, 0, 0, 0}
  // 使用过程中可以不加星号
  (*av)[1] = 8
  av[2] = 10
  fmt.Printf("av: %p %#v \n", &av, av) //av: 0xc000006028 &[5]int{0, 8, 0, 0, 0}
````
### Slice
```go
  // 对于slice,使用的时候需要加星号，不带星号就用[]取值会编译不通过
  var a *[]int
  fmt.Printf("a: %p %#v \n", &a, a) //a: 0xc042004028 (*[]int)(nil)
  av := new([]int)
  fmt.Printf("av: %p %#v \n", &av, av) //av: 0xc000074018 &[]int(nil)
  *av = append(*av, 2, 3)
  (*av)[0] = 8
  //av[3] = 10
  fmt.Printf("av: %p %#v \n", &av, av) //panic: runtime error: index out of range
```
### Map
```go
  // map这里，使用new之后，也只是产生一个相当于刚刚声明的map，在make之前无法使用
  var m map[string]string
  fmt.Printf("m: %p %#v \n", &m, m) //m: 0xc042068018 map[string]string(nil)
  mv := new(map[string]string)
  fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc000006028 &map[string]string(nil)
  //(*mv)["a"] = "a"
  //fmt.Printf("mv: %p %#v \n", &mv, mv) //这里会报错panic: assignment to entry in nil map
```
### Channel
```go 
// channel也不能直接使用
cv := new(chan string)
fmt.Printf("cv: %p %#v \n", &cv, cv) //cv: 0xc000074018 (*chan string)(0xc000074020)
//cv <- "good" //会报 invalid operation: cv <- "good" (send to non-chan type *chan string)
```
## make操作
make的基本操作
```go
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
```
## New 和make结合
new操作可以分配指针内容，让指针不再是nil，但是对于slice，map，channel，无法分配空间。把make和new结合，可以分配空间的同时获取地址。
```go
// 先new，让指针有值，然后对于指针指向的内容进行make填充
var mv *map[string]string
fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc042004028 (*map[string]string)(nil)
mv = new(map[string]string)
fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc000006028 &map[string]string(nil)
*mv = make(map[string]string)
(*mv)["a"] = "a"
fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc042004028 &map[string]string{"a":"a"}
```

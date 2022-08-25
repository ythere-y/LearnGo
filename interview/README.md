# Go 语言学习

> 来自 [博客](https://geektutu.com/post/qa-golang.html) 




## notes

### go build

> 参考[博客](http://c.biancheng.net/view/120.html)

#### 无参数build
源码中没有依赖GOPATH的包引用，那么可以直接无参数`go build`。自动生成可执行文件`gobuild`。

#### build+文件列表
go build + 文件列表。默认使用第一个文件名字作为输出文件，但是得到的结果都是main包的main函数。

使用-o参数可以指定名称如，`go build -o main.go lib.go`。

#### go build+包

`go build -o main chapter11/goinstall`

#### 编译时的附加参数

-v 编译时显示包名
-p n 开启并发编译，默认情况下该值为CPU逻辑核数
-a 强制重新构建
-n 打印编译时会用到的所有命令，但不真正执行
-x 打印编译时会用到的所有命令
-race 开启竞态检测



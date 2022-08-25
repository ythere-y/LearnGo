package interview

import "fmt"

// region what's about init()

/*
init()函数什么时候执行？

在main函数执行之前

详细：init()函数是go初始化的一部分，由runtime初始化每个导入的包，
初始化不是按照从上到下的导入顺序，而是按照解析的依赖关系，没有依赖的包最先初始化。

每个包首先初始化包作用域的常量和变量（常量优先于变量），
然后执行包的init()函数。
同一个包，甚至是同一个源文件可以有多个init()函数。
init()函数没有入参和返回值，不能被其他函数调用，
同一个包内多个init()函数的执行顺序不作保证。

执行顺序：import –> const –> var –>init()–>main()

一个文件可以有多个init()函数！
*/

// endregion

// region object on stack or heap?

/*
Go和C++不同，Go局部变量会进行逃逸分析。
如果变量离开作用域后没有被引用，则优先分配到栈上，否则分配到堆上。

那么如何判断是否发生了逃逸呢？

go build -gcflags '-m -m -l' xxx.go.

关于逃逸的可能情况：
1. 变量大小不确定，
2. 变量类型不确定，
3. 变量分配的内存超过用户栈最大值，
4. 暴露给了外部指针。
*/

// endregion

// region can interface be compared?
/*
Go语言中，interface的内部实现包含了2个字段，类型T和值V，interface 可以用`==`或者`!=`来比较。
2个interface相等有以下2种情况：
1. 两个interface均等于nil（此时V和T都处于unset状态）
2. 类型T相同，且对应的值V相等。
*/

type Stu struct {
	Name string
}
type StuInt interface{}

func interfaceCompare() {
	var stu1, stu2 StuInt = &Stu{"Tom"}, &Stu{"Tom"}
	var stu3, stu4 StuInt = Stu{"Tom"}, Stu{"Tom"}
	// stu1 和 stu2 类型是*Stu,值是Stu结构体的地址，二者地址不同
	fmt.Printf("compare stu1 and stu2 = %v\n", stu1 == stu2)
	// stu1 和 stu3 类型不同
	fmt.Printf("compare stu1 and stu3 = %v\n", stu1 == stu3)
	// stu3 和 stu4 类型相同，各字段也想通
	fmt.Printf("compare stu3 and stu4 = %v\n", stu3 == stu4)
}

// endregion

// region can two nil be Not equal?

/*
可能不等。interface在运行时绑定值，只有值为nil接口值才为nil，但是与指针的nil不相等
而类型不同的nil之间无法直接比较，会报错。

总结：两个nil只有在类型相同时才相等
*/

func nilNotEqual() {
	var p *int = nil
	//var pStr *string = nil
	var pStu *Stu = nil
	var i interface{} = nil
	if p == i {
		fmt.Printf("*int(nil) and interface{}(nil) are equal\n")
	} else {
		fmt.Printf("*int(nil) and interface{}(nil) are NOT equal\n")
	}
	//if p == pStr {
	//	fmt.Printf("*int(nil) and *string(nil) are equal\n")
	//} else {
	//	fmt.Printf("*int(nil) and *string(nil) are NOT equal\n")
	//}

	if i == pStu {
		fmt.Printf("*int(nil) and *string(nil) are equal\n")
	} else {
		fmt.Printf("*int(nil) and *string(nil) are NOT equal\n")
	}

}

// endregion

// region garbage collection
/*
垃圾回收机制是Go一大特色。

Go1.3采用标记清除法，Go1.5采用三色标记法，Go1.8采用三色标记法+混合写屏障

标记清除法：
分为两个阶段：标记 and 清除

标记阶段：从根对象出发寻找并标记所有存活的对象。
清除阶段：遍历堆中所有的对象，回收未标记的对象，并加入空闲链表。

缺点：需要暂停程序STW


三色标记法：
将对象标记为白色、灰色或灰色。

白色：不确定对象（默认色）；黑色：存活对象； 灰色：存活对象，但是子对象待处理。

标记开始时，现将所有对象加入白色集合（需要STW）。首先将根对象标记为灰色，
然后将一个对象从灰色集合取出，遍历其子对象，放入灰色集合。
同时将取出的对象放入黑色集合，直到灰色集合为空。
最后的白色集合对象就是需要清理的对象。

缺陷：如果对象的引用被用户修改了，那么之前的标记就无效了。因此Go采用了写屏障技术，
当对象新增或者更新都会将其着色为灰色。

一次完整的GC分为四个阶段：
1. 准备标记（需要STW），开启写屏障
2. 开始标记
3. 标记结束（STW），关闭写屏障
4. 清理（并发）

基于插入写屏障和删除写屏障在结束时需要STW来重新扫描栈，带来性能瓶颈。混合写屏障分为以下四步：

1. GC开始时，将栈上的全部对象标记为黑色（不需要二次扫描，无需STW）
2. GC期间，任何栈上创建的新对象均为黑色
3. 被删除引用的对象标记为灰色
4. 被添加引用的对象标记为灰色

总而言之就是确保黑色对象不能引用白色对象，这个改进直接使得GC时间从2s降低到2us。
*/
// endregion

// region safe to return local variable in function?
/*
这一点和C++不同，在Go里面返回局部变量的指针是安全的。
因为Go会进行逃逸分析，如果发现局部变量的作用域超过该函数则会把指针分配到堆区，避免内存泄漏。
*/
// endregion

// region if any type T() can call method of *T?

/*
一个T类型的值可以调用*T类型申明的方法，当且仅当T是可寻址的。（调用的时候加上&）
反之，*T可以调用T()的方法，因为指针可以解引用。（调用的时候加上*）
*/
// endregion

// region how slice expend?
/*
如果当前容量小于1024，则判断所需容量是否大于原来容量2倍，如果大于，当前容量加上所需容量；否则当前容量乘2。

如果当前容量大于1024，则每次按照1.25倍速度递增容量，也就是每次加上cap/4。
*/

func sliceExpand() {
	var a []int
	var lastCap = cap(a)
	fmt.Printf("append by step = 1 >>> \n")
	for i := 0; i < 2000; i++ {
		a = append(a, i)
		if cap(a) != lastCap {
			fmt.Printf("cap change : [%v->%v],expand ratio = %v\n", lastCap, cap(a), float64(cap(a))/float64(lastCap))
		}
		lastCap = cap(a)
	}

	var b []int
	var bSize int = 500
	fmt.Printf("append by step = %v >>> \n", bSize)
	for i := 0; i < bSize; i++ {
		b = append(b, i)
	}
	a = []int{}
	lastCap = cap(a)
	for i := 0; i < 20; i++ {
		a = append(a, b...)
		if cap(a) != lastCap {
			fmt.Printf("cap change : [%v->%v],expand ratio = %v\n", lastCap, cap(a), float64(cap(a))/float64(lastCap))
		}
		lastCap = cap(a)
		//fmt.Printf("a.cap = %v\n", cap(a))
	}

}

// endregion

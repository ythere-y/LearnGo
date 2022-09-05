# Golang的sync.Map

sync.Map(在并发环境中使用的map)

## 特性
1. 无需初始化，声明之后即可使用，且不能使用make创建。
2. 不能用[]的方式直接取值、设置，必须的调用方法。Store是写，Load是读，Delete是删除。key和value以interface{}类型保存。
3. 使用Range配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range参数中回调函数的返回值在需要继续迭代遍历时，返回true,终止迭代遍历时，返回false.
```go

scene.Range(func(k, v interface{}) bool {

fmt.Println("iterate:", k, v)
return true
})
```

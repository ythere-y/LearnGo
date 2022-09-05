

# Golang 断言 type assertion
> [run golang online](https://go.dev/tour/methods/15)
>
> [聊聊golang的断言](https://segmentfault.com/a/1190000038318158)
> 
> [golang断言实现原理](https://segmentfault.com/a/1190000039894161)
>
> [learn in English](https://betterprogramming.pub/golang-type-assertion-d5517d81c366)

原理总结：
- 空接口类型断言实现流程：空接口类型断言实质是将eface中_type与要匹配的类型进行对比，匹配成功在内存中组装返回值，匹配失败直接清空寄存器，返回默认值。
- 非空接口类型断言的实质是 iface 中 *itab 的对比。*itab 匹配成功会在内存中组装返回值。匹配失败直接清空寄存器，返回默认值
- 泛型是在编译期做的事情，使用类型断言会消耗一点性能，类型断言使用方式不同，带来的性能损耗也不同，具体请看上面的章节。

性能总结：
- 空接口类型的类型断言代价并不高，与直接类型转换几乎没有性能差异
- 空接口类型使用type switch进行类型断言时，随着case的增多性能会直线下降
- 非空接口类型进行类型断言时，随着接口中方法的增多，性能会直线下降
- 直接进行方法调用要比非接口类型进行类型断言要高效很多
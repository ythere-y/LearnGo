# call visualization

[知乎专栏](https://zhuanlan.zhihu.com/p/463849627)

按照教程安装后使用的时候，发现并不成功

## 初运行尝试
运行的命令是
```bash
go-callvis .
```

在测试项目中
报错如下
```log
panic: T

goroutine 743 [running]:
golang.org/x/tools/go/types/typeutil.Hasher.hashFor({0x46831a0?}, {0x47e7880?, 0xc0003baea0?})
        /Users/bytedance/go/pkg/mod/golang.org/x/tools@v0.0.0-20200305224536-de023d59a5d1/go/types/typeutil/map.go:302 +0x426
golang.org/x/tools/go/types/typeutil.Hasher.Hash({0x0?}, {0x47e7880, 0xc0003baea0})
        /Users/bytedance/go/pkg/mod/golang.org/x/tools@v0.0.0-20200305224536-de023d59a5d1/go/types/typeutil/map.go:226 +0x65
golang.org/x/tools/go/types/typeutil.Hasher.hashFor({0x46831a0?}, {0x47e77b8?, 0xc00076cf90?})
        /Users/bytedance/go/pkg/mod/golang.org/x/tools@v0.0.0-20200305224536-de023d59a5d1/go/types/typeutil/map.go:269 +0x106
golang.org/x/tools/go/types/typeutil.Hasher.Hash({0xc008abe2a0?}, {0x47e77b8, 0xc00076cf90})
        /Users/bytedance/go/pkg/mod/golang.org/x/tools@v0.0.0-20200305224536-de023d59a5d1/go/types/typeutil/map.go:226 +0x65
golang.org/x/tools/go/types/typeutil.Hasher.hashTuple({0x0?}, 0xc000611308)
        /Users/bytedance/go/pkg/mod/golang.org/x/tools@v0.0.0-20200305224536-de023d59a5d1/go/types/typeutil/map.go:310 +0x5f
golang.org/x/tools/go/types/typeutil.Hasher.hashFor({0x46831a0?}, {0x47e7858?, 0xc000611308?})
        /Users/bytedance/go/pkg/mod/golang.org/x/tools@v0.0.0-20200305224536-de023d59a5d1/go/types/typeutil/map.go:300 +0x391
golang.org/x/tools/go/types/typeutil.Hasher.Hash({0xc0077dfa00?}, {0x47e7858, 0xc000611308})
        /Users/bytedance/go/pkg/mod/golang.org/x/tools@v0.0.0-20200305224536-de023d59a5d1/go/types/typeutil/map.go:226 +0x65
golang.org/x/tools/go/types/typeutil.(*Map).At(0xc005a49ef8, {0x47e7858, 0xc000611308})
        /Users/bytedance/go/pkg/mod/golang.org/x/tools@v0.0.0-20200305224536-de023d59a5d1/go/types/typeutil/map.go:88 +0x48
golang.org/x/tools/go/ssa.(*Program).needMethods(0xc005a49ea0, {0x47e7858?, 0xc000611308?}, 0x0)
        /Users/bytedance/go/pkg/mod/golang.org/x/tools@v0.0.0-20200305224536-de023d59a5d1/go/ssa/methods.go:156 +0x66
golang.org/x/tools/go/ssa.(*Program).needMethods(0xc005a49ea0, {0x47e77b8?, 0xc008ab38f0?}, 0x0)
        /Users/bytedance/go/pkg/mod/golang.org/x/tools@v0.0.0-20200305224536-de023d59a5d1/go/ssa/methods.go:181 +0x1b4
golang.org/x/tools/go/ssa.(*Program).needMethods(0xc005a49ea0, {0x47e7790?, 0xc0003d61c0?}, 0x0)
        /Users/bytedance/go/pkg/mod/golang.org/x/tools@v0.0.0-20200305224536-de023d59a5
```

在业务项目中运行信息中有如下提示
```log
panic("Unsupported Go version. Supported versions are: 1.15, 1.16, 1.17, 1.18") (no value) used as value
```

于是去解决go版本问题，


直接在goland里面切换go版本，然后关闭terminal重新打开，运行`go version`显示新的版本。在我的尝试下，使用1.16版本是可以成功运行的。
针对业务项目运行，需要的时间非常长。
# 常量

```go
const pi = 3.1415926

// 批量声明常量，如果下面一行没有值，那么默认跟上面一行是一样的
const {
    status = 200
    notFound = 404
    n2
    n3
}
```

# iota
iota在关键字const出现时将被重置为0，const中每新增一行iota计数一次

```go
const {
    a1 = iota // 0
    a2 // 1
    a3 // 2
}

const {
    b1 = iota //0
    b2 // 1
    _ // 2
    b3 //3
}

const {
    b1 = iota //0
    b2 = 100 //iota = 1
    b3 = iota //2
}

// 定义数量级
const {
    _ = iota //0
    KB = 1 << (10 * iota)
    MB = 1 << (10 * iota)
}
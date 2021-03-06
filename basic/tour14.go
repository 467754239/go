package main

import (
    "fmt"
    "math/cmplx"
)

// 变量组
var (
    ToBe bool = false
    MaxInt uint64 = 1 << 64 - 1
    z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
    // fmt.PrintF("%T") 可以打印变量的类型
    const f = "%T(%v)\n"
    fmt.Printf(f, ToBe, ToBe)
    fmt.Printf(f, MaxInt, MaxInt)
    fmt.Printf(f, z, z)
}

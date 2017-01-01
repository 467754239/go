package main

import (
    "fmt"
)

func main() {
    // 声明一个数组，数组中有2个字符串类型要===元素
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1])
    fmt.Println(a)
}

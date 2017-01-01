package main

import (
    "fmt"
)

// 函数add 接收x, y两个参数 分别是int类型 返回一个参数 int类型
func add(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(2, 3))
}

package main

import (
    "fmt"
)

// 结构体
type Vertex struct {
    X int
    Y int
}

func main() {
    p := Vertex{1, 2}
    // & 表示该变量的指针地址
    q := &p
    // 结构体元素重新赋值
    q.X = 1e9
    fmt.Println(p)
}

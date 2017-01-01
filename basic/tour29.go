package main

import (
    "fmt"
)

// 结构体
type Vertex struct {
    // int 零值是0
    X, Y int
}

var (
    p = Vertex{1, 2}    // has type Vertex
    q = &Vertex{1, 2}   // has type *Vertex
    r = Vertex{X: 1}    // Y: 0 is implicit
    s = Vertex{}        // X: 0 and Y: 0
)

func main() {
    fmt.Println(p, q, r, s)
}


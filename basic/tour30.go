package main

import "fmt"

type Vertex struct {
    X, Y int
}

func main() {
    v := new(Vertex)    // new 指定内存地址
    fmt.Println(v)
    v.X, v.Y = 11, 2
    fmt.Println(v)
}

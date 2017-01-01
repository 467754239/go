package main

import (
    "fmt"
)

func swap(x, y string) (string, string) {
    return x, y
}

func swap_alais(x, y string) (a, b string) {
    /*
	由于返回值a, b string 局部变量已经存在，所以无需
	a := x
	b := y
     */
    a = x
    b = y
    return
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b)

    c, d := swap_alais("vip", "kid")
    fmt.Println(c, d)
}

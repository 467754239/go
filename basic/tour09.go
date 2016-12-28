package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return x, y
}

func swap_alais(x, y string) (x, y string) {
	return x, y
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	c, d := swap_alais("vip", "kid")
	fmt.Println(c, d)
}

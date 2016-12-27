package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 包rand实现伪随机生成器
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println("My favorite number is", rand.Seed(42))
}

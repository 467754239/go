package main

import (
	"fmt"
)

/*
func main() {
	// 方式1
	a := 10

	if a > 0 {
		a += 100
	} else if a == 0 {
		a = 0
	} else {
		a -= 100
	}

	fmt.Println(a)

	// 方式2
	b := 101
	if b > 0 {
		b += 100
	} else {
		b -= 100
	}
	fmt.Println(b)

	// 方式3
	//支持一个初始化参数，变量c的作用域尽在if/else范围内.
	if c := 1; c < 10 {
		c += 1
		fmt.Println("c =", c)
	}

	// 方式4
	if a, b := 1, 2; a+b == 10 {
		fmt.Println("into if")
	} else {
		fmt.Println("into else")
	}

    // 方式5
    if x := computedValue(); x > 10 {
        ...
    } else {
        ...
    }

}

func main() {

	// 第一种形式
	a := 1
	for {
		a++
		if a > 3 {
			break
		}
		fmt.Println(a)
	}
	fmt.Println("a finish")

	// 第二种形式
	b := 1
	for b < 3 {
		b++
		fmt.Println(b)
	}
	fmt.Println("b finish")

	// 第三种形式
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("i finish")
}



func main() {

	// 第一种形式
	a := 1
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}

	// 第二种形式
	ab := 1
	switch {
	case ab >= 0:
		fmt.Println("ab=0")
		fallthrough // 不终止，继续检查下一个case.
	case ab >= 1:
		fmt.Println("ab=1")
	default:
		fmt.Println("None")
	}

	// 第三种形式
	switch ab := 1; {
	case ab >= 0:
		fmt.Println("ab=0")
		fallthrough // 不终止，继续检查下一个case.
	case ab >= 1:
		fmt.Println("ab=1")
	default:
		fmt.Println("None")
	}

}
*/

func main() {
LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				fmt.Println(i)
				break LABEL1
			}
		}
	}
	fmt.Println("ok")

	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				fmt.Println(i)
				goto LABEL2
			}
		}
	}
LABEL2:
	fmt.Println("ok")

}

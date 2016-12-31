package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	/*
		func Sqrt(x float64) float64
	    返回x的平方根
	    25的平方根是5
	*/
	var f float64 = math.Sqrt(float64(3*3 + 4*4))
	var z int = int(f)
	fmt.Println(x, y, z)
}

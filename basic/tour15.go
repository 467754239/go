package main

import (
    "fmt"
    "math"
)

func Sqrtd(x float64) float64 {
    // 初始点
    z := 1.0

    // 重复计算10次
    for i := 0; i < 10; i++ {
	z = z - (z * z - x) / ( 2 * z)
	//fmt.Println(z)
    }
    return z
}

func main() {
    var x, y int = 3, 4

    var f float64 = math.Sqrt(float64(3 * 3 + 4 * 4))
    var z int = int(f)
    fmt.Println(x, y, z)

    // 用牛顿法实现开方函数
    fmt.Println(Sqrtd(2))

    // build-in sqrt
    fmt.Println(math.Sqrt(2))
}

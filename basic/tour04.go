package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // 这个程序的运行环境是固定的，因此 rand.Intn 总是会返回相同的数字。
    fmt.Println("My favoritenumber is", rand.Intn(10))

    // 为了得到不同的数字，需要生成不同的种子数。rand.Seed
    // 根据时间设置随机数种子，这时重复执行会得到不同的随机数。
    rand.Seed(int64(time.Now().Nanosecond()))
    fmt.Println(rand.Intn(10))

    // 获取指定范围内的随机数
    for i := 0; i < 10; i++ {
	//fmt.Printf("%v", rand.Intn(10))
        fmt.Println(rand.Intn(10))
    }

}

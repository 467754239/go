package main

import (
    "fmt"
)

func main() {
    /*python
    sum = 1
    while sum < 1000:
	    sum += sum
    print sum
     */

    sum := 1
    for sum < 1000 {
	fmt.Println(sum)
	sum += sum
    }
    fmt.Println(sum)
}

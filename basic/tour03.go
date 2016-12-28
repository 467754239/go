package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	fmt.Println("And if you try to open a file:")
	fmt.Println(os.Open("C:/Windows/SynInst.log"))

	fmt.Println("Or access the network:")
	fmt.Println(net.Dial("tcp", "54.223.222.180:80"))
}

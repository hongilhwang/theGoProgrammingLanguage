package main

import (
	"fmt"
)

func main() {
	var x uint64 = 84
	fmt.Println(byte(x >> (0 * 8)))
	fmt.Println(byte(x >> (1 * 8)))
	fmt.Printf("%v", 1/2)
}

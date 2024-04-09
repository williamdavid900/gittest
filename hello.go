package main

import (
	"fmt"
	"math/bits"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello There Janice Trenker"))
	fmt.Println("try git, again.", bits.Reverse64(123))
}

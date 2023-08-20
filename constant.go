package main

import (
	"fmt"
)

func main() {
	const b = 10
	fmt.Println(b)
	b = 30 // Error, since b is a constant and we can not change its value
	fmt.Println(b)
}

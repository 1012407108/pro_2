package main

import (
	"fmt"
)

func main() {
	fv := func() {
		fmt.Println("HelloWorld!!")
	}
	fmt.Println(fv())
}

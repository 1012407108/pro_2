package main

import (
	"fmt"
)

func leiJia(a int) (sum int) {
	if a == 1 {
		return 1
	}
	sum = leiJia(a-1) + a
	return
}

func main() {
	var sum int
	sum = leiJia(100)
	fmt.Println(sum)
}

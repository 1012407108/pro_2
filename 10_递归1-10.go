package main

import "fmt"

func main() {
	for i := 10; i > 0; i-- {
		fmt.Println(digui(i))
	}

}
func digui(n int) (a int) {
	if n == 1 {
		a = n
	} else {
		a = digui(n-1) + 1

	}
	return

}

package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
	)
	fmt.Printf("%d%d%d", a, b, c)
}

package main

import "fmt"

func main() {
	n := 0
	Multiply(6, 5, &n)
	fmt.Println(n)
}
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

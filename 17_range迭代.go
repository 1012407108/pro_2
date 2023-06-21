package main

import (
	"fmt"
)

func main() {
	var s string
	s = "abcdefghijklmnopqrstuvwxyz"
	for key, password := range s {
		fmt.Printf("%d,%c,\t", key, password)
	}
	for key := range s {
		fmt.Printf("%d,\t", key)
	}
	for _, password := range s {
		fmt.Printf("%c,\t", password)
	}
}

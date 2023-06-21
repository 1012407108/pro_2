package main

import "fmt"

func main() {
	var result int64 = 0
	result = jieCeng(30)
	fmt.Println(result)
}
func jieCeng(n int64) (result int64) {
	if n == 1 {
		result = 1
	} else {
		result = n * jieCeng(n-1)
	}
	return
}

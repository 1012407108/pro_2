package main

import "fmt"

func test() (a, b, c int) {
	return 1, 2, 3
}

func main() {
	a := 10
	fmt.Printf("%d\n", a)
	b, c := 20, 30
	fmt.Println("a= ", a, "b=", b, "c=", c)
	fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c)
	var h, j, k int
	h, j, k = test()
	fmt.Printf("h = %d,j=%d,k=%d\n", h, j, k)
	_, j, k = test()
	fmt.Printf("j=%d,k=%d\n", j, k)
	j, k = k, j
	fmt.Print(j, k)

}

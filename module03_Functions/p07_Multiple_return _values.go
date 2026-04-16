package main

import "fmt"

func calc(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	sum, product := calc(3, 4)
	fmt.Println(sum, product)
}

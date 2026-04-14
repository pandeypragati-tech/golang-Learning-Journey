package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Println("Before Swap:", a, b)

	a, b = b, a

	fmt.Println("After Swap:", a, b)
}

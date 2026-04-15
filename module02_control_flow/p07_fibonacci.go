package main

import "fmt"

func main() {
	var n int
	a := 0
	b := 1

	fmt.Println("Enter number of terms :")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		fmt.Println(a, " ")
		a, b = b, a+b
	}
}

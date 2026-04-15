package main

import "fmt"

func main() {

	var n int
	fact := 1
	fmt.Println("Enter n:")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fact *= i
	}
	fmt.Println("Factorial:", fact)

}

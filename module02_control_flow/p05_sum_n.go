package main

import "fmt"

func main() {
	var n, sum int

	fmt.Println("Enter n:")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)
}

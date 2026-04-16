package main

import "fmt"

func fibonacci(n int) {
	a, b := 0, 1
	for i := 1; i <= n; i++ {
		fmt.Println(a, " ")
		a, b = b, a+b
	}
}
func main() {
	fibonacci(7)
}

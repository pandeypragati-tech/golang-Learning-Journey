package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Enter Three Numbers:")
	fmt.Scanln(&a, &b, &c)

	if a >= b && a >= c {
		fmt.Println("The Largest Number is :", a)
	} else if b >= a && b >= c {
		fmt.Println("The Largest Number is:", b)
	} else {
		fmt.Println("The Largest Nujmber is :", c)
	}
}

package main

import "fmt"

func main() {
	var p, r, t float64

	fmt.Print("Enter Principal: ")
	fmt.Scanln(&p)
	fmt.Print("Enter Rate: ")
	fmt.Scanln(&r)
	fmt.Print("Enter Time: ")
	fmt.Scanln(&t)

	si := (p * r * t) / 100

	fmt.Println("Simple Interest:", si)
}

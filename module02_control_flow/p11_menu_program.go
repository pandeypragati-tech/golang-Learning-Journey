package main

import "fmt"

func main() {
	var choice int
	var a, b int

	for {
		fmt.Println("1.Add 2.Subtract 3.Exit")
		fmt.Scan(&choice)

		if choice == 3 {
			break
		}

		fmt.Scan(&a, &b)

		if choice == 1 {
			fmt.Println(a + b)
		} else if choice == 2 {
			fmt.Println(a - b)
		}
	}
}

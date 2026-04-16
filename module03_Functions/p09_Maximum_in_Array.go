package main

import "fmt"

func max(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	arr := []int{3, 7, 2, 9}
	fmt.Println(max(arr))
}

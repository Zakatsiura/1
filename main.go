package main

import "fmt"

func main() {
	fmt.Println(findMin(2, 4, 5, -7, -3))
}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]

	for _, i := range numbers {
		if i < min {
			min = i
		}
	}
	return min
}

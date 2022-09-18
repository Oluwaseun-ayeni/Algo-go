package main

import "fmt"

func main() {

	var size int
	fmt.Println("Enter the size of the array")
	fmt.Scan(&size)

	arr := make([]int, size)
	fmt.Println("Enter the number you to add up:")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])

	}
	total := 0
	for _, a := range arr {
		total += a
	}
	fmt.Println("Here is your output", total)
}

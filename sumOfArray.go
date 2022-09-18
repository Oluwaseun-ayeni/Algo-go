package main

import "fmt"

func main() {

	total := 0
	arr := []int{2, 7, 9, 7, 7}

	for i := 0; i < len(arr); i++ {
		total += arr[i]

	}
	fmt.Println(total)
}

package main

import "fmt"

func sum(x int, y int) {

	total := 10
	total = x + y
	fmt.Println(total)

}

func main() {
	sum(20, 30)

}

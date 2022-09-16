package main

import "fmt"

func main() {
	// to create a map,you need to use the builtin function make
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map", m)

	r1 := m["k2"]
	r1 = 16
	fmt.Println("r1:", r1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, boo := m["k2"]
	fmt.Println("prs:", boo)

	n := map[string]int{"f1": 45, "f2": 64}
	fmt.Println("map2:", n)

}

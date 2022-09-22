package main

import "fmt"

type youtubers struct {
	speed int
	ross  int
	ksi   int
	kai   int
}

func (a youtubers) rating() {
	fmt.Println("speed rating is:", a.speed)
	fmt.Println("ross rating is:", a.ross)
	fmt.Println("ksi rating is :", a.ksi)
	fmt.Println("kai rating is :", a.kai)

}

func main() {
	res := youtubers{
		speed: 8,
		ross:  6,
		ksi:   7,
		kai:   9,
	}

	res.rating()

}

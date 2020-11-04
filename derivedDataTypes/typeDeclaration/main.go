package main

import "fmt"

type mile int
type kilometer int
type library []string

func distanceToEnemy(distance mile) {
	fmt.Println("Distance to enemy ")
	fmt.Println(distance, "miles")
}

func printBooks(lib library) {
	for _,value := range lib {
		fmt.Println(value)
	}
}

func main() {
	var distance mile = 5
	//fmt.Println(distance)
	//distance++
	//fmt.Println("distance ++")
	//fmt.Println(distance)
	distanceToEnemy(distance)
	//var distance2 kilometer = 8
	//distanceToEnemy(distance2) // type kilometer as type mile
	var lib library = library{"Book1", "Book2", "Book3"}

	printBooks(lib)
}

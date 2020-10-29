package main

import "fmt"

// var variable_name [size]variable_type

//var numbers [3]int = [3]int{1, 2, 3}
var numbers = [3]int{1, 2, 3} // short option
var symbols = [...]string{"a", "b", "c", "d"}

func main() {
	words := [3]string{"apple", "axe", "ant"}
	fmt.Println(numbers[0],numbers[1],numbers[2])
	fmt.Println(words)
	fmt.Println(symbols)
}

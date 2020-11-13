package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "hello cold")

	//Println(), Print() and  Printf()
	//same as Fprintln() Fprint() and Fprintf() with first parameter os.Stdout

	tom := person {
		name:"Tom",
		age: 24,
		weight: 68.5,
	}
	fmt.Printf("%-10s %-10d %-10.3f\n",
		tom.name, tom.age, tom.weight)
	fmt.Print("Hello ")
	fmt.Println("cold!")
}



type person struct {
	name string
	age int32
	weight float64
}
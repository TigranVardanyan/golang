package main

import "fmt"

type person struct {
	name string
	age int
}

func (p *person) updateAge(newAge int)  {
	(*p).age = newAge
}

func main() {
	var tigran = person{name: "Tigran", age: 23}
	var tigranPointer = &tigran
	fmt.Println(tigranPointer.age)
	tigranPointer.updateAge(24)
	fmt.Println(tigranPointer.age)
	fmt.Println(tigran.age)
	tigran.updateAge(25)
	fmt.Println(tigran.age)
}

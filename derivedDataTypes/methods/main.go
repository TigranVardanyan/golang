package main

import "fmt"

//func (param_name reciver_type) method_name (params) (return_values_type){
//	action
//}

type library []string

func (l library) print()  {
	for _,val := range l {
		fmt.Println(val)
	}
}

type person struct {
	name string
	age int
}

func (p person) printPerson()  {
	fmt.Println(p.name, p.age)
}

func (p person) printPersonEat(food string)  {
	fmt.Println(p.name, "eat " + food)
}

func main() {
	var lib library = library{ "Book1", "Book2", "Book3" }
	lib.print()
	var Tigran person = person{name: "Tigran", age: 24}
	Tigran.printPerson()
	Tigran.printPersonEat("apple")
}

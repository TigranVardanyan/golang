package main

import "fmt"

//type struct_name struct{
//	struct_field
//}

type person struct {
	name string
	age int
	height float64
}

func main() {
	var tom person = person{"Tom", 20, 170}
	var alice person = person{"Alice", 19, 166}
	fmt.Println(tom.name,alice.name)
	var tomPointer *person = &tom
	tomPointer.age = 24
	fmt.Println(tom.age)


}

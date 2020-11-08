package main

import "fmt"

//add struct as struct field
type contact1 struct {
	email string
	phone string
}

type person1 struct {
	name string
	age int
	contactInfo contact1
}



//immediately add fields to the structure

type contact2 struct {
	email string
	phone string
}

type person2 struct {
	name string
	age int
	contact2
}

type node struct {
	value int
	next *node
}


func main() {
	var tom1 = person1{
		"Tom",
		24,
		contact1{"tom@gmail.com",
			"+123456789"}}

	//tom.contactInfo.email = "supporttom@gmail.com"

	fmt.Println(tom1.contactInfo.email)      // supertom@gmail.com
	fmt.Println(tom1.contactInfo.phone)      // +1234567899


	var tom2 = person2{
		name: "Tom",
		age: 24,
		contact2 : contact2{"tom@gmail.com",
			"+123456789"}}

	tom2.email = "supporttom@gmail.com"

	fmt.Println(tom2.email)      // supertom@gmail.com
	fmt.Println(tom2.phone)      // +1234567899



	////////

	first := node{value: 4}
	second := node{value: 5}
	third := node{value: 6}

	first.next = &second
	second.next = &third

	var current *node = &first
	for current != nil{
		fmt.Println(current.value)
		current = current.next
	}
}
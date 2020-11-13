package main

import "fmt"

func main() {
	var name string
	var age string

	//fmt.Print("Enter your name: ")
	//fmt.Fscan(os.Stdin, &name)
	//fmt.Print("Enter your age: ")
	//fmt.Fscan(os.Stdin, &age)


	//with fmt

	fmt.Print("Введите имя: ")
	fmt.Scan(&name) // there can be multi data scan
	//this function reads the value of a string and other data types up to a space
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)


	fmt.Println(name, age)

}

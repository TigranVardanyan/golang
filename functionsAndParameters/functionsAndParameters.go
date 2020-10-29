package main

import "fmt"

func main() {
	fmt.Println("Hello Go")
	sayHello()
	sayHelloWithName("Tigran")
	sayHelloWithName("Vahag")

	add(10, 15)

	addMoreThanTwo(1,2,3,4,5,6,7,8,9)

	addMoreThanTwo([]int{1, 2, 3}...)
}

/*
func function_name (parameters) (return type){
    action
}
 */

func sayHello() {
	fmt.Println("Hello man")
}

func sayHelloWithName(name string) {
	fmt.Println("Hello " + name)
}

func add(a int, b int) {
	fmt.Println(a + b)
}

func addMoreThanTwo(numbers ...int)  {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Print("SUM = ")
	fmt.Println(sum)
}
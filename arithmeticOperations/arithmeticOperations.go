package main

import "fmt"

func main() {
	var a = 10
	var b = 5
	additions(a, b)
	subtract(a, b)
	multi(a, b)
	division(a, b)
}

func additions(a int, b int)  {
	var c = a + b
	fmt.Println(c)
}

func subtract(a int, b int)  {
	var c = a - b
	fmt.Println(c)
}

func multi(a int, b int)  {
	var c = a * b
	fmt.Println(c)
}

func division(a int, b int)  {
	var c = a / b
	fmt.Println(c)
}
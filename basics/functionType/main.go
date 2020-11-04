package main

import "fmt"

func main() {
	var f func(int, int) int = add
	fmt.Println(f(3, 4))

	var x = f(4, 5)
	fmt.Println(x)

	display("Hello, world")

	mainAdd := add

	fmt.Println(mainAdd(1,2))

	action(5,10,add)
	action(5,10,multiply)
	f1 := selectFn(1)
	fmt.Println(f1(4,6))
	f2 := selectFn(2)
	fmt.Println(f2(4,6))
	f3 := selectFn(3)
	fmt.Println(f3(4,6))
}

//type func(int,int) int
func add(x int, y int) int{
	return x + y
}
//type func(int,int) int
func multiply(x int, y int) int{
	return x * y
}
func subtract(x int, y int) int {
	return x - y
}
func action(n1 int, n2 int, operation func(int, int) int) {
	result := operation(n1, n2)
	fmt.Println(result)
}

//type func(string,string) string
func concatStr(x string, y string) string {
	return x + " " + y
}
//type func(string)
func display(message string) {
	fmt.Println(message)
}

func selectFn(n int) func(int, int) int{
	if n==1 {
		return add
	}else if n==2{
		return subtract
	}else{
		return multiply
	}
}
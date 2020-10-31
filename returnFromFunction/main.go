package main

import "fmt"

//func function_name (parameters) return_type {
//	action
//	return return_value
//}

func main() {
	var a = 1
	var b = 2
	var str1 = "qwerty"
	var str2 = "asdfgh"
	var c = add(a, b)
	fmt.Println(c)
	var d = multi(a, b)
	fmt.Println(d)
	var e, f = multiReturn(a, b, str1, str2)
	fmt.Println(e)
	fmt.Println(f)
}

func add(x, y int) int {
	return x + y
}

//immediately return
func multi(x, y int) (z int) {
	z = x * y
	return
}

func multiReturn(x, y int, firstName, lastName string) (int, string) {
	var z int = x + y
	var fullName = firstName + " " + lastName
	return z, fullName
}

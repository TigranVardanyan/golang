package main

import "fmt"

func main() {
	var a int = 8
	var b int = 3
	var c bool = a == b
	fmt.Println(c)      // false

	var i bool = 4 > 5 && 6 > 8       // false
	var j bool = 3 <= 5 && 10 > 8     // true
	fmt.Println(i)
	fmt.Println(j)
}

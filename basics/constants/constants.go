package main

import "fmt"

//const pi float64 = 3.14

//pi = 15 //error

const (
	a = 1
	b
	c
	d = 3
	f
)

var first = 15
//const second = first //error! you cannot set const value to other variables

func main() {
	fmt.Println(a,b,c,d,f)
	fmt.Println(first)
}

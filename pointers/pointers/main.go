package main

import "fmt"

func main() {
	var i int = 15
	var p *int
	p = &i
	fmt.Println("Address", p)
	fmt.Println("Value", *p)
	*p = 25
	fmt.Println("*p changed")
	fmt.Println("i = ",i)

	//var pf *float64
	//fmt.Println("Value:", *pf) //invalid memory address or nil

	var num = new(int)
	fmt.Println("new(int) def value",*num)
	*num = 50
	fmt.Println("changed value", *num)
}

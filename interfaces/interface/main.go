package main

import "fmt"

//type interface_name interface{
//	function
//}

type Vehicle interface {
	move()
}

//var v vehicle = vehicle{} can not do this

type Car struct {}
type Aircraft struct {}

func (c Car) move() {
	fmt.Println("The car is driving")
}
func (c Aircraft) move()  {
	fmt.Println("The aircraft is flying")
}

func main() {
	var tesla Vehicle = Car{}
	var boing Vehicle = Aircraft{}
	tesla.move()
	boing.move()
}
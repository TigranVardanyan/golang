package main

import "fmt"

type Vehicle interface {
	move()
}

type Car struct {
	model string
}

type Aircraft struct {
	model string
}

func (c Car) move()  {
	fmt.Println(c.model, "move")
}

func (a Aircraft) move()  {
	fmt.Println(a.model, "fly")
}

func main() {
	tesla := Car{"Tesla"}
	boeing := Aircraft{"Boing"}
	volvo := Car{"Volvo"}

	vehicles := [...]Vehicle{tesla, boeing, volvo}

	for _, vehicle := range vehicles{
		vehicle.move()
	}
}
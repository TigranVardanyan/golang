package main
import "fmt"

func main() {

	a := 6
	b := 7
	if a < b {
		fmt.Println("a is less than b")
	}

	c := 10
	d := 15
	if c > d {
		fmt.Println("if block")
	} else {
		fmt.Println("else block")
	}
	
	
	key := "Tigran"
	switch key {
	case "Vahag":
		fmt.Println("hello Vahag")
	case "Alberto":
		fmt.Println("hello Alberto")
	case "Tigran":
		fmt.Println("hello Tigran")
	default:
		fmt.Println("no matches")
	}
}
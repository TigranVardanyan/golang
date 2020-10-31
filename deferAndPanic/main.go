package main
import "fmt"

func main() {
	defer finish()
	fmt.Println("Program has been started")
	fmt.Println("Program is working")

	print1()
	print2()
	print3()

	defer print3() // first defer is last
	defer print2()
	defer print1() // last defer is first

	divide(5, 0)
}

func finish(){
	fmt.Println("Program has been finished")
}

func print1()  {
	fmt.Println("1")
}

func print2()  {
	fmt.Println("2")
}

func print3()  {
	fmt.Println("3")
}

func divide(x, y float64) float64{
	if y == 0{
		panic("Division by zero!")
	}
	return x / y
}
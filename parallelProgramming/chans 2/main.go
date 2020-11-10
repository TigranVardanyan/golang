package main

import "fmt"

func main() {

	intCh := make(chan int, 2)

	go factorial(6, intCh)
	fmt.Println(<- intCh)
	fmt.Println("The end")
}

func factorial(n int, ch chan<- int) {
	result := 1
	for i := 1; i <= n; i++{
		result *= i
	}
	ch <- result
	fmt.Println("New Start")
	fmt.Println(<-createChan(5))
	fmt.Println("End")
}

//chan as function return

func createChan(n int) chan int{
	ch := make(chan int) 	//make chan
	go func() {ch <- n}()   // send data to chan
	return ch 				//return chan
}
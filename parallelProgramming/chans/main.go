package main

import "fmt"

//chan element_type

func main() {
	//anonymous function

	//var intCh chan int = make(chan int)
	////short version
	////strCh := make(chan string)
	//
	//go func() {
	//	fmt.Println("Go routine starts")
	//	intCh <- 5
	//}()
	//fmt.Println(<-intCh)
	//fmt.Println("The end")

	intCh := make(chan int)

	go factorial(6, intCh)

	fmt.Println(<-intCh)
	fmt.Println("The End")

	intCh2 := make(chan int, 3)
	intCh2 <- 10
	intCh2 <- 3
	intCh2 <- 24
	fmt.Println(<-intCh2)     // 10
	fmt.Println(<-intCh2)     // 3
	fmt.Println(<-intCh2)
}

func factorial(n int, ch chan int)  {
	result := 1

	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println(n,"-",result)

	ch <- result
}
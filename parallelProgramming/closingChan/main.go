package main

import "fmt"

func main() {
	intCh := make(chan int, 3)
	intCh <- 3
	intCh <- 6
	close(intCh)
	//intCh <- 9  // error, chan is closed

	fmt.Println(<- intCh)
	fmt.Println(<- intCh)
	fmt.Println(<- intCh)


	strCh := make(chan string, 3)
	strCh <- "Hello"
	strCh <- "World"
	close(strCh)    // канал закрыт

	for i := 0; i < cap(strCh); i++ {
		if val, opened := <-strCh; opened {
			fmt.Println(val)
		} else {
			fmt.Println("Channel closed!")
		}
	}
}

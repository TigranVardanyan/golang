package main

import "fmt"

func factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * factorial(n - 1)
}

func fib(n uint) uint  {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n - 1) + fib(n - 2)

}

func main()  {

	fmt.Println(factorial(4))
	fmt.Println(factorial(5))
	fmt.Println(factorial(6))
	fmt.Println()
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
	fmt.Println(fib(6))
	fmt.Println(fib(7))
}

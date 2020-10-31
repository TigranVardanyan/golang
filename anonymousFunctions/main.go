package main

import "fmt"

func main() {
	f := func(x,y int) int { return x + y}
	fmt.Println(f(3,4))
	fmt.Println(f(6,7))

	f = selectFn(1)
	fmt.Println(f(2,3))
	fmt.Println(f(4,5))
	fmt.Println(f(5,6))

	g := square()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
}

func selectFn(n int) (func(int, int) int){
	if n==1 {
		return func(x int, y int) int{ return x + y}
	}else if n==2{
		return func(x int, y int) int{ return x - y}
	}else{
		return func(x int, y int) int{ return x * y}
	}
}

func square() func() int {
	var x int = 2
	return func() int {
		x++
		return x*x
	}
}
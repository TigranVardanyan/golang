package main

import "fmt"
import "math"
import "./pkg"


//import (  //short type
//	"fmt"
//	"math"
//)

func main()  {
	fmt.Println("Hello Go")
	fmt.Println(math.Pow(4,2))
	pkg.Factorial(5)
}

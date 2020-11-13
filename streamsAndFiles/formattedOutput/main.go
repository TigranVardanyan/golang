package main

import (
	"fmt"
	"os"
)

func main() {
	fprintFprintln()
	tom := person{"Tom",24,77.5}
	file,err := os.Create("tom.dat")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fprintf(file, "%-10s %-10d %-10.3f\n",tom.name, tom.age, tom.weight)
}

func fprintFprintln()  {
	file, err := os.Create("confeve.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Fprint(file, "Today ")
	fmt.Fprintln(file, "good weather")
}

type person struct {
	name string
	age int
	weight float64
}

//func Fprint(w io.Writer, a ...interface{}) (n int, err error)  {}
//
//func Fprintln(w io.Writer, a ...interface{}) (n int, err error)  {}
//
//func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)

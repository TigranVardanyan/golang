package main

//func Fscan(r io.Reader, a ...interface{}) (n int, err error)
//func Fscanln(r io.Reader, a ...interface{}) (n int, err error)

import (
	"fmt"
	"os"
)

type person struct {
	name string
	age int
	weight float64
}

func main() {
	filename := "hello4.txt"
	writeData(filename)
	readData(filename)
}

func writeData(filename string) {
	var name string = "Tom"
	var age int = 24
	var weight float64 = 77.5

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fprintln(file, name)
	fmt.Fprintln(file, age)
	fmt.Fprintln(file, weight)

}
func readData(filename string)  {
	var name string
	var age int
	var weight float64

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fscanln(file, &name)
	fmt.Fscanln(file, &age)
	fmt.Fscanln(file, &weight)

	fmt.Println(name, age, weight)
}
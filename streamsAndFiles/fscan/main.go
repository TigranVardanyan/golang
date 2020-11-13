package main

//func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)

import (
	"fmt"
	"os"
)

type person struct {
	name string
	age int32
	weight float64
}
func main() {
	filename := "person.dat"
	writeData(filename)
	readData(filename)
}

func writeData(filename string){
	tom := person { name:"Tom", age: 24, weight: 68.5 }

	//short variant
	//tom := person{"Tom",24,68.5}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// save data
	fmt.Fprintf(file, "%s %d %.2f\n", tom.name, tom.age, tom.weight)
}
func readData(filename string){
	//vars for scan file
	var name string
	var age int
	var weight float64

	file, err := os.Open(filename)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = fmt.Fscanf(file, "%s %d %f\n", &name, &age, &weight)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%-8s %-8d %-8.2f\n", name, age, weight)
}
package main

import (
	"fmt"
	"os"
)

func main() {
	file,err := os.Create("Hello.txt")
	if err != nil {
		fmt.Println("Unable to create file", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Println(file.Name())
}


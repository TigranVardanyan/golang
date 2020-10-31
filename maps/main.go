package main

import (
	"fmt"
	"strconv"
)



func main() {
	//var people map[string]int

	var people = map[string]int{
		"Tom": 1,
		"Bob": 2,
		"Sam": 4,
		"Alice": 8,
	}
	fmt.Println(people)

	fmt.Println(people["Tom"])
	if val, ok := people["Sam"]; ok {
		fmt.Println(val)
		fmt.Println(ok)
	}
	delete(people, "Sam")
	for key,value := range people {
		fmt.Println(key + " " + strconv.Itoa(value))
	}

}

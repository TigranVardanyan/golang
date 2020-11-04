package main

import "fmt"

//for counter initialization;condition;counter change {
// action
//}

func main() {
	for i:=0;i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("")
	//counter init outside loop
	var j = 1
	for ; j < 10; j++{
		fmt.Println(j * j)
	}
	fmt.Println("")
	//only condition
	var k = 1
	for ; k < 10;{
		fmt.Println(k * k)
		k++
	}
	fmt.Println("")
	//only condition - short
	var p = 1
	for p < 10{
		fmt.Println(p * p)
		p++
	}
	fmt.Println("")
	for q := 1; q < 10; q++{
		for e := 1; e < 10; e++{
			fmt.Print(q * e, "\t")
		}
		fmt.Println()
	}

	//Iterating over arrays

	/*
	for index, value := range array{
		action
	}
	 */
	fmt.Println("")
	var users = [3]string{"Tom", "Alice", "Kate"}
	for index, value := range users{
		fmt.Println(index, value)
	}
	fmt.Println("")
	for _, value2 := range users{
		fmt.Println(value2)
	}

	//iterate array with standart loop
	var food = [5]string{"Apple", "Potato", "Pizza", "Cheese", "Banana"}
	fmt.Println("")
	for g := 0;g < len(food); g++ {
		if food[g] == "Potato" {
			continue
		}
		if food[g] == "Cheese" {
			break
		}
		fmt.Println(food[g])
	}
}

package main

import "fmt"

func main() {
	var users []string
	var users2 = []string{"Marko", "Antonio", "Kate"}

	fmt.Println(users2)

	for _,value := range users2 {
		fmt.Println(value)
	}

	users = append(users, "Zara", "Novak")
	fmt.Println(users)

	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	userGroupOne := initialUsers[:2]
	userGroupTwo := initialUsers[2:5]
	userGroupThree := initialUsers[5:]

	fmt.Println(userGroupOne)
	fmt.Println(userGroupTwo)
	fmt.Println(userGroupThree)
	newUsers := []int{1,2,3,4,5,6,7,8,9}
	newUsers = append(newUsers[:3], newUsers[3+1:]...) // append add element,
	// not array(slice) need to add spread operator
}



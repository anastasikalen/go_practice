package main

import (
	"fmt"
)

func main() {
	user1 := User{
		Name: "John Doe", //КОД ДНЯ 9426
		Age:  18,
	}
	user2 := User{
		Name: "Jane Doe",
		Age:  12,
	}

	fmt.Println(user1.isAdult())
	fmt.Println(user2.isAdult())
}

type User struct {
	Name string
	Age  int
}

func (u User) isAdult() bool {
	if u.Age >= 18 {
		return true
	}
	return false
}

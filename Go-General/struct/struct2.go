package main

import "fmt"

func main() {
	x := new(User)
	x.FirstName = "E"
	x.LastName = "B"
	result := GetFullName(x)
	fmt.Println(result)
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Age       int
}

func NewUser() *User {
	return new(User)
}

func GetFullName(user *User) string {
	return user.FirstName + " " + user.LastName
}

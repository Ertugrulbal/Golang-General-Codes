package main

import (
	"fmt"
	"strconv"
)

func main() {

	//printHello()
	//printX("hi!")
	// response := login("Hi")
	// println(response)
	// w(par(add(3, 4)))
	// x := add(3, 4)
	// w(par(x))
	// message := "Hello"
	// sayHello(&message)
	// fmt.Println(message)

	x, y := addMultipleInput(5, 5, 5, 5)
	println(x, y)

}

//variadic parameters
func addMultipleInput(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}
	return result, len(terms)
}

//pointer
func sayHello(msg *string) {
	fmt.Println(*msg)
	*msg = "Hi"
}

func w(p string) {
	fmt.Println(p)
}

func par(i int) string {
	return strconv.Itoa(i)
}

func add(x int, y int) int {
	return x + y
}

func printHello() {
	println("Hi")
}

func printX(y string) {
	println(y)
}

func login(uname string) string {
	return "Hoşgeldin, " + uname
}

func login2(uname, surname string) (string, string) {
	return uname, surname
}

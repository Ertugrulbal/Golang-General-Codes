package main

import (
	"fmt"
	"strconv"
)

func main() {
	// printHello()
	// printX("Hi!")
	// x := login1("Gamze")
	// printX(x)

	// x := add(3, 4)
	// w(par(x))

	// w(par(add(3, 4)))

	// message := "Hello"
	// sayHello(&message)
	// fmt.Println(message)

	x, y := add(3, 4, 5, 6, 7)
	fmt.Println(x, y)
}

// variadic parameters
func add(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}
	return result, len(terms)
}

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

// func add(x int, y int) int {
//   return x + y
// }

// func add(x, y int, z, o string) int {
//   return x + y
// }

// func printHello() {
//   println("Hi")
// }

// func printX(y string) {
//   fmt.Println(y)
// }

// func login1(uname string) string {
//   return "Hoşgeldin, " + uname
// }

// func login2(uname, surname string) (string, string) {
//   return uname, surname
// }

// public string abc()
// {

// }

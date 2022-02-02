package main

import "fmt"

// closure
// closure bir çeşit anonim fonksiyondur(anomymous function)
// parameter as a function

func main() {

	a, b := 5, 8
	fn := func(sum int) (int, int) {
		x := sum * a / b
		y := sum - x
		return x, y
	}

	splitValues(fn)

	// xx, yy := fn(10)
	// fmt.Println(xx, yy)
}

func split(sum int) (x int, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

func splitValues(f func(sum int) (int, int)) {
	x, y := f(40)
	fmt.Println(x, y)
}

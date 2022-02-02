package main

// named results

func main() {

}

// func ada() (int, int) {

// }

// func split(sum int) (int, int) {
//   x := sum * 4 / 9
//   y := sum - x
//   return x, y
// }

func split(sum int) (x int, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

func add(terms ...int) (numTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}
	return len(terms), sum
	// numTerms = len(terms)
	// return
}

// func X() (num1 int, num2 int, count int, number int, age int, bla int) {
//   return 234, 3432, 232, 1231, 989, 6
// }

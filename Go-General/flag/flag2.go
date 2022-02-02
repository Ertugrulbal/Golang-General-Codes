package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		name string
		age  int
		cute bool
	)
	flag.StringVar(&name, "name", defaultName(), "name of the gopher")
	flag.StringVar(&name, "n", defaultName(), "name of the gopher(shorthand)")

	flag.IntVar(&age, "age", defaultAge(), "age of the gopher")
	flag.IntVar(&age, "a", defaultAge(), "age of the gopher(shorthand)")

	flag.BoolVar(&cute, "cute", defaultCuteness(), "is the gopher cute")
	flag.BoolVar(&cute, "c", defaultCuteness(), "is the gopher cute(shorthand)")

	flag.Parse()

	fmt.Printf("Gopher stats\nName: %s\nAge: %d\nCute: %t", name, age, cute)

}

func defaultCuteness() bool {
	if os.Getenv("GOPHER_DEFAULT_CUTENESS") != "" {
		cuteness, err := strconv.ParseBool(os.Getenv("GOPHER_DEFAULT_CUTENESS"))
		if err == nil {
			return cuteness
		}
	}
	return false
}

func defaultAge() int {
	if os.Getenv("GOPHER_DEFAULT_AGE") != "" {
		age, err := strconv.Atoi(os.Getenv("GOPHER_DEFAULT_AGE"))
		if err == nil {
			return age
		}
	}
	return 0
}

func defaultName() string {
	if os.Getenv("GOPHER_DEFAULT_NAME") != "" {
		return os.Getenv("GOPHER_DEFAULT_NAME")
	}
	return "Gopher"
}

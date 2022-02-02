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
	flag.IntVar(&age, "a", defaultAge(), "age of the gopher (shorthand)")
	flag.BoolVar(&cute, "cute", defaultCuteness(), "is gopher cute")
	flag.BoolVar(&cute, "c", defaultCuteness(), "is gopher cute (shorthand)")
	flag.Parse()

	fmt.Printf("Name: %s\nAge: %d\nCute:%t\n", name, age, cute)
}

func defaultName() string {
	if os.Getenv("GOPHER_DEFAULT_NAME") != "" {
		return os.Getenv("gopher_default_name")
	}
	return "Gopher"
}

func defaultAge() int {
	if os.Getenv("GOPHER_DEFAULT_AGE") != "" {
		age, err := strconv.Atoi(os.Getenv("GOPHER_DEFAULT_AGE"))
		if err == nil {
			return age
		}
	}
	return 7
}

func defaultCuteness() bool {
	if os.Getenv("GOPHER_DEFAULT_CUTENESS") != "" {
		cute, err := strconv.ParseBool(os.Getenv("GOPHER_DEFAULT_CUTENESS"))
		if err == nil {
			return cute
		}
	}
	return true
}

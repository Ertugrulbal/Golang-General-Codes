package main

import (
	"fmt"
	"strings"
)

type Car struct {
	Make    string
	Model   string
	Options []string
}

func main() {
	dashes := strings.Repeat("-", 50)

	is250 := &Car{"Lexus", "IS250", []string{"GPS", "Alaşım Janslar", "Tavan Rafı", "Güç Çıkışları", "Isıtmalı Koltuk"}}
	accord := &Car{"Honda", "Accord", []string{"Alaşım Janslar", "Tavan Rafı"}}
	blazer := &Car{"Chevy", "Blazer", []string{"GPS", "Alaşım Janslar", "Tavan Rafı", "Güç Çıkışları"}}

	cars := []*Car{is250, accord, blazer} // cars is a slice!
	fmt.Printf("Cars:\n%v\n\n", cars)

	car_options := make(map[string][]*Car)

	fmt.Printf("Cars:\n%s\n", dashes)
	for _, car := range cars {
		fmt.Printf("%v\n", car)
		for _, option := range car.Options {
			car_options[option] = append(car_options[option], car)
			fmt.Printf("car options[option]: %s\n", option)
		}
		fmt.Println(dashes)
	}
	fmt.Println(dashes)

	for _, p := range car_options["GPS"] {
		fmt.Println(p.Make + " GPS var.")
	}

	fmt.Println("-->")
}

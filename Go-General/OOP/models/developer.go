package models

import "fmt"

type Developer struct {
	Employee // type embedding for composition
	Skills   []string
}

func (d Developer) PrintDetails() {
	d.Employee.PrintDetails()
	// d.PrintDetails()
	fmt.Println("Technical Skills: ")
	for _, v := range d.Skills {
		fmt.Println("-> " + v)
	}
}

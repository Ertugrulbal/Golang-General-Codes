package main

import (
	"fmt"

	"github.com/ertugrulbal/userpayment/models"
	"github.com/ertugrulbal/userpayment/utils"
	// . "github.com/ertugrulbal/userpayment/models"
	// x "github.com/ertugrulbal/userpayment/models"
	// _ "github.com/ertugrulbal/userpayment/models"
)

func main() {

	// v2

	// User Creation
	user := models.NewUser()
	user.FirstName = "Ertuğrul"
	user.LastName = "BAL"
	user.Age = 25
	user.UserName = "ertugrulbal"

	// Payment Creation v1

	// user.Pay.Salary = 50
	// user.Pay.Bonus = 5
	// // Printing
	// fmt.Println(user.Pay)
	// fmt.Println(user.GetPayment())

	// Payment Creation v2

	// user.Pay = &models.Payment{Salary: 45, Bonus: 3}
	// fmt.Println(user.GetFullName())
	// // fmt.Println(user.Pay)
	// fmt.Println(user.GetPayment())

	// Random

	// randX := utils.Random(10, 99)
	// fmt.Println("Rand Value: " + strconv.Itoa(randX))
	fmt.Println(GetSalary())

}

func GetSalary() int {
	return 500 + utils.Random(10, 99)
}

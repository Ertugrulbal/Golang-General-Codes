package models

type Payment struct {
	ID     int
	Salary float64
	Bonus  float64
}

func NewPayment() *Payment {
	return new(Payment)
}

func (u *User) GetPayment() float64 {
	return u.Pay.Salary + u.Pay.Bonus
}

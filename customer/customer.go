package customer

import (
	// "time"
	"math/rand"
)

type Customer struct {
	Id   int
	Name string
	//Code         string
	Amount float64
	// LoanDate     time.Time
	// LoanDuration int
	// ExpiryDate   time.Time
}

var Customers []Customer

func AddToCustomers(c *Customer) {
	Customers = append(Customers, *c)
}

var Names = []string{"Adeleke", "John", "Ahmed"}

func GenerateCustomers(customerCount int) []Customer {
	for i := 0; i <= customerCount; i++ {
		randomFirstName := Names[rand.Intn(len(Names))]
		randomLastName := Names[rand.Intn(len(Names))]
		customerName := randomFirstName + " " + randomLastName
		customer := &Customer{
			Id:     len(Customers) + 1,
			Name:   customerName,
			Amount: 1000,
		}
		AddToCustomers(customer)
	}
	return Customers
}

package customer

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	filemanager "github.com/adeisbright/2mdp/file_manager"
	"github.com/adeisbright/2mdp/lib"
)

type Customer struct {
	Id           int
	Name         string
	Code         string
	Amount       float64
	LoanDate     time.Time
	LoanDuration int
	ExpiryDate   time.Time
}

var Customers []Customer

func AddToCustomers(c *Customer) {
	Customers = append(Customers, *c)
}

var Names = []string{
	"Adeleke",
	"John",
	"Ahmed",
	"Benita",
	"Emeka",
	"Okon",
	"Mogaji",
	"Bello",
	"Abdulrahman",
	"Uchenna",
	"Daberechi",
	"Amadi",
	"Orkan",
	"Pepple",
	"Adekunle",
	"Idongesit",
	"Owoicho",
	"Yohana",
	"Galadima",
	"Ifiok",
	"Essien",
	"Gbenga",
	"Okezie",
	"Obi",
	"Anyanwu",
}

const HIGHEST_OWED_AMOUNT_CAP = 1000000

func (c *Customer) StringnifyCustomer() string {
	return fmt.Sprintf("%d ,  %s , %s ,  %f  , %q , %d , %q \n", c.Id, c.Name, c.Code, c.Amount, c.LoanDate, c.LoanDuration, c.ExpiryDate)
}
func GenerateCustomers(customerCount int) []Customer {
	for i := 0; i <= customerCount; i++ {
		randomFirstName := Names[rand.Intn(len(Names))]
		randomLastName := Names[rand.Intn(len(Names))]
		customerName := randomFirstName + " " + randomLastName
		randomAmount := rand.Float64() * HIGHEST_OWED_AMOUNT_CAP
		customerCode, err := lib.GenerateRandomString(9)

		if err != nil {
			fmt.Println(err.Error())
		}

		randomDuration := rand.Intn(12) + 1
		customer := &Customer{
			Id:           len(Customers) + 1,
			Name:         customerName,
			Amount:       randomAmount,
			Code:         customerCode,
			LoanDate:     time.Now(),
			LoanDuration: randomDuration,
			ExpiryDate:   time.Now().Local().AddDate(0, randomDuration, 0),
		}
		AddToCustomers(customer)
		customerString := customer.StringnifyCustomer()
		filemanager.AddToFile("customers.txt", customerString)
	}
	return Customers
}

func DisplayCustomers(customers []Customer) {
	fmt.Println("A view of Generated Customers")
	for _, customer := range customers {
		fmt.Println(customer)
	}
}

func GetCustomer(customerCode string, customers []Customer) (*Customer, error) {
	for _, value := range customers {
		if value.Code == customerCode {
			return &value, nil
		}
	}
	return nil, errors.New("not found")
}

package main

import (
	"fmt"

	"github.com/adeisbright/2mdp/customer"
)

func main() {
	fmt.Println("Start of the 2 Million Debtors Program")
	customers := customer.GenerateCustomers(50)
	fmt.Println(customers)
}

package main

import (
	"fmt"

	"github.com/adeisbright/2mdp/customer"
	filemanager "github.com/adeisbright/2mdp/file_manager"
)

func main() {
	fmt.Println("Start of the 2 Million Debtors Program")
	customersFile, err := filemanager.CreateFile("debtors.txt")
	if err != nil {
		fmt.Println(err)
	}
	//Generate Customers and store them in the provided store
	customers := customer.GenerateCustomers(2000000, customersFile.Name())
	fmt.Println(len(customers))
}

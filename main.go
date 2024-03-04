package main

import (
	"fmt"

	//filemanager "github.com/adeisbright/2mdp/file_manager"
	"github.com/adeisbright/2mdp/customer"
)

func main() {
	fmt.Println("Start of the 2 Million Debtors Program")
	customers := customer.GenerateCustomers(2000000)
	fmt.Println(len(customers))
	// // customer.DisplayCustomers(customers)
	// filemanager.ReadFromFile("Readme.md")
	// customersFile, err := filemanager.CreateFile("customers.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// filemanager.AddToFile(customersFile.Name(), "Hello, this is a test")
}

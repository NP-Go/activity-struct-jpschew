package main

import "fmt"

//Insert struct declaration here

// activity 1
type customer struct {
	fname            string
	lname            string
	age              int
	subscriber       bool
	homeAddress      string
	phone            int
	creditAvailable  float32
	currentCartCost  float32
	currentOrderCost float32
}

func main() {
	//Insert code here

	// activity 2
	customer1 := customer{
		fname:            "Annakin",
		lname:            "Skywalker",
		age:              45,
		subscriber:       true,
		homeAddress:      "Death Star",
		phone:            1234567,
		creditAvailable:  10000.00,
		currentCartCost:  0.00,
		currentOrderCost: 0.00,
	}

	customer2 := customer{
		fname:            "Han",
		lname:            "Solo",
		age:              50,
		subscriber:       false,
		homeAddress:      "Tatooine",
		phone:            4321765,
		creditAvailable:  20000.00,
		currentCartCost:  0.00,
		currentOrderCost: 0.00,
	}

	fmt.Println(customer1)
	fmt.Println(customer2)

}

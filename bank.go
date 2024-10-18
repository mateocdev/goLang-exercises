package main

import "fmt"

func main() {
	var accountBalance float64 = 1000.0

	for {
		fmt.Println("Please select an option:")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Println("Enter the deposit amount:")
			fmt.Scanln(&depositAmount)
			accountBalance += depositAmount
		case 3:
			var withdrawAmount float64
			fmt.Println("Enter the withdraw amount:")
			fmt.Scanln(&withdrawAmount)
			accountBalance -= withdrawAmount
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Your final balance is:", accountBalance)
			return
		default:
			fmt.Println("Invalid choice")
		}
	}

}

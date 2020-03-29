package main

import (
	"fmt"
)

func main() {
	input := -1
	for input != 0 {
		fmt.Printf("1 – Print Covid19 cases in Pakistan\n2 – Print Covid19 cases in SouthKorea\n3 – Print Covid19 cases in France\n4 – Print my personalized message to Coronavirus\n0 – Exit\n>> ")
		fmt.Scan(&input)
		if input == 1 {
			fmt.Println("Covid19 cases in Pakistan: 1,526")
		} else if input == 2 {
			fmt.Println("Covid19 cases in SouthKorea: 2000")
		} else if input == 3 {
			fmt.Println("Covid19 cases in Pakistan: 3000")
		} else if input == 4 {
			fmt.Println("You can spread the virus even if you don’t have symptoms.")
		} else if input != 0 {
			fmt.Println("Wrong Option.")
		}
	}
	fmt.Println("Exiting...")
}

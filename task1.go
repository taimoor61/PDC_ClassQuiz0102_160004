package quiz0102

import "fmt"

func Menu() { //To display the stats
	input := -1
	pak, sk, fr := false, false, false
	for true {
		fmt.Printf("1 – Print Covid19 cases in Pakistan\n2 – Print Covid19 cases in SouthKorea\n3 – Print Covid19 cases in France\n4 – Print my personalized message to Coronavirus\n0 – Exit\n>> ")
		fmt.Scan(&input)
		if input == 1 {
			fmt.Println("Covid19 cases in Pakistan: 1,526")
			pak = true
		} else if input == 2 {
			fmt.Println("Covid19 cases in SouthKorea: 9,583")
			sk = true
		} else if input == 3 {
			fmt.Println("Covid19 cases in France: 37,575")
			fr = true
		} else if input == 4 {
			fmt.Println("You can spread the virus even if you don’t have symptoms.")
		} else if input == 0 {
			if pak && sk && fr {
				break
			} else {
				fmt.Println("You have not seen all the stats yet")
			}
		} else {
			fmt.Println("Wrong Option")
		}

	}
	fmt.Println("Exiting...")
}

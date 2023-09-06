package main

import (
	"fmt"
	"slices"
)

// global variable
var religions = []string{"catholique", "protestant", "athee"}

func main() {

	var born_year int

	fmt.Println("what year were you born")
	fmt.Scan(&born_year)
	my_age = Get_age(born_year)

	//get my religion
	my_religion := Get_religion(religions)
	fmt.Println("your religin is", my_religion)

	//check if you can drink

	if()

	/*
		age := Get_age(born_year)
		if age < 18 {
			fmt.Println(" cant drink")
		} else {
			fmt.Println("GET DRUNK")
		}
	*/
}

//END MAIN

//FUNCTIONS

// return your age based on year born
func Get_age(year_born int) int {

	var current_year = 2023
	var age int

	age = current_year - year_born
	return age

}

// return a string religion based on slice of religions
func Get_religion(array_religion []string) string {

	for {
		var my_religion string

		fmt.Println(" here are the religions :")
		for _, s := range array_religion {
			fmt.Println(s)
		}

		fmt.Scan(&my_religion)

		// check if contains
		exist := slices.Contains(array_religion, my_religion)

		if exist {
			return my_religion
		}

		fmt.Println("didnt get your choise , try again")
	}

}

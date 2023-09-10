package main

import "fmt"

func main() {
	var number int
	fmt.Println("enter a number: ")
	fmt.Scanln(&number)

	if number%2 == 0 {
		fmt.Println("this is a even number. ")
	} else {
		fmt.Println("thsi is odd number. ")
	}

}

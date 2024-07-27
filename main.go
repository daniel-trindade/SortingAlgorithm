package main

import (
	"fmt"
	"main/sorting"
)

//MAIN FUNCTION
func main() {

	var num [20]int = [20]int{20, 5, 68, 15, 9, 1, 44, 4, 88, 25, 10, 3, 99, 77, 55, 100, 2, 7, 6, 42}

	fmt.Println("Array desordenado:")
	for i:=0; i<len(num); i++{
		fmt.Print(num[i], " ")
	}

	fmt.Println()

	sorting.InsertionSort(num[:])
	
	fmt.Println("Array ordenado:")
	for i:=0; i<len(num); i++{
		fmt.Print(num[i], " ")
	}
	fmt.Println()
}

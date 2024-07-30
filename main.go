package main

import (
	"fmt"
	"main/sorting"
)

func main() {

	var num [9]int = [9]int{3, 6, 2, 5, 4, 3, 7, 1, 10}

	fmt.Println("Array desordenado:")
	for i:=0; i<len(num); i++{
		fmt.Print(num[i], " ")
	}

	fmt.Println()

	sorting.CountingSort(num[:])
	
	fmt.Println("Array ordenado:")
	for i:=0; i<len(num); i++{
		fmt.Print(num[i], " ")
	}
	fmt.Println()
}

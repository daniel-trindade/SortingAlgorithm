package main

import (
	"fmt"
)

//MAIN FUNCTION
func main() {

	var num [20]int = [20]int{20, 5, 68, 15, 9, 1, 44, 4, 88, 25, 10, 3, 99, 77, 55, 100, 2, 7, 6, 42}

	fmt.Println("Array desordenado:")
	for i:=0; i<len(num); i++{
		fmt.Print(num[i], " ")
	}

	fmt.Println()

	SelectionSort(num[:])
	
	fmt.Println("Array ordenado pelo Selection Sort:")
	for i:=0; i<len(num); i++{
		fmt.Print(num[i], " ")
	}
	fmt.Println()
}



//SELECTION SORT
func SelectionSort(arr[]int){

	var min int
	var index int

	for i:=0; i<len(arr); i++{
		min = arr[i]
		index = i
		for j:=i+1; j < len(arr); j++{
			if arr[j] < min{
				min = arr[j]
				index = j
			}
		}
		arr[index] = arr[i]
		arr[i] = min
	}
}

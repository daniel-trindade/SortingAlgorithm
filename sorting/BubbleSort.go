package sorting

import "fmt"

func BubbleSort(arr[]int){

  var sorted bool = false
  var temp int

  for sorted!=true{
    sorted = true
    for i:=0; i<len(arr)-1; i++{
      if arr[i] > arr[i+1] {
        sorted = false
        temp = arr[i]
        arr[i] = arr[i+1]
        arr[i+1] = temp
      }
    }
    for j:=0;j<len(arr);j++{
      fmt.Print(arr[j], " ")
    }
    fmt.Println()
  } 
}
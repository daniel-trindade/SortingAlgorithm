package sorting

import "fmt"

func InsertionSort(arr[] int){

  var temp int
  
  for i:=1;i<len(arr);i++{
    for j:=i; j>0 && arr[j]<arr[j-1]; j--{
      temp = arr[j-1]
      arr[j-1] = arr[j]
      arr[j] = temp
    }
    for j:=0;j<len(arr);j++{
      fmt.Print(arr[j], " ")
    }
    fmt.Println()
  }
  
}
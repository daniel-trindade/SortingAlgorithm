package sorting

func SelectionSort(arr[]int){

  var temp int
  var index int
  
  for i:=0; i<len(arr); i++{
    for j:=0; j < len(arr); j++{
      if arr[i] > arr[j]{
        temp = arr[j]
        index = j
      }
    }
    arr[index] = arr[i]
    arr[i] = temp
  }
}
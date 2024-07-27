package sorting

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
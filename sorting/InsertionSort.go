package sorting

func InsertionSort(arr[] int){

  var temp int
  
  for i:=1;i<len(arr);i++{
    for j:=i; j>0 && arr[j]<arr[j-1]; j--{
      temp = arr[j-1]
      arr[j-1] = arr[j]
      arr[j] = temp
    }
  }
}
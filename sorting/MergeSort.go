package sorting

func MergeSort (arr[]int){
  if len(arr) > 1{
    
    mid := len(arr)/2

    leftValues := make([]int, mid)
    for i:=0; i<mid; i++{
      leftValues[i]=arr[i]
    }

    rightValues := make([]int, len(arr)-mid)
    for i:=mid; i<len(arr); i++{
      rightValues[i-mid]=arr[i]
    }

    MergeSort(leftValues)
    MergeSort(rightValues)
    Merge(arr, leftValues, rightValues)
  }
}

func Merge (arr[]int, arrLeft[] int, arrRight[]int){
  iArr := 0
  iLeft := 0
  iRight := 0

  for iLeft < len(arrLeft) && iRight < len(arrRight){
    if arrLeft[iLeft] <= arrRight[iRight]{
      arr[iArr] = arrLeft[iLeft]
      iLeft++
    }else{
      arr[iArr] = arrRight[iRight]
      iRight++
    }
    iArr++
  }

  for iLeft < len(arrLeft){
    arr[iArr] = arrLeft[iLeft]
    iLeft++
    iArr++
  }

  for iRight < len(arrRight){
    arr[iArr] = arrRight[iRight]
    iRight++
    iArr++
  }
  
}
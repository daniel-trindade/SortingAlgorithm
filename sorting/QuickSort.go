package sorting

import (
	"math/rand"
  "time"
)

func QuickSort(arr[] int, start int, end int){
  
  if start < end{
    iPivot := partition(arr, start, end)
    QuickSort(arr, start, iPivot-1)
    QuickSort(arr, iPivot+1, end)
  }
  
}

func partition(arr[]int, start int, end int) int{
  
  rand.Seed(time.Now().UnixNano())
  randPivot := rand.Intn(end-start) + 1

  arr[end], arr[randPivot] = arr[randPivot], arr[end]
  iPivot := end
  pIndex := start

  for i:= start; i<iPivot; i++{
    if arr[i] <= arr[iPivot]{
      arr[i], arr[iPivot] =  arr[iPivot], arr[i]
      pIndex++
    }
  }

  arr[pIndex], arr[iPivot] = arr[iPivot], arr[pIndex]

  return pIndex
}
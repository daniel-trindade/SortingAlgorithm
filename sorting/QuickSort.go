package sorting

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )



func QuickSort(arr[] int, start int, end int){
  
  if start < end{
    iPivot := partition(arr, start, end)
    QuickSort(arr, start, iPivot-1)
    QuickSort(arr, iPivot+1, end)
  }
  
}

func partition(arr[]int, start int, end int) int{
  
  // rand.Seed(time.Now().UnixNano())
  // randPivot := rand.Intn(end-start+1)

  // fmt.Println("indice do pivot: ", randPivot, " Valor: ", arr[randPivot])

  // arr[end], arr[randPivot] = arr[randPivot], arr[end]

  // fmt.Println("Valor do umtimo indice: ", arr[end])
  
  pivot := arr[end]
  pIndex := start

  for i:= start; i<end; i++{
    if arr[i] <= pivot{
      arr[pIndex], arr[i] =  arr[i], arr[pIndex]
      pIndex++
    }
  }

  arr[pIndex], arr[end] = arr[end], arr[pIndex]

  return pIndex
}
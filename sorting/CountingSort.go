package sorting

func CountingSort (arr[]int){

  max := 0
  
  for i:=0; i<len(arr); i++{
    if arr[i]>max{
      max=arr[i]
    }
  }

  max++
  
  arrAux := make([]int, max)

  for i:=0; i<max;i++{
    arrAux[i] = 0
  }

  for i:=0; i<len(arr); i++{
    aux := arr[i]
    arrAux[aux] += 1 
  }

  c := 0
  
  for i:=0; i<len(arrAux); i++{
    for j:=0; j<arrAux[i]; j++{
      arr[c] = i
      c++
    }
  }
}
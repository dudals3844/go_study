package main

import "fmt"

func main() {
    sliceA := []int{1,2,3}
    sliceB := []int{4,5,6}
    
    sliceA = append(sliceA,sliceB...)//sliceB...은 모든 배열에 있는 값을 추가한다는 의미

    fmt.Println(sliceA)
}

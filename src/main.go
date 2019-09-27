package main

import "fmt"

func main() {
	s := []int{0,1}
	s = append(s,2)

    s = append(s,3,4,5)

    fmt.Println(s)
}

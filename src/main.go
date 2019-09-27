package main

import "fmt"

//define struct
type person struct {
    name string
    age int
}


func main() {
    p := person{}//객체 생성
    d := new(person)
    //필드값 설정
    p.name = "Lee"
    p.age = 20
    
    d.name = "Choi"
    d.age = 21
    
    fmt.Println(d)
    
    fmt.Println(p)
}

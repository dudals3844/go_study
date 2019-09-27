package main

import (
    //"fmt"
    //"math"
)
//define Rect
type Rect struct {
    width, height float64//return type이랑 맞춰줘야 된다.
}

type Circle struct {
    radius float64
}
//Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area2() float64 {
    
    return r.width * r.height
}

// func (r Rect) perimeter() float64{
//     return 2 * (r.width+r.height)
// }

// //Circle 타입에 대한 Shape 인터페이스 구현
// func (c Circle) area() float64{
//     return math.Pi * c.radius * c.radius
// }
// func (c Circle)perimeter()float64{
//     return 2*math.Pi*c.radius
// }

func main() {
    r := Rect{10.,20.}
    //c := Circle{10.}
    println(r.area2())
    //println(c.perimeter())
}

// func showArea(shapes ...Shape){
//     for _,s := range shapes{
//         a := s.area()//인터페이스 메서드 호출
//         println(a)
//     }
// }
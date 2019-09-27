package main

//import "fmt"

type Rect struct {
    width, height int
}
 
func (r Rect) area() int {
    return r.width * r.height
}

func main() {
    rect := Rect{10,20}
    area := rect.area()//메소드 호출
    println(area)
}

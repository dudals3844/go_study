package main

//import "fmt"

type Rect struct {
    width, height int
}
 
func (r *Rect) area2() int {
    r.width++
    return r.width * r.height
}

func main() {
    rect := Rect{10,20}
    area := rect.area2()//메소드 호출
    println(rect.width,area)
}

package main

import (
    //"fmt"
    //"math"
    //"log"
    "os"
)


func main() {
    openFile("invalid.txt")
    println("Done")
    
}

func openFile(fn string){
    f, err := os.Open(fn)
    if err != nil {
        panic(err)//panic함수는 현재함수를 즉시 멈추고 현재함수의 defer함수들을 모두 실행한뒤 즉시 리턴한다.
    }
    
    defer f.Close()
}



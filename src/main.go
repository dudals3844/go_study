package main

import (
    //"fmt"
    //"math"
    //"log"
    "os"
)


func main() {
    f , err := os.Open("1.txt")
    if err != nil {
        panic(err)
    }
    
    //main 마지막에 파일 close
    defer f.Close()
    bytes := make([]byte, 1024)
    f.Read(bytes)
    println(len(bytes))
    
}



package main

import (
    //"fmt"
    //"math"
    "log"
    "os"
)
type error interface {
    Error() string
}

func main() {
    f, err := os.Open("./")
    if(err != nil){
       log.Fatal(err.Error())
    }
    println(f.Name())
    
    _,err := otherFunc()
    switch err.(type){
    default: //no error
        println("ok")
    case Myerror:
        log.Print("Log my error")
    case error:
        log.Fatal(err.Error())
            
    }
}



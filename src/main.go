package main

import (
	//"fmt"
	//"math"
	//"log"
	//"os"
	//"time"
	//"sync"
)

func main() {
	ch := make(chan int, 2)
	
	ch <- 1
	ch <- 2
	
	// 채널이 닫힌 것을 감지할 때까지 계속 수신
    /*
    for {
        if i, success := <-ch; success {
            println(i)
        } else {
            break
        }
    }
    */
 
    // 방법2
    // 위 표현과 동일한 채널 range 문
	
	for i := range ch {
		println(i)
	
	}
	
}

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

	close(ch)

	println(<-ch)
	println(<-ch)

	if _, success := <-ch; !success { //_,success := <-ch은 값을 대입받는다는 의미이다.
		println("no more data")
	}
}

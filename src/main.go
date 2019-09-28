package main

import (
	"fmt"
	//"math"
	//"log"
	//"os"
	//"time"
	//"sync"
)

func main() {
	// c := make(chan int)
	// c <- 1           //수신루틴이 없어 데드락
	// fmt.Println(<-c) //별도의 고루틴이 없어서 데드락에 걸린다

	//buffer channel
	ch := make(chan int, 1) //chan int, 1에서 1이 버퍼의 개수이다.
	//수신자가 없어도 보낼 수 있다
	ch <- 101
	fmt.Println(<-ch)
}

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
	ch := make(chan string, 1)
	sendChan(ch)
	receiveChan(ch)
}

func sendChan(ch chan<- string) {
	ch <- "Data"
	//x := <- ch
}

func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}

// 해당 채널로 송신만 할 것인지 혹은 수신만할 것인지를 지정할 수도 있다. 송신 파라미터는 (p chan<- int)와 같이 chan<- 을 사용하고,
// 수신 파라미터는 (p <-chan int)와 같이 <-chan 을 사용한다. 만약 송신 채널 파라미터에서 수신을 한다거나,
// 수신 채널에 송신을 하게되면, 에러가 발생한다.

// 아래 예제에서 만약 sendChan() 함수 안에서 x := <- ch 를 실행하면 송신전용 채널에 수신을 시도하므로 에러가 발생한다.

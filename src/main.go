package main

import (
	"fmt"
	//"math"
	//"log"
	//"os"
	"time"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	//동기 실행
	say("Sync")

	//비동기 실행
	go say("Async1")
	go say("Async2")
	go say("Async3")

	time.Sleep(time.Second * 3)
}

package main

import (
	"fmt"
	//"math"
	//"log"
	//"os"
	//"time"
	"sync"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	var wait sync.WaitGroup
	wait.Add(2)

	//익명함수를 이용한 goroutine
	go func() {
		defer wait.Done() //끝나면 .Done()호출
		fmt.Println("Hello")
	}()

	go func(msg string) {
		defer wait.Done() //끝나면 .Done()호출
		fmt.Println(msg)
	}("HI")

	wait.Wait()
}

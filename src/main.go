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
	ch := make(chan int) //채널의 데이터 타입을 정해준다

	go func() {
		ch <- 123 //고루틴문에서 채널에 123을 보낸다
	}()

	var i, j int
	i = <-ch //메인루티에서 채널에 데이저를 받는다
	println(i)

	j = <-ch
	println(j)
}

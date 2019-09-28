package main

import (
	//"fmt"
	//"math"
	//"log"
	"os"
	//"time"
	//"sync"
	"io"
)

func main() {
	fi, err := os.Open("/workspace/go/src/github.com/demo-apps/go-gin-app/Go_Study/src/1.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	//create ouput file
	fo, err := os.Create("/workspace/go/src/github.com/demo-apps/go-gin-app/Go_Study/src/2.txt")
	if err != nil{
		panic(err)
	}
	defer fo.Close()
	buff := make([]byte, 1024)

	//loop
	for {
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if cnt == 0{
			break
		}

		//write
		_, err = fo.Write(buff[:cnt])
		if err != nil{
			panic(err)
		}
	}
}



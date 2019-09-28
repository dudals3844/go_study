package main

import (
	"fmt"
	//"math"
	//"log"
	"os"
)

func main() {
	openFile("1.txt")
	println("Done")

}

func openFile(fn string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Open Error", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()
}

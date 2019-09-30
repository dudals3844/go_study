package main

import (
	//"fmt"
	//"math"
	"log"
	"os"
	//"time"
	//"sync"
	//"bytes"
	//"io/ioutil"
	//"net/http"
	//"net/url"
	//"encoding/xml"
)

var myLogger *log.Logger

func main() {
	fpLog, err := os.OpenFile("/src/1.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	defer fpLog.Close()

	myLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	run()
	myLogger.Println("End of Program")
}

func run() {
	myLogger.Print("Test")
}

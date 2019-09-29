package main

import (
	//"fmt"
	//"math"
	//"log"
	//"os"
	//"time"
	//"sync"
	"io/ioutil"
	"net/http"
	"bytes"
	//"net/url"
	"encoding/xml"
)


type Person struct {
	Name string
	Age  int
}
//HTTP Postform 호출
func main() {
	person := Person{"Alex",10}
	pbytes, _ := xml.Marshal(person)
	buff := bytes.NewBuffer(pbytes)

	//Create Request obj
	req, err := http.NewRequest("POST","http://httpbin.org/post",buff)
	if err != nil{
		panic(err)
	}

	//Content-Type header add
	req.Header.Add("Content-Type","application/xml")

	//Request start in Client obj
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}



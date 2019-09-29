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
	//"bytes"
	"net/url"
)
//HTTP Postform 호출
func main() {
	resp, err := http.PostForm("http://httpbin.org/post",url.Values{"Name":{"Lee"},"Age": {"10"}})
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	//Response check
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}



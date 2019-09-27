package main

//import "fmt"

func main() {
    tickers := map[string]string{
        "GOOG": "Google Inc",
        "MSFT": "Microsoft",
        "FB":   "FaceBook",
        "AMZN": "Amazon",//마지막도 , 넣어야 된다.
    }
    
    // map key check
    val,exists := tickers["MSFT"]
    if !exists {
        print("NO MSFT tickers")
    }
    
    println(val)//키가 존재하면 Microsoft를 리턴한다.
}

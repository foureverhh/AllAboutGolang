package main

import (
	"encoding/xml" //parse xml object to go object
	"fmt"
	"io/ioutil" //receive and parse responds
	"net/http"  //send request
	"time"
)

func main() {
	start := time.Now()

	resp, error := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=googl")
	if error != nil {
		fmt.Println("http response error occurs!")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	quote := new(QuoteResponse)
	xml.Unmarshal(body, &quote)
	fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)

	elapsed := time.Since(start)

	fmt.Printf("Execution time: %s\n", elapsed)
}

//QuoteResponse is representing the whole xml respence from markitondemond api
type QuoteResponse struct {
	Status           string
	Name             string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	TimeStamp        string
	MSDate           float32
	MarketCap        int
	Volume           int
	ChangeYTD        float32
	ChangePercentYTD float32
	High             float32
	Low              float32
	open             float32
}

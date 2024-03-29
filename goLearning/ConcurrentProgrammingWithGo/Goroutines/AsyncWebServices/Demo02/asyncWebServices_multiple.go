package main

import (
	"encoding/xml" //parse xml object to go object
	"fmt"
	"io/ioutil" //receive and parse responds
	"net/http"  //send request
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	start := time.Now()

	stockSymbols := []string{
		"googl",
		"msft",
		"aapl",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmus",
		"s",
	}

	numComplete := 0

	for _, symbol := range stockSymbols {
		go func(symbol string) {
			resp, error := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=" + symbol)
			if error != nil {
				fmt.Println("http response error occurs!")
			}
			defer resp.Body.Close()

			body, _ := ioutil.ReadAll(resp.Body)

			quote := new(QuoteResponse)
			xml.Unmarshal(body, &quote)
			fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
			numComplete++
		}(symbol)
	}

	for numComplete < len(stockSymbols) {
		time.Sleep(10 * time.Millisecond)
	}

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

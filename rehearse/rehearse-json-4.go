package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// read file!
	data, _ := ioutil.ReadFile("api_keys")

	api_url := fmt.Sprintf("http://api.marketstack.com/v1/tickers?access_key=%v&exchange=XNAS&symbols=AAPL,MSFT", string(data))

	resp, _ := http.Get(api_url)

	// close HTTP connection!
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)

	fmt.Println(bodyString)
}

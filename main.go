package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
)

func main() {

	godotenv.Load()
	key := os.Getenv("API_KEY")
	symbol := "AAPL" // Replace with the ticker symbol you want

	if len(os.Args) >= 2 {
		symbol = os.Args[1]
	}

	url := fmt.Sprintf("https://api.twelvedata.com/time_series?symbol=%s&apikey=%s", symbol, key)

	//get req
	req, err := http.NewRequest("GET", url, nil)
	// req.Header.Add("Authorization", "apikey "+key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Stock API not Available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

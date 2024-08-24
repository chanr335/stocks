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
	url := "https://api.finazon.io/latest/finazon/us_stocks_essential/tickers?page_size=1000"

	//get request
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "apikey "+key)
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

	fmt.Println(res)
	fmt.Println(string(body))
}

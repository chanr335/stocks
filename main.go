package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
)

type Stock struct {
	Data []struct {
		AssetType     string `json:"asset_type"`
		Cik           string `json:"cik"`
		CompositeFigi string `json:"composite_figi"`
		Currency      string `json:"currency"`
		Lei           string `json:"lei"`
		Mic           string `json:"mic"`
		Security      string `json:"security"`
		ShareFigi     string `json:"share_figi"`
		Ticker        string `json:"ticker"`
	} `json:"data"`
	Meta struct {
		Pagination struct {
			Page    int `json:"page"`
			PerPage int `json:"per_page"`
		} `json:"pagination"`
	} `json:"meta"`
}

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

	// fmt.Println(string(body))

	var stock Stock
	err = json.Unmarshal(body, &stock)
	if err != nil {
		panic(err)
	}

	fmt.Println(stock)
}

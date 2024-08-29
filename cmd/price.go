/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
)

type Fund struct {
	Price string `json:"price"`
}

// priceCmd represents the price command
var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "A brief description of your command",
	Long: `Get the price of a stock, \n
  Usage: price\(ticker, datetime\)`,

	Args: cobra.RangeArgs(1, 3),
	Run: func(cmd *cobra.Command, args []string) {
		godotenv.Load()
		key := os.Getenv("API_KEY")
		symbol := args[0] // Replace with the ticker symbol you want
		datetime := ""
		if len(args) > 1 {
			date := args[1]
			time := args[2]
			datetime = date + " " + time
		}

		url := fmt.Sprintf("https://api.twelvedata.com/price?symbol=%s&datetime=%s&apikey=%s", symbol, datetime, key)

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

		var fund Fund
		err = json.Unmarshal(body, &fund)
		if err != nil {
			panic(err)
		}

		price := fund.Price
		fmt.Printf("%s's Current Price: $%s USD\n", symbol, price)
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)
}

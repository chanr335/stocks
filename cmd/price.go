/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
)

// priceCmd represents the price command
var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "A brief description of your command",
	Long:  ``,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		godotenv.Load()
		key := os.Getenv("API_KEY")
		symbol := args[0] // Replace with the ticker symbol you want

		url := fmt.Sprintf("https://api.twelvedata.com/price?symbol=%s&apikey=%s", symbol, key)

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
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)
}

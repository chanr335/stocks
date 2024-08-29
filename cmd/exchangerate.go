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

type ExchangeRate struct {
	Symbol string  `json:"symbol"`
	Rate   float64 `json:"rate"`
}

// exchangerateCmd represents the exchangerate command
var exchangerateCmd = &cobra.Command{
	Use:   "exchangerate",
	Short: "A brief description of your command",
	Long:  ``,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		godotenv.Load()
		key := os.Getenv("API_KEY")
		symbol := args[0]

		url := fmt.Sprintf("https://api.twelvedata.com/exchange_rate?symbol=%s&apikey=%s", symbol, key)
		req, err := http.NewRequest("GET", url, nil)

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

		var exchangerate ExchangeRate
		err = json.Unmarshal(body, &exchangerate)
		if err != nil {
			panic(err)
		}

		fmt.Printf("The exchange rate for %s is %f", symbol, exchangerate.Rate)

	},
}

func init() {
	rootCmd.AddCommand(exchangerateCmd)
}

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

type Quote struct {
	L52 struct {
		Low        string `json:"low"`
		High       string `json:"high"`
		Lowchange  string `json:"low_change"`
		Highchange string `json:"high_change"`
	} `json:"fifty_two_week"`
}

// exchangerateCmd represents the exchangerate command
var quoteCmd = &cobra.Command{
	Use:   "quote",
	Short: "A brief description of your command",
	Long:  ``,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		godotenv.Load()
		key := os.Getenv("API_KEY")
		symbol := args[0]

		url := fmt.Sprintf("https://api.twelvedata.com/quote?symbol=%s&apikey=%s", symbol, key)
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

		// fmt.Println(string(body))

		var quote Quote
		err = json.Unmarshal(body, &quote)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Quote for %s: \nLow %s\nHigh: %s\n", symbol, quote.L52.Low, quote.L52.High)

	},
}

func init() {
	rootCmd.AddCommand(quoteCmd)
}

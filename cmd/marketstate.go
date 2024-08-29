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

type MarketState struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	Country string `json:"country"`
	Open    bool   `json:"is_market_open"`
	After   string `json:"time_after_open"`
	ToOpen  string `json:"time_to_open"`
	ToClose string `json:"time_to_close"`
}

// exchangerateCmd represents the exchangerate command
var marketstateCmd = &cobra.Command{
	Use:   "marketstate",
	Short: "A brief description of your command",
	Long:  ``,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		godotenv.Load()
		key := os.Getenv("API_KEY")
		name := args[0]

		url := fmt.Sprintf("https://api.twelvedata.com/market_state?name=%s&apikey=%s", name, key)
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

		// Unmarshal into a slice of MarketState
		var marketStates []MarketState
		err = json.Unmarshal(body, &marketStates)
		if err != nil {
			panic(err)
		}

		// Loop through the slice and print each market's state
		for _, marketstate := range marketStates {
			fmt.Printf("Name: %s, Country: %s, Is Market Open: %t, Time After Open: %s, Time to Open: %s, Time to Close: %s\n",
				marketstate.Name, marketstate.Country, marketstate.Open, marketstate.After, marketstate.ToOpen, marketstate.ToClose)
		}

	},
}

func init() {
	rootCmd.AddCommand(marketstateCmd)
}

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
	Data []struct {
		Name string `json:"name"`
		Open string `json:"is_market_open"`
	} `json:"data"`
}

// exchangerateCmd represents the exchangerate command
var marketstateCmd = &cobra.Command{
	Use:   "exchangerate",
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

		var marketstate MarketState
		err = json.Unmarshal(body, &marketstate)
		if err != nil {
			panic(err)
		}

		for _, marketstate := range marketstate.Data {
			fmt.Printf("The market is open: %s", marketstate.Open)
		}

	},
}

func init() {
	rootCmd.AddCommand(marketstateCmd)
}

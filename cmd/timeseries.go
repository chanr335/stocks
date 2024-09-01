package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"stocks_cli/cmd/model"
	"stocks_cli/cmd/utils" // Import utils package
)

type Timeseries struct {
	Values []struct {
		DateTime string `json:"datetime"`
		Close    string `json:"close"`
	} `json:"values"`
}

var timeseriesCmd = &cobra.Command{
	Use:   "timeseries",
	Short: "A brief description of your command",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		slice := []model.DataPoint{}
		godotenv.Load()
		key := os.Getenv("API_KEY")

		url := fmt.Sprintf("https://api.twelvedata.com/time_series?symbol=AAPL&interval=1min&apikey=%s", key)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			panic(err)
		}

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

		var timeseries Timeseries
		err = json.Unmarshal(body, &timeseries)
		if err != nil {
			panic(err)
		}

		for _, set := range timeseries.Values {
			slice = append(slice, model.DataPoint{DateTime: set.DateTime, Close: set.Close})
		}

		utils.Graph("Stock Prices: AAPL", slice)
	},
}

func init() {
	rootCmd.AddCommand(timeseriesCmd)
}

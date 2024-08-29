package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
)

// timeseriesCmd represents the timeseries command
var timeseriesCmd = &cobra.Command{
	Use:   "timeseries",
	Short: "A brief description of your command",
	Long:  ``,

	// Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		godotenv.Load()
		key := os.Getenv("API_KEY")
		// fundType := args[0]

		url := fmt.Sprintf("https://api.twelvedata.com/time_series?symbol=AAPL,CAD&interval=1min&apikey=%s", key)
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

		fmt.Println(string(body))

	},
}

func init() {
	rootCmd.AddCommand(timeseriesCmd)
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	// "time"
)

// timezoneCmd represents the timezone command

var timezoneCmd = &cobra.Command{
	Use:   "timezone",
	Short: "Get the current time in a given timezone",
	Long: `Get the current time in a given timezone.
               This command takes one argument, the timezone you want to get the current time in.
               It returns the current time in RFC1123 format.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		timezone := args[0]
		currentTime, err := getTimeInTimezone(timezone)
		if err != nil {
			log.Fatalln("The timezone string is invalid")
		}
		fmt.Println(currentTime)

		// timezone := args[0]
		// location, _ := time.LoadLocation(timezone)
		// dateFlag, _ := cmd.Flags().GetString("date")
		// var date string
		//
		// if dateFlag != "" {
		// 	date = time.Now().In(location).Format(dateFlag)
		// } else {
		// 	date = time.Now().In(location).Format(time.RFC3339)[:10]
		// }
		// fmt.Printf("Current date in %v: %v\n", timezone, date)
	},
}

func init() {
	rootCmd.AddCommand(timezoneCmd)
	timezoneCmd.PersistentFlags().String("date", "", "returns the date in a time zone in a specified format")
	// timezoneCmd.Flags().String("date", "", "Date for which to get the time (format: yyyy-mm-dd)")
}

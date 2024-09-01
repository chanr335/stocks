package utils

import (
	"fmt"
	"github.com/guptarohit/asciigraph"
	"log"
	"stocks_cli/cmd/model"
	"strconv"
)

func Graph(title string, dataPoints []model.DataPoint) {
	var data []float64
	// first := 0

	for _, point := range dataPoints {
		closeValue, err := strconv.ParseFloat(point.Close, 64)
		// if first == 0 {
		// 	first += 1
		// 	data = append(data, closeValue/1.02)
		// }
		if err != nil {
			log.Printf("Failed to parse close value: %v", err)
			continue
		}
		data = append(data, closeValue)
	}

	graph := asciigraph.Plot(data)
	fmt.Println(title)
	fmt.Println(graph)
}

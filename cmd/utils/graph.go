package utils

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
	"math"
	// "stocks_cli/cmd"
)

func Graph(title string, length int /* , slice []cmd.DataPoint */) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	sinData := func() [][]float64 {
		n := length
		data := make([][]float64, 2)
		data[0] = make([]float64, n)
		data[1] = make([]float64, n)
		for i := 0; i < n; i++ {
			data[0][i] = 1 + math.Sin(float64(i)/5)
			data[1][i] = 1 + math.Cos(float64(i)/5)
		}
		return data
	}()

	p0 := widgets.NewPlot()
	p0.Title = title
	p0.Data = sinData
	p0.SetRect(0, 0, 50, 15)
	p0.AxesColor = ui.ColorWhite
	p0.LineColors[0] = ui.ColorGreen

	ui.Render(p0)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

package main

import (
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/orellazri/realtime_devops/performance/http"
)

func generateHTTPItems() []opts.BarData {
	items := make([]opts.BarData, 0)

	totalTime, err := http.Benchmark(100)
	if err != nil {
		log.Fatal(err)
	}
	items = append(items, opts.BarData{Value: totalTime.Milliseconds()})

	totalTime, err = http.Benchmark(1000)
	if err != nil {
		log.Fatal(err)
	}
	items = append(items, opts.BarData{Value: totalTime.Milliseconds()})

	return items
}

func generateRedisItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	items = append(items, opts.BarData{Value: 1})
	items = append(items, opts.BarData{Value: 1})
	return items
}

func main() {
	chart := charts.NewBar()

	chart.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Message Transport Performance",
			Subtitle: "Performance benchmark of different message transport infrastructures",
			Right:    "40%",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "No. Iterations",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			AxisLabel: &opts.AxisLabel{Show: true, Formatter: "{value} ms"},
		}),
	)

	chart.SetXAxis([]string{"100", "1000"}).
		AddSeries("HTTP", generateHTTPItems()).
		AddSeries("Redis", generateRedisItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:     true,
				Position: "inside",
			}),
		)

	f, _ := os.Create("chart.html")
	chart.Render(f)
}

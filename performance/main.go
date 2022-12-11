package main

import (
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/orellazri/realtime_devops/performance/helpers"
	"github.com/orellazri/realtime_devops/performance/http"
	"github.com/orellazri/realtime_devops/performance/kafka"
	"github.com/orellazri/realtime_devops/performance/redis"
)

func generateHTTPItems() []opts.BarData {
	conn := http.NewConnection()

	return generateItems(conn)
}

func generateRedisItems() []opts.BarData {
	conn := redis.NewConnection()

	return generateItems(conn)
}

func generateKafkaItems() []opts.BarData {
	conn, err := kafka.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	return generateItems(conn)
}

func generateItems(b helpers.Benchmarkable) []opts.BarData {
	items := make([]opts.BarData, 0)

	totalTime, err := b.BenchmarkWrite(100)
	if err != nil {
		log.Fatal(err)
	}
	items = append(items, opts.BarData{Value: totalTime.Milliseconds()})

	totalTime, err = b.BenchmarkWrite(1000)
	if err != nil {
		log.Fatal(err)
	}
	items = append(items, opts.BarData{Value: totalTime.Milliseconds()})

	return items
}

func main() {
	chart := charts.NewBar()

	chart.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Message Transport Performance",
			Subtitle: "Performance benchmark of different message transport infrastructures",
			Right:    "50%",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Right: "80%"}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Iterations",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			AxisLabel: &opts.AxisLabel{Show: true, Formatter: "{value} ms"},
		}),
	)

	chart.SetXAxis([]string{"100", "1000"}).
		AddSeries("Local HTTP", generateHTTPItems()).
		AddSeries("Local Redis", generateRedisItems()).
		AddSeries("Local Kafka", generateKafkaItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      true,
				Position:  "inside",
				Formatter: "{a}\n{c}",
			}),
		)

	f, _ := os.Create("chart.html")
	chart.Render(f)
}

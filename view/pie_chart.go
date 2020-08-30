package view

import (
	"fmt"
	"math"
	"sort"

	"github.com/Ar3q/corona-info-cli/info"
	"github.com/gizak/termui/v3/widgets"
)

type PieChartCords struct {
	TopLeft, BottomRight Point
}

type Point struct {
	X, Y int
}

type ChartData struct {
	title     string
	countries []string
	fractions []float64
}

func PrepareDataForChartByCases(data info.ListCountryData, threshold int) ChartData {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Cases > data[j].Cases
	})

	countries := make([]string, threshold+1)
	fractions := make([]float64, threshold+1)

	totalCases := sumCases(data)
	fractionOfOthers := totalCases

	for i, el := range data {
		if i < threshold {
			countries[i] = el.Country
			fractions[i] = float64(el.Cases) / float64(totalCases)
			fractionOfOthers -= el.Cases
		}
	}

	countries[threshold] = "Others"
	fractions[threshold] = float64(fractionOfOthers) / float64(totalCases)

	return ChartData{title: "By Cases", countries: countries, fractions: fractions}
}

func sumCases(data info.ListCountryData) (sum int) {
	for _, el := range data {
		sum += el.Cases
	}
	return sum
}

func NewPieChart(cords PieChartCords, chartData ChartData) *widgets.PieChart {
	pc := widgets.NewPieChart()
	pc.Title = chartData.title
	pc.SetRect(cords.TopLeft.X, cords.TopLeft.Y, cords.BottomRight.X, cords.BottomRight.Y)
	pc.Data = chartData.fractions
	pc.AngleOffset = -.4 * math.Pi
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%s %.02f", chartData.countries[i], v)
	}

	return pc
}

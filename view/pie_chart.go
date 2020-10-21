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

type extractValue func(info.CountryData) int

func prepareDataForChar(data info.ListCountryData, threshold int, extractFunc extractValue) ([]string, []float64) {
	descendingSortBy(data, extractFunc)

	countries := make([]string, threshold+1)
	fractions := make([]float64, threshold+1)

	totalCases := sumValue(data, extractFunc)
	fractionOfOthers := totalCases

	for i, el := range data {
		if i < threshold {
			countries[i] = el.Country
			value := extractFunc((el))
			fractions[i] = float64(value) / float64(totalCases)
			fractionOfOthers -= value
		}
	}

	countries[threshold] = "Others"
	fractions[threshold] = float64(fractionOfOthers) / float64(totalCases)

	return countries, fractions
}

func sumValue(data info.ListCountryData, extractFunc extractValue) (sum int) {
	for _, el := range data {
		sum += extractFunc(el)
	}
	return sum
}

func descendingSortBy(data info.ListCountryData, extractFunc extractValue) {
	sort.Slice(data, func(i, j int) bool {
		return extractFunc(data[i]) > extractFunc((data[j]))
	})
}

func PrepareDataForChartByCases(data info.ListCountryData, threshold int) ChartData {
	extractCases := func(el info.CountryData) int { return el.Cases }

	countries, fractions := prepareDataForChar(data, threshold, extractCases)

	return ChartData{title: "By Cases", countries: countries, fractions: fractions}
}

func PrepareDataForChartByDeaths(data info.ListCountryData, threshold int) ChartData {
	extractDeaths := func(el info.CountryData) int { return el.Deaths }

	countries, fractions := prepareDataForChar(data, threshold, extractDeaths)

	return ChartData{title: "By Deaths", countries: countries, fractions: fractions}
}

func NewPieChart(cords PieChartCords, chartData ChartData) *widgets.PieChart {
	pc := widgets.NewPieChart()
	pc.Title = chartData.title
	pc.SetRect(cords.TopLeft.X, cords.TopLeft.Y, cords.BottomRight.X, cords.BottomRight.Y)
	pc.Data = chartData.fractions
	pc.AngleOffset = -.3 * math.Pi
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%s %.02f", chartData.countries[i], v)
	}

	return pc
}

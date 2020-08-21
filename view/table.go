package view

import (
	"strconv"

	"github.com/Ar3q/corona-info-cli/info"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var headers []string = []string{"Number", "Country", "Cases", "Deaths", "Today Cases", "Today Deaths"}

func NewCountryTable(countryStats []info.CountryData, width, height int) *widgets.Table {
	table := widgets.NewTable()
	table.Title = " Countries "

	// +1 for header
	rows := make([][]string, len(countryStats)+1)
	rows[0] = make([]string, len(headers))
	rows[0] = headers

	for i, country := range countryStats {
		incI := i + 1
		rows[incI] = make([]string, len(headers))
		rows[incI][0] = strconv.Itoa(incI)
		rows[incI][1] = country.Country
		rows[incI][2] = strconv.Itoa(country.Cases)
		rows[incI][3] = strconv.Itoa(country.Deaths)
		rows[incI][4] = strconv.Itoa(country.TodayCases)
		rows[incI][5] = strconv.Itoa(country.TodayDeaths)
	}

	table.Rows = rows
	table.TextStyle = ui.NewStyle(ui.ColorWhite)
	table.SetRect(0, 5, width, height)

	sequenceNumWidth := getWidthForColumn(0.05, width)
	nameWitdth := getWidthForColumn(0.3, width)
	numbersWidth := getWidthForColumn(0.16, width)
	table.ColumnWidths = []int{sequenceNumWidth, nameWitdth, numbersWidth, numbersWidth, numbersWidth, numbersWidth}
	return table
}

func getWidthForColumn(fraction float64, width int) int {
	return int(fraction * float64(width))
}

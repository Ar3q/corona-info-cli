package view

import (
	"math"
	"strconv"

	"github.com/Ar3q/corona-info-cli/info"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var headers []string = []string{"Number", "Country", "Cases", "Deaths", "Today Cases", "Today Deaths"}

const minYHeight = 5

// Hardcoded ratio
const ratio = 2

// NumberOfRowsInTable is maximum number of rows that can be shown in one table
var NumberOfRowsInTable int

// NewCountryTables returns slice of Tables filled with data
func NewCountryTables(countryStats []info.CountryData, width, height int) []*widgets.Table {
	NumberOfRowsInTable = (height - minYHeight - 3) / ratio
	rowsPerTable := getRowsForEachTable(countryStats, NumberOfRowsInTable)

	tables := make([]*widgets.Table, len(rowsPerTable))
	for i, rows := range rowsPerTable {
		tables[i] = newCountryTable(rows, width, height)
	}

	return tables
}

func getRowsForEachTable(countryStats []info.CountryData, maxNumberOfRowsInTable int) [][][]string {
	numberOfTables := int(math.Ceil(float64(len(countryStats)) / float64(maxNumberOfRowsInTable)))
	rowsPerTable := make([][][]string, numberOfTables)

	for i := 0; i < numberOfTables; i++ {
		// +1 for header
		rows := make([][]string, maxNumberOfRowsInTable+1)
		rows[0] = make([]string, len(headers))
		rows[0] = headers

		for j, country := range countryStats[i*maxNumberOfRowsInTable : (i+1)*maxNumberOfRowsInTable] {
			incJ := j + 1
			orderNumber := i*maxNumberOfRowsInTable + incJ
			rows[incJ] = make([]string, len(headers))
			rows[incJ][0] = strconv.Itoa(orderNumber)
			rows[incJ][1] = country.Country
			rows[incJ][2] = strconv.Itoa(country.Cases)
			rows[incJ][3] = strconv.Itoa(country.Deaths)
			rows[incJ][4] = strconv.Itoa(country.TodayCases)
			rows[incJ][5] = strconv.Itoa(country.TodayDeaths)
		}

		rowsPerTable[i] = rows
	}

	return rowsPerTable

}

func newCountryTable(rows [][]string, width, height int) *widgets.Table {
	table := widgets.NewTable()
	table.Title = " Countries "

	table.Rows = rows
	table.TextStyle = ui.NewStyle(ui.ColorWhite)
	table.SetRect(0, minYHeight, width, height)

	sequenceNumWidth := getWidthForColumn(0.05, width)
	nameWitdth := getWidthForColumn(0.3, width)
	numbersWidth := getWidthForColumn(0.16, width)
	table.ColumnWidths = []int{sequenceNumWidth, nameWitdth, numbersWidth, numbersWidth, numbersWidth, numbersWidth}
	return table
}

func getWidthForColumn(fraction float64, width int) int {
	return int(fraction * float64(width))
}

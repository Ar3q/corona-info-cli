package view

import (
	"os"
	"strconv"

	"github.com/Ar3q/corona-info-cli/info"
	"github.com/olekukonko/tablewriter"
)

// PrintTable prints table to standard output.
// Uses passed data given from API and limits number of rows to numberOfCountries
func PrintTable(responseObject *info.Response, numberOfCountries int) {
	var data [][]string
	for i, countryObject := range responseObject.Data {
		if i == numberOfCountries {
			break
		}
		cases := strconv.Itoa(countryObject.Cases)
		todayCases := strconv.Itoa(countryObject.TodayCases)
		deaths := strconv.Itoa(countryObject.Deaths)
		todayDeaths := strconv.Itoa(countryObject.TodayDeaths)
		countryStringified := []string{strconv.Itoa(i + 1), countryObject.Country, cases, deaths, todayCases, todayDeaths}
		data = append(data, countryStringified)
	}

	tableHeaders := []string{"Number", "Country", "Cases", "Deaths", "Today Cases", "Today Deaths"}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(tableHeaders)
	table.AppendBulk(data)
	table.Render()
}

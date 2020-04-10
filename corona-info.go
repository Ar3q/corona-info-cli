package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type response struct {
	Data []countryData `json:"data"`
}

type countryData struct {
	Country     string `json:"country"`
	Cases       int    `json:"cases"`
	TodayCases  int    `json:"todayCases"`
	Deaths      int    `json:"deaths"`
	TodayDeaths int    `json:"todayDeaths"`
	Recovered   int    `json:"recovered"`
}

func getURL(country *string) string {
	var urlBuilder strings.Builder

	urlBuilder.WriteString("https://corona-stats.online/")

	if *country != "" {
		urlBuilder.WriteString(*country)
	}

	urlBuilder.WriteString("?format=json")

	return urlBuilder.String()
}

func printTable(responseObject *response, numberOfCountries int) {
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

func main() {
	country := flag.String("c", "", "country name")
	top := flag.Int("t", 255, "top x countries")
	flag.Parse()

	url := getURL(country)

	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var responseObject response
	err = json.Unmarshal(jsonData, &responseObject)
	if err != nil {
		log.Fatalln(err)
	}

	printTable(&responseObject, *top)
}

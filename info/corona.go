package info

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Data ListCountryData `json:"data"`
}

type CountryData struct {
	Country     string `json:"country"`
	Cases       int    `json:"cases"`
	TodayCases  int    `json:"todayCases"`
	Deaths      int    `json:"deaths"`
	TodayDeaths int    `json:"todayDeaths"`
	Recovered   int    `json:"recovered"`
}

type ListCountryData []CountryData

const urlForAllCountries string = "https://corona-stats.online/?format=json"

// Fetch returns response which contains info about all countries
func Fetch() (*Response, error) {
	return fetchInfoFromAPI(urlForAllCountries)
}

func getURL(country string) string {
	var urlBuilder strings.Builder

	urlBuilder.WriteString("https://corona-stats.online/")

	if country != "" {
		urlBuilder.WriteString(country)
	}

	urlBuilder.WriteString("?format=json")

	return urlBuilder.String()
}

// FetchForOneCountry returns response which contains info about specified country
func FetchForOneCountry(country string) (*Response, error) {
	url := getURL(country)
	return fetchInfoFromAPI(url)
}

func fetchInfoFromAPI(url string) (*Response, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer res.Body.Close()

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var responseObject Response
	err = json.Unmarshal(jsonData, &responseObject)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &responseObject, nil
}

func (data ListCountryData) FilterByCountry(match string) ListCountryData {
	lowerMatch := strings.ToLower(match)

	var filtered ListCountryData
	for _, countryData := range data {
		lowerCountry := strings.ToLower(countryData.Country)
		if strings.Contains(lowerCountry, lowerMatch) {
			filtered = append(filtered, countryData)
		}
	}

	return filtered
}

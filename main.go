package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Ar3q/corona-info-cli/info"
	"github.com/Ar3q/corona-info-cli/view"
)

func main() {
	country := flag.String("c", "", "Country name")
	top := flag.Int("t", 255, "Top x countries")
	help := flag.Bool("h", false, "Show help")
	flag.Parse()

	if *help {
		fmt.Println("Help:")
		fmt.Println("-c COUNTRY\tReturns table for given COUNTRY")
		fmt.Println("-t NUMBER\tReturns table for first NUMBER countries")
		os.Exit(0)
	}

	var data *info.Response
	var err error
	if *country == "" {
		data, err = info.Fetch()
	} else {
		data, err = info.FetchForOneCountry(*country)
	}
	if err != nil {
		fmt.Printf("Error: %v appeared", err)
		os.Exit(1)
	}

	view.PrintTable(data, *top)
}

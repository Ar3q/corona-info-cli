package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Ar3q/corona-info-cli/info"
	"github.com/Ar3q/corona-info-cli/view"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	country := flag.String("c", "", "Country name")
	help := flag.Bool("h", false, "Show help")
	flag.Parse()

	if *help {
		fmt.Println("Help:")
		fmt.Println("-c COUNTRY\tReturns table for given COUNTRY")
		os.Exit(0)
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

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

	termWidth, termHeight := ui.TerminalDimensions()
	// fmt.Printf("H: %d, W: %d\n", termHeight, termWidth)

	tablesOfCountries := view.NewCountryTables(data.Data, termWidth, termHeight)

	panes := make([]string, len(tablesOfCountries))
	for i := 1; i <= len(tablesOfCountries); i++ {
		panes[i-1] = strconv.Itoa(i)
	}

	tabpane := widgets.NewTabPane(panes...)
	tabpane.SetRect(0, 2, termWidth, 5)
	tabpane.Border = true

	header := widgets.NewParagraph()
	header.Text = "Press q to quit, Press h or l to switch tabs"
	header.SetRect(0, 1, termWidth, 2)
	header.Border = false

	renderTab := func() {
		ui.Clear()
		ui.Render(header, tabpane)
		ui.Render(tablesOfCountries[tabpane.ActiveTabIndex])
	}

	ui.Render(header, tabpane, tablesOfCountries[0])

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "h":
			tabpane.FocusLeft()
			renderTab()
		case "l":
			tabpane.FocusRight()
			renderTab()
		}
	}
}

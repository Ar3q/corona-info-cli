package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Ar3q/corona-info-cli/view"

	"github.com/Ar3q/corona-info-cli/info"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	country := flag.String("c", "", "Country name")
	// top := flag.Int("t", 255, "Top x countries")
	help := flag.Bool("h", false, "Show help")
	flag.Parse()

	if *help {
		fmt.Println("Help:")
		fmt.Println("-c COUNTRY\tReturns table for given COUNTRY")
		fmt.Println("-t NUMBER\tReturns table for first NUMBER countries")
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
	tabpane := widgets.NewTabPane("1", "2")
	tabpane.SetRect(0, 2, termWidth, 5)
	tabpane.Border = true

	// view.PrintTable(data, *top)
	tableOfCountries := view.NewCountryTable(data.Data, termWidth, termHeight)

	header := widgets.NewParagraph()
	header.Text = "Press q to quit, Press h or l to switch tabs"
	header.SetRect(0, 1, termWidth, 2)
	header.Border = false
	header.TextStyle.Bg = ui.ColorBlue

	renderTab := func() {
		switch tabpane.ActiveTabIndex {
		case 0:
			ui.Render(tableOfCountries)
		case 1:
			ui.Render(header)
		}
	}

	ui.Render(tabpane)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "h":
			tabpane.FocusLeft()
			ui.Clear()
			ui.Render(tabpane)
			renderTab()
		case "l":
			tabpane.FocusRight()
			ui.Clear()
			ui.Render(tabpane)
			renderTab()
		}
	}
}

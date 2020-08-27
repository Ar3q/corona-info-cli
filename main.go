package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Ar3q/corona-info-cli/info"
	"github.com/Ar3q/corona-info-cli/view"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var find string
var searching bool = false
var found bool = false

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

	findParagraph := widgets.NewParagraph()
	findParagraph.Text = ""
	findParagraph.SetRect(0, termHeight-3, termWidth, termHeight)

	renderTab := func(showFinder bool) {
		ui.Clear()
		ui.Render(header, tabpane)
		ui.Render(tablesOfCountries[tabpane.ActiveTabIndex])
		if showFinder {
			findParagraph.Text = find
			ui.Render(findParagraph)
		}
	}

	ui.Render(header, tabpane, tablesOfCountries[0])

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		if searching {
			switch e.ID {
			case "<C-c>":
				return
			case "/", "<Escape>":
				renderTab(false)
				searching = false
			case "<C-u>":
				find = ""
				renderTab(true)
			case "<C-l>":
				if found {
					find = ""
					renderTab(true)
				}
			case "<Backspace>":
				if length := len(find); length > 0 {
					lastCharacter := find[length-1]
					find = strings.TrimSuffix(find, string(lastCharacter))
					renderTab(true)
				}
			case "<Enter>":
				filteredData := data.Data.FilterByCountry(find)
				tablesOfCountries = view.NewCountryTables(filteredData, termWidth, termHeight)
				find = "Results for: " + find
				renderTab(true)
				found = true
			default:
				if !found {
					find = find + e.ID
					renderTab(true)
				}
			}

		} else {
			switch e.ID {
			case "q", "<C-c>":
				return
			case "h":
				tabpane.FocusLeft()
				renderTab(false)
			case "l":
				tabpane.FocusRight()
				renderTab(false)
			case "/":
				find = ""
				renderTab(true)
				searching = true
				found = false
			}

		}
	}
}

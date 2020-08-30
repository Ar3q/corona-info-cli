package main

import (
	"fmt"
	"log"
	"os"
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
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	var data *info.Response
	var err error
	data, err = info.Fetch()
	if err != nil {
		fmt.Printf("Error: %v appeared", err)
		os.Exit(1)
	}

	termWidth, termHeight := ui.TerminalDimensions()

	tablesOfCountries := view.NewCountryTables(data.Data, termWidth/2, termHeight)

	tabpane := view.NewTabPane(termWidth/2, len(tablesOfCountries))
	helper := view.NewHelper(termWidth, "")

	findParagraph := widgets.NewParagraph()
	findParagraph.Text = ""
	findParagraph.SetRect(0, termHeight-3, termWidth/2, termHeight)

	pieChartData := view.PrepareDataForChartByCases(data.Data, 5)
	pc1 := view.NewPieChart(view.PieChartCords{view.Point{termWidth/2 + 1, 5}, view.Point{termWidth, termHeight}}, pieChartData)

	renderTab := func(showFinder bool) {
		ui.Clear()
		ui.Render(helper, tabpane, pc1)
		ui.Render(tablesOfCountries[tabpane.ActiveTabIndex])
		if showFinder {
			findParagraph.Text = find
			ui.Render(findParagraph)
		}
	}

	refreshTablesAndTabpane := func() {
		filteredData := data.Data.FilterByCountry(find)
		tablesOfCountries = view.NewCountryTables(filteredData, termWidth/2, termHeight)

		tabpane = view.NewTabPane(termWidth/2, len(tablesOfCountries))
	}

	ui.Render(helper, tabpane, tablesOfCountries[0], pc1)

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
				find = ""
				refreshTablesAndTabpane()
				renderTab(false)
			case "<Backspace>":
				if length := len(find); length > 0 {
					lastCharacter := find[length-1]
					find = strings.TrimSuffix(find, string(lastCharacter))
					renderTab(true)
				}
			case "<Enter>":
				if !found {
					refreshTablesAndTabpane()
					find = "Results for: " + find
					renderTab(true)
					found = true
				}
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
			case "<C-l>":
				find = ""
				refreshTablesAndTabpane()
				renderTab(false)
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

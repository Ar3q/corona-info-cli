package view

import (
	"strconv"

	"github.com/gizak/termui/v3/widgets"
)

func NewTabPane(tabPaneWidth, numberOfPanes int) *widgets.TabPane {
	panes := make([]string, numberOfPanes)
	for i := 1; i <= numberOfPanes; i++ {
		panes[i-1] = strconv.Itoa(i)
	}

	tabpane := widgets.NewTabPane(panes...)
	tabpane.SetRect(0, 2, tabPaneWidth, 5)
	tabpane.Border = true

	return tabpane
}

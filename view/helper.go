package view

import "github.com/gizak/termui/v3/widgets"

func NewHelper(width int, differentText string) *widgets.Paragraph {
	helper := widgets.NewParagraph()

	if differentText == "" {
		helper.Text = "Press q to quit, Press h or l to switch tabs"
	} else {
		helper.Text = differentText
	}
	helper.SetRect(0, 1, width, 2)
	helper.Border = false

	return helper
}
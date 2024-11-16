package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type SimpleListItem struct {
	widget.BaseWidget
	LblText *widget.Label
}

func NewSimpleListItem(text string) *SimpleListItem {
	i := &SimpleListItem{
		LblText: widget.NewLabel(text),
	}

	i.ExtendBaseWidget(i)

	return i
}

func (si *SimpleListItem) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(si.LblText)
}

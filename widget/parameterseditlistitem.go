package widget

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/halsten-dev/komok/data"
	kManager "github.com/halsten-dev/komok/manager"
	"reflect"
)

// ParametersEditListItem widget definition
type ParametersEditListItem struct {
	widget.BaseWidget
	LblParam   *widget.Label
	EntryParam *Entry
	CbParam    *widget.Check
}

// NewParametersEditListItem creates a new widget
func NewParametersEditListItem(sm *kManager.ShortcutsManager) *ParametersEditListItem {
	item := &ParametersEditListItem{
		LblParam:   widget.NewLabel(""),
		EntryParam: NewEntry(sm),
		CbParam:    widget.NewCheck("", nil),
	}

	item.ExtendBaseWidget(item)

	return item
}

// CreateRenderer return the way the widget should be rendered
func (peli *ParametersEditListItem) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewVBox(peli.LblParam, peli.EntryParam)

	return widget.NewSimpleRenderer(c)
}

// UpdateData updates the widget's fields based on the given data.Param
func (peli *ParametersEditListItem) UpdateData(param *data.Param) {
	if param != nil {
		peli.LblParam.Text = param.Label + " (" + fmt.Sprintf("%v", param.Type) + ")"

		// Manage entry visibility
		switch param.Type {
		case reflect.Bool:
			peli.CbParam.Show()
			peli.EntryParam.Hide()
			peli.CbParam.Checked = param.Value.(bool)
		default:
			peli.CbParam.Hide()
			peli.EntryParam.Show()
			peli.EntryParam.Text = fmt.Sprintf("%v", param.Value)
		}

	} else {
		peli.LblParam.Text = ""
		peli.EntryParam.Text = ""
	}

	peli.Refresh()
}

package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/halsten-dev/komok/data"
	kManager "github.com/halsten-dev/komok/manager"
)

type ParamEditListItem struct {
	widget.BaseWidget

	LblParam   *widget.Label
	EntryParam *Entry
}

func NewParamEditListItem(sm *kManager.ShortcutsManager) *ParamEditListItem {
	item := &ParamEditListItem{
		LblParam:   widget.NewLabel(""),
		EntryParam: NewEntry(sm),
	}

	item.ExtendBaseWidget(item)

	return item
}

func (p *ParamEditListItem) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewVBox(p.LblParam, p.EntryParam)

	return widget.NewSimpleRenderer(c)
}

func (p *ParamEditListItem) UpdateData(param *data.ParamString) {
	p.LblParam.Text = param.Label
	p.EntryParam.Text = param.Value

	p.Refresh()
}

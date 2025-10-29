package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/halsten-dev/komok/data"
	kManager "github.com/halsten-dev/komok/manager"
)

type ParamEditList struct {
	widget.BaseWidget
	LstParams     *widget.List
	Data          data.ParamStringListData
	ItemOnChanged func()
	OnSave        func()
}

func NewParamEditList(sm *kManager.ShortcutsManager) *ParamEditList {
	item := &ParamEditList{
		Data: data.ParamStringListData{},
	}

	item.LstParams = widget.NewList(
		func() int {
			if item.Data != nil {
				return len(item.Data)
			}

			return 0
		},
		func() fyne.CanvasObject {
			return NewParamEditListItem(sm)
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			rowItem := object.(*ParamEditListItem)

			if item.Data != nil {
				rowItem.UpdateData(item.Data[id])
			} else {
				rowItem.UpdateData(nil)
			}

			rowItem.EntryParam.OnChanged = func(s string) {
				item.Data[id].Value = s

				item.ItemOnChanged()
			}
		})

	item.ExtendBaseWidget(item)

	return item
}

// CreateRenderer returns the way the widget should be rendered
func (p *ParamEditList) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(p.LstParams)
}

func (p *ParamEditList) UpdateData(data data.ParamStringListData) {
	p.Data = data
	p.LstParams.Refresh()
}

func (p *ParamEditList) Save() {
	if len(p.Data) == 0 {
		return
	}

	p.OnSave()
}

func (p *ParamEditList) Show() {
	if len(p.Data) == 0 {
		return
	}

	p.BaseWidget.Show()
}

func (p *ParamEditList) Reset() {
	p.UpdateData(nil)
	p.Hide()
}

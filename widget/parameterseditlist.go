package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/halsten-dev/komok/data"
	kManager "github.com/halsten-dev/komok/manager"
	"log"
	"reflect"
	"strconv"
)

// ParametersEditList widget definition
type ParametersEditList struct {
	widget.BaseWidget
	LstParams     *widget.List
	Data          *data.ParamListData
	ItemOnChanged func()
	OnSave        func()
}

// NewParametersEditList creates a new widget
func NewParametersEditList(sm *kManager.ShortcutsManager) *ParametersEditList {
	item := &ParametersEditList{
		Data: &data.ParamListData{},
	}

	item.LstParams = widget.NewList(
		func() int {
			if item.Data != nil {
				if item.Data.Params != nil {
					return len(item.Data.Params)
				}
			}

			return 0
		},
		func() fyne.CanvasObject {
			return NewParametersEditListItem(sm)
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			rowItem := object.(*ParametersEditListItem)

			if item.Data != nil {
				rowItem.UpdateData(&item.Data.Params[id])
			} else {
				rowItem.UpdateData(nil)
			}

			rowItem.EntryParam.OnChanged = func(s string) {
				switch item.Data.Params[id].Type {
				case reflect.String:
					item.Data.Params[id].Value = s
				case reflect.Int:
					item.Data.Params[id].Value, _ = strconv.Atoi(s)
				case reflect.Float64:
					item.Data.Params[id].Value, _ = strconv.ParseFloat(s, 64)
				case reflect.Uint, reflect.Uint8:
					item.Data.Params[id].Value, _ = strconv.ParseUint(s, 10, 64)
				default:
					item.Data.Params[id].Value = nil
				}

				item.ItemOnChanged()
			}

			rowItem.CbParam.OnChanged = func(checked bool) {
				switch item.Data.Params[id].Type {
				case reflect.Bool:
					item.Data.Params[id].Value = checked
				}

				item.ItemOnChanged()
			}
		})

	item.ExtendBaseWidget(item)
	return item
}

// CreateRenderer returns the way the widget should be rendered
func (pel *ParametersEditList) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(pel.LstParams)
}

// UpdateData update the data and refresh the list
func (pel *ParametersEditList) UpdateData(data *data.ParamListData) {
	pel.Data = data

	pel.LstParams.Refresh()
}

// Save updates the data and call the save callback
func (pel *ParametersEditList) Save() {
	if pel.Data == nil {
		return
	}

	err := pel.Data.Save()

	if err != nil {
		log.Println(err)
		return
	}

	pel.OnSave()
}

// Show is overridden from the base widget to only show the form if there is data.
func (pel *ParametersEditList) Show() {
	if pel.Data == nil {
		return
	}

	pel.BaseWidget.Show()
}

func (pel *ParametersEditList) Reset() {
	pel.UpdateData(nil)
	pel.Hide()
}

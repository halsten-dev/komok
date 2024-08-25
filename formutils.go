package komok

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"reflect"
)

func NewFormFromStruct(sm *ShortcutsManager, s interface{}) *widget.Form {
	var items []*widget.FormItem
	var formEntry fyne.Widget

	items = make([]*widget.FormItem, 0)

	t := reflect.TypeOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Name
		fieldType := field.Type.Kind()

		switch fieldType {
		case reflect.String:
			formEntry = NewEntry(sm)
		case reflect.Int:
			formEntry = NewNumericalEntry(sm, true)
		case reflect.Float64:
			formEntry = NewNumericalEntry(sm, false)
		default:
			fmt.Printf("Unsupported field type: %s\n", fieldType)
		}

		item := widget.NewFormItem(fieldName, formEntry)
		items = append(items, item)
	}

	form := widget.NewForm(items...)

	return form
}

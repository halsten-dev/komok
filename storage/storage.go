package storage

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

var appDataLister fyne.ListableURI

func Init(app fyne.App) error {
	var err error

	appDataLister, err = storage.ListerForURI(app.Storage().RootURI())

	if err != nil {
		return err
	}

	return nil
}

func GetPath() string {
	return appDataLister.Path()
}

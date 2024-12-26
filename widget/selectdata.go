package widget

import (
	"fyne.io/fyne/v2/widget"
	"sort"
)

type SelectData[T any] struct {
	widget.Select

	data map[string]T
}

// NewSelectData creates a new select without any data attached.
func NewSelectData[T any]() *SelectData[T] {
	s := &SelectData[T]{}

	s.data = make(map[string]T)

	s.ExtendBaseWidget(s)

	return s
}

// NewSelectDataWithData creates a new select with data.
func NewSelectDataWithData[T any](data map[string]T) *SelectData[T] {
	s := &SelectData[T]{}

	s.data = data

	s.Options = s.generateOptions()

	s.ExtendBaseWidget(s)

	return s
}

// SetData sets data and init options.
func (s *SelectData[T]) SetData(data map[string]T) {
	s.data = data
	s.Options = s.generateOptions()
}

// GetSelectedData returns the data linked with the selected option.
func (s *SelectData[T]) GetSelectedData() T {
	return s.data[s.Select.Selected]
}

// generateOptions generates a []string based on the linked data.
func (s *SelectData[T]) generateOptions() []string {
	var options []string

	for key := range s.data {
		options = append(options, key)
	}

	sort.Strings(options)

	return options
}

package widget

import (
	"sort"

	"fyne.io/fyne/v2/widget"
)

type SelectData[T any] struct {
	widget.Select

	data     map[string]T
	keyOrder []string
}

// NewSelectData creates a new select without any data attached.
func NewSelectData[T any]() *SelectData[T] {
	s := &SelectData[T]{}

	s.data = make(map[string]T)
	s.keyOrder = make([]string, 0)

	s.ExtendBaseWidget(s)

	return s
}

// NewSelectDataWithData creates a new select with data.
func NewSelectDataWithData[T any](data map[string]T) *SelectData[T] {
	s := &SelectData[T]{}

	s.data = data
	s.keyOrder = make([]string, 0)

	s.Options = s.generateOptions()

	s.ExtendBaseWidget(s)

	return s
}

// NewSelectDataWithOrderedData creates a new select with data and key order slice.
func NewSelectDataWithOrderedData[T any](data map[string]T, keyOrder []string) *SelectData[T] {
	s := &SelectData[T]{}

	s.data = data
	s.keyOrder = keyOrder

	s.Options = s.generateOptions()

	s.ExtendBaseWidget(s)

	return s
}

// SetData sets data and init options.
func (s *SelectData[T]) SetData(data map[string]T) {
	s.data = data
	s.keyOrder = make([]string, 0)
	s.Options = s.generateOptions()
}

// SetOrderedData sets data and the keyOrder and init options.
func (s *SelectData[T]) SetOrderedData(data map[string]T, keyOrder []string) {
	s.data = data
	s.keyOrder = keyOrder
	s.Options = s.generateOptions()
}

// GetSelectedData returns the data linked with the selected option.
func (s *SelectData[T]) GetSelectedData() T {
	return s.data[s.Select.Selected]
}

// generateOptions generates a []string based on the linked data.
func (s *SelectData[T]) generateOptions() []string {
	var options []string

	if len(s.keyOrder) > 0 {
		for _, key := range s.keyOrder {
			options = append(options, key)
		}
	} else {
		for key := range s.data {
			options = append(options, key)
		}

		sort.Strings(options)
	}

	return options
}

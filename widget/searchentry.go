package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	kManager "github.com/halsten-dev/komok/manager"
)

// SearchEntryItem interface represents an individual item in the search list.
type SearchEntryItem[T any] interface {
	// CanvasObject Inherits from fyne.CanvasObject interface.
	fyne.CanvasObject

	// UpdateData should be used to init the item widget content.
	UpdateData(T)
}

// SearchEntryItemGenerator represents a struct capable of generating new SearchEntry list item.
type SearchEntryItemGenerator[T any] interface {
	// NewItem create a new instance of SearchEntryItem.
	NewItem() SearchEntryItem[T]
}

// SearchEntry is a widget that is bound to data and allow for searching in it and selecting a value.
type SearchEntry[T any] struct {
	// Entry represents the base of the SearchEntry.
	Entry

	// LstSearch is the list to show the results of the search.
	LstSearch *searchList[T]

	// PopSearch is the popup dialog holding the result list.
	PopSearch *widget.PopUp

	// OnSearch must be defined and should give the SearchEntry widget a new set of results by using SetSearchResult.
	OnSearch func(s string)

	// OnChangedCustom allows to set custom behaviour in the OnChanged event.
	OnChangedCustom func(s string)

	// OnSelectedKeyCustom allows to set custom behaviour when a key is selected.
	OnSelectedKeyCustom func(s string)

	// GetDataValue must be defined to return a calculated string value of the given key.
	GetDataValue func(key string) string

	// data represents the whole data.
	data map[string]T

	// searchResult represents the slice of found keys for the search value.
	searchResult []string

	// keyOrder can be defined to keep the needed order in result list.
	// TODO: Implement feature
	keyOrder []string

	// selectedKey represents the currently selected key.
	selectedKey string

	// CanBeEmpty is true by default and authorize or not to select an empty value.
	CanBeEmpty bool

	// generator represents the item generator.
	generator SearchEntryItemGenerator[T]

	// canvas represents the canvas the object was added to.
	canvas fyne.Canvas
}

func NewSearchEntry[T any](sm *kManager.ShortcutsManager, ig SearchEntryItemGenerator[T],
	canvas fyne.Canvas) *SearchEntry[T] {
	s := &SearchEntry[T]{
		generator: ig,
	}

	s.CanBeEmpty = true

	s.Entry.ShortcutsManager = sm
	s.Entry.MultiLine = false
	s.Entry.TabManagement = false
	s.Entry.SelectAllOnFocus = true

	s.data = make(map[string]T)
	s.keyOrder = make([]string, 0)
	s.searchResult = make([]string, 0)
	s.selectedKey = ""

	s.canvas = canvas

	s.LstSearch = newSearchList[T](s)

	s.OnChanged = func(str string) {
		if s.OnChangedCustom != nil {
			s.OnChangedCustom(str)
		}

		if s.OnSearch != nil {
			s.OnSearch(str)
		}

		if len(s.searchResult) > 0 {
			s.ShowSearchResults()
		} else {
			s.HideSearchResults()
		}

		s.LstSearch.Refresh()
	}

	s.PopSearch = widget.NewPopUp(s.LstSearch, s.canvas)

	s.ExtendBaseWidget(s)

	return s
}

func (se *SearchEntry[T]) ShowSearchResults() {
	se.LstSearch.UnselectAll()
	se.PopSearch.Resize(se.searchResultsMaxSize())
	se.PopSearch.ShowAtPosition(se.searchResultPos())
	se.canvas.Focus(se.LstSearch)
	se.LstSearch.selectElement(0)
}

func (se *SearchEntry[T]) HideSearchResults() {
	se.PopSearch.Hide()
}

func (se *SearchEntry[T]) searchResultsMaxSize() fyne.Size {
	count := len(se.searchResult)

	if count > 10 {
		count = 10
	}

	height := (se.LstSearch.MinSize().Height + theme.Padding()) * float32(count)
	height += theme.Padding()

	maxHeight := se.canvas.Size().Height - se.Position().Y - se.Size().Height - 2*theme.Padding()

	if height > maxHeight {
		height = maxHeight
	}

	return fyne.NewSize(se.Size().Width, height)
}

func (se *SearchEntry[T]) searchResultPos() fyne.Position {
	entryPos := fyne.CurrentApp().Driver().AbsolutePositionForObject(se)
	return entryPos.Add(fyne.NewPos(0, se.Size().Height))
}

func (se *SearchEntry[T]) Move(pos fyne.Position) {
	se.Entry.Move(pos)

	se.PopSearch.Resize(se.searchResultsMaxSize())
	se.PopSearch.Move(se.searchResultPos())
}

// FocusLost is a hook called by the focus handling logic after this object lost the focus.
func (se *SearchEntry[T]) FocusLost() {
	se.Entry.FocusLost()
	se.HideSearchResults()
	se.SetSelectedKey(se.selectedKey)
}

// TypedRune is a hook called by the input handling logic on text input events if this object is focused.
func (se *SearchEntry[T]) TypedRune(r rune) {
	se.Entry.TypedRune(r)
}

// TypedKey is a hook called by the input handling logic on key events if this object is focused.
func (se *SearchEntry[T]) TypedKey(ke *fyne.KeyEvent) {
	switch ke.Name {
	case fyne.KeyEscape:
		se.HideSearchResults()
		se.SetSelectedKey(se.selectedKey)
	default:
		se.Entry.TypedKey(ke)
	}
}

func (se *SearchEntry[T]) SetSelectedKey(key string) {
	var value string

	if key == "" {
		se.selectedKey = ""
		se.Entry.Text = ""

		se.Entry.Refresh() // Need to refresh before the workaround. If not, the old value will be reapplied to the entry. Thanks Fyne...
		se.Entry.TypedKey(&fyne.KeyEvent{
			Name: fyne.KeyLeft,
		}) // workaround to remove the selection
	} else {
		value = se.GetDataValue(key)

		if value == "" {
			value = key
		}

		se.selectedKey = key
		se.Entry.Text = value
	}

	se.Refresh()
	se.HideSearchResults()

	if se.OnSelectedKeyCustom != nil {
		se.OnSelectedKeyCustom(se.selectedKey)
	}
}

func (se *SearchEntry[T]) SetData(data map[string]T) {
	se.data = data
}

func (se *SearchEntry[T]) SetKeyOrder(keys []string) {
	se.keyOrder = keys
}

func (se *SearchEntry[T]) SetSearchResult(searchResult []string) {
	if se.CanBeEmpty && len(searchResult) > 0 {
		searchResult = append([]string{""}, searchResult...)
	}

	se.searchResult = searchResult
}

func (se *SearchEntry[T]) GetSelectedData() T {
	data, _ := se.data[se.selectedKey]

	return data
}

type searchList[T any] struct {
	widget.List

	entry *SearchEntry[T]

	selectedIndex int
}

func newSearchList[T any](entry *SearchEntry[T]) *searchList[T] {
	s := &searchList[T]{entry: entry}

	s.selectedIndex = -1

	s.List = widget.List{
		Length: func() int {
			return len(s.entry.searchResult)
		},
		CreateItem: func() fyne.CanvasObject {
			return s.entry.generator.NewItem()
		},
		UpdateItem: func(id widget.ListItemID, item fyne.CanvasObject) {
			rowItem := item.(SearchEntryItem[T])

			rowItem.UpdateData(s.entry.data[s.entry.searchResult[id]])
		},
		OnSelected: func(id widget.ListItemID) {
			s.selectedIndex = id
		},
		OnUnselected: func(id widget.ListItemID) {
			s.selectedIndex = -1
		},
	}

	s.ExtendBaseWidget(s)

	return s
}

func (sl *searchList[T]) TypedKey(ke *fyne.KeyEvent) {
	switch ke.Name {
	case fyne.KeyDown:
		if sl.selectedIndex < len(sl.entry.searchResult)-1 {
			sl.selectedIndex++
		} else {
			sl.selectedIndex = 0
		}

		sl.selectElement(sl.selectedIndex)

	case fyne.KeyUp:
		if sl.selectedIndex > 0 {
			sl.selectedIndex--
		} else {
			sl.selectedIndex = len(sl.entry.searchResult) - 1
		}

		sl.selectElement(sl.selectedIndex)

	case fyne.KeyEnter, fyne.KeyReturn:
		if sl.selectedIndex > -1 {
			sl.entry.SetSelectedKey(sl.entry.searchResult[sl.selectedIndex])
		}
	default:
		sl.entry.TypedKey(ke)
	}
}

func (sl *searchList[T]) TypedRune(r rune) {
	sl.entry.TypedRune(r)
}

func (sl *searchList[T]) selectElement(index int) {
	sl.Select(index)
	sl.Refresh()
}

func (sl *searchList[T]) FocusGained() {
}

func (sl *searchList[T]) FocusLost() {}

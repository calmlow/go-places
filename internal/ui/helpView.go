package ui

import (
	"fmt"

	"github.com/calmlow/go-places/internal/io"
	"github.com/calmlow/go-places/internal/types"
	"github.com/rivo/tview"
)

func GetHelpTextView(place types.Place) *tview.TextView {

	var textContent string

	contents, err := io.GetReadmeFileContents(place.Path)
	if err != nil {
		textContent = "Cannot read readme.md file. " + err.Error() + "\n"
	} else {
		textContent = string(contents)
	}

	textView := tview.NewTextView().
		SetDynamicColors(false).
		SetRegions(true)

	textView.SetBorder(true)
	textView.Clear()
	textView.SetTitle(fmt.Sprintf("Readme.md for %s (%s)", place.Name, place.Path))
	textView.SetText(textContent)

	return textView
}

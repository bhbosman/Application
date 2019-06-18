/*
A presentation of the tview package, implemented with tview.

Navigation

The presentation will advance to the next slide when the primitive demonstrated
in the current slide is left (usually by hitting Enter or Escape). Additionally,
the following shortcuts can be used:

  - Ctrl-N: Jump to next slide
  - Ctrl-P: Jump to previous slide
*/
package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// Starting point for the presentation.
func main() {
	var app = tview.NewApplication()
	screen := func() (content tview.Primitive) {

		list := tview.NewList().ShowSecondaryText(false).
			AddItem("Start feed", "", '1', nil).
			AddItem("Stop Feed", "", '2', nil)

		logWriter := tview.NewTextView().
			SetChangedFunc(func() {
				app.Draw()
			}).
			SetScrollable(false)
		return tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(
				tview.NewBox().
					SetDrawFunc(
						func(screen tcell.Screen, x, y, width, height int) (i int, i2 int, i3 int, i4 int) {
							tview.Print(screen, "fdsfsdddd", 1, 1, 100, tview.AlignLeft, tcell.ColorGreen)
							return x, y, width, height

						}), 5, 0, false).
			AddItem(list, 2, 0, true).
			AddItem(logWriter, 0, 1, false)

	}

	// The bottom row has some info on where we are.

	// Create the main layout.
	layout := screen()

	// Shortcuts to navigate the slides.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		return event
	})

	// Start the application.
	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}

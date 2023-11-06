// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"log"
	"os"

	"gioui.org/app"             // app contains Window handling.
	"gioui.org/io/system"       // system is used for system events (e.g. closing the window).
	"gioui.org/layout"          // layout is used for layouting widgets.
	"gioui.org/op"              // op is used for recording different operations.
	"gioui.org/unit"            // unit is used to define pixel-independent sizes
	"gioui.org/widget/material" // material contains material design widgets.
)

func main() {
	// The ui loop is separated from the application window creation
	// such that it can be used for testing.
	ui := NewUI()

	go func() {
		w := app.NewWindow(
			// Set the window title.
			app.Title("Hello, Gio!"),
			// Set the size for the window.
			app.Size(unit.Dp(800), unit.Dp(400)),
		)
		if err := ui.Run(w); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}()

	app.Main()
}

// UI holds all of the application state.
type UI struct {
	// Theme is used to hold the fonts used throughout the application.
	Theme *material.Theme
}

// NewUI creates a new UI using the Go Fonts.
func NewUI() *UI {
	ui := &UI{}
	// Load the theme and fonts.
	ui.Theme = material.NewTheme()
	return ui
}

// Run handles window events and renders the application.
func (ui *UI) Run(w *app.Window) error {
	// ops will be used to encode different operations.
	var ops op.Ops

	// listen for events happening on the window.
	for e := range w.Events() {
		// detect the type of the event.
		switch e := e.(type) {
		// this is sent when the application should re-render.
		case system.FrameEvent:
			// gtx is used to pass around rendering and event information.
			gtx := layout.NewContext(&ops, e)
			// handle all UI logic.
			ui.Layout(gtx)
			// render and handle the operations from the UI.
			e.Frame(gtx.Ops)

		// this is sent when the application is closed.
		case system.DestroyEvent:
			return e.Err
		}
	}

	return nil
}

// Layout handles rendering and input.
func (ui *UI) Layout(gtx layout.Context) layout.Dimensions {
	return material.H1(ui.Theme, "Hello, Gio!").Layout(gtx)
}

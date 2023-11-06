// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image/color"

	"gioui.org/layout"          // layout is used for layouting widgets.
	"gioui.org/text"            // text contains constants for text layouting.
	"gioui.org/widget/material" // material contains material design widgets.

	"github.com/egonelbre/diy-visuals-gio/internal/qapp" // qapp contains convenience funcs for the examples
)

var Theme = material.NewTheme()

func main() { qapp.Layout(Layout) }

// Layout handles rendering and input.
func Layout(gtx layout.Context) layout.Dimensions {
	return Title(Theme, "Hello, Gio!").Layout(gtx)
}

// Title creates a center aligned H1.
func Title(th *material.Theme, caption string) material.LabelStyle {
	l := material.H1(th, caption)
	l.Color = color.NRGBA{R: 127, G: 0, B: 0, A: 255}
	l.Alignment = text.Middle
	return l
}

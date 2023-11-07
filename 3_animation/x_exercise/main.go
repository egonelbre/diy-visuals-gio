// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"gioui.org/layout"
	"gioui.org/widget/material"

	"github.com/egonelbre/diy-visuals-gio/internal/qapp"
)

var Theme = material.NewTheme()

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {

		/*

			Continue your house drawing exercise.

			Here are a few ideas on what to do:

			- animate a door, window or a fence opening and closing
			- add a bird / cat that picks a random spot and then moves there
			  using some easing curve

		*/

		return layout.Dimensions{}
	})
}

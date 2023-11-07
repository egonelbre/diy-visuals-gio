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

			- add a fence by using translate and drawing the same shape multiple times
			- change the trees to use the recursive approach
			   - create multiple trees with variations
			   - remember to keep them slightly animated
			   - you can add leaves at end of branches to make it nicer
			- make the cat or dog look at the direction of the mouse
			- add perspective by translating things based on mouse position
			   - things in the background should move slower
			   - than things in the foreground
		*/

		return layout.Dimensions{}
	})
}

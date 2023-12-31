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

			Continue your house drawing exercise by making it animated.

			Here are a few ideas on what to animate:

			- grass moving
			- bushes or tree waving in the wind
			- a bird flapping wings in the sky
			- clouds moving
			- animate a whole day-night cycle by changing colors
			  - also you can animate a sun and moon

		*/

		return layout.Dimensions{}
	})
}

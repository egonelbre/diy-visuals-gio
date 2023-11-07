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

			- animate a smoking chimney using particles
			- if your house doesn't have a chimney, show it burning instead

			The basic idea for simulating particles is having few
			properties:

			  speed / velocity
			  position

			It's usually more interesting to set them to some random
			value.

			// we add together multiple forces acting on the particle
			force.X = 0.1 // wind
			force.Y = -5  // e.g. gravity

			// calculate the acceleration
			acceleration.X = force.X * deltaTime / mass
			acceleration.Y = force.Y * deltaTime / mass

			// note, it's possible to skip force calculations and mass
			// for simplicity and setting them to some constant

			// calculate the new speed value
			speed.X = speed.X + acceleration.X * deltaTime
			speed.Y = speed.Y + acceleration.Y * deltaTime

			// apply dampening to the particle, this is to prevent particles
			// becoming too fast
			speed.X = speed.X * 0.95
			speed.Y = speed.Y * 0.95

			// finally update the position
			position.X = speed.X * deltaTime
			position.Y = speed.Y * deltaTime


			// additionally whenever the particle leaves the screen you
			// can reset it to the starting position and starting speed
			// values
		*/

		return layout.Dimensions{}
	})
}

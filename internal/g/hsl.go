// SPDX-License-Identifier: Unlicense OR MIT

package g

import (
	"image/color"
)

// HSLA returns color based on HSLA in range 0..1
func HSLA(h, s, l, a float32) color.NRGBA { return RGBA(hsla(h, s, l, a)) }

// HSL returns color based on HSL in range 0..1
func HSL(h, s, l float32) color.NRGBA { return HSLA(h, s, l, 1) }

func hue(v1, v2, h float32) float32 {
	if h < 0 {
		h += 1
	}
	if h > 1 {
		h -= 1
	}
	if 6*h < 1 {
		return v1 + (v2-v1)*6*h
	} else if 2*h < 1 {
		return v2
	} else if 3*h < 2 {
		return v1 + (v2-v1)*(2.0/3.0-h)*6
	}

	return v1
}

func hsla(h, s, l, a float32) (r, g, b, ra float32) {
	if s == 0 {
		return l, l, l, a
	}

	h = Mod(h, 1)

	var v2 float32
	if l < 0.5 {
		v2 = l * (1 + s)
	} else {
		v2 = (l + s) - s*l
	}

	v1 := 2*l - v2
	r = hue(v1, v2, h+1.0/3.0)
	g = hue(v1, v2, h)
	b = hue(v1, v2, h-1.0/3.0)
	ra = a

	return
}

// SPDX-License-Identifier: Unlicense OR MIT

// Package g provides graphics helpers.
package g

import (
	"image/color"
)

var (
	White  = color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF}
	Black  = color.NRGBA{0x00, 0x00, 0x00, 0xFF}
	Red    = color.NRGBA{0xFF, 0x00, 0x00, 0xFF}
	Green  = color.NRGBA{0x00, 0xFF, 0x00, 0xFF}
	Blue   = color.NRGBA{0x00, 0x00, 0xFF, 0xFF}
	Yellow = color.NRGBA{0xFF, 0xFF, 0x00, 0xFF}

	Transparent = color.NRGBA{0xFF, 0xFF, 0xFF, 0x00}
)

// HexRGB1 takes number of 0xRGB and converts to color.NRGBA.
func HexRGB1(v uint16) color.NRGBA {
	r := byte(v>>8) & 0xF
	g := byte(v>>4) & 0xF
	b := byte(v>>0) & 0xF
	return color.NRGBA{
		R: r | r<<4,
		G: g | g<<4,
		B: b | b<<4,
		A: 0xFF,
	}
}

// HexRGBA1 takes number of 0xRGBA and converts to color.NRGBA.
func HexRGBA1(v uint16) color.NRGBA {
	r := byte(v>>12) & 0xF
	g := byte(v>>8) & 0xF
	b := byte(v>>4) & 0xF
	a := byte(v>>0) & 0xF
	return color.NRGBA{
		R: r | r<<4,
		G: g | g<<4,
		B: b | b<<4,
		A: a | a<<4,
	}
}

// HexRGB takes number of 0xRRGGBB and converts to color.NRGBA.
func HexRGB(v uint32) color.NRGBA {
	return color.NRGBA{
		R: byte(v >> 16),
		G: byte(v >> 8),
		B: byte(v >> 0),
		A: 0xFF,
	}
}

// HexRGBA takes number of 0xRRGGBBAA and converts to color.NRGBA.
func HexRGBA(v uint32) color.NRGBA {
	return color.NRGBA{
		R: byte(v >> 24),
		G: byte(v >> 16),
		B: byte(v >> 8),
		A: byte(v >> 0),
	}
}

// RGB is a helper for creating color.NRGBA.
func RGB(r, g, b float32) color.NRGBA {
	return color.NRGBA{
		R: Sat8(r),
		G: Sat8(g),
		B: Sat8(b),
		A: 0xff,
	}
}

// RGBA is a helper for creating color.NRGBA.
func RGBA(r, g, b, a float32) color.NRGBA {
	return color.NRGBA{
		R: Sat8(r),
		G: Sat8(g),
		B: Sat8(b),
		A: Sat8(a),
	}
}

// RGBAFloat returns RGBA scaled to 0..1
func RGBAFloat(c color.NRGBA) (r, g, b, a float32) {
	return float32(c.R) / 0xFF, float32(c.G) / 0xFF, float32(c.B) / 0xFF, float32(c.A) / 0xFF
}

// Lerp linearly interpolates each RGBA component separately
func RGBALerp(a, b color.NRGBA, p float32) color.NRGBA {
	ar, ag, ab, aa := RGBAFloat(a)
	br, bg, bb, ba := RGBAFloat(b)
	return RGBA(
		LerpClamp(ar, br, p),
		LerpClamp(ag, bg, p),
		LerpClamp(ab, bb, p),
		LerpClamp(aa, ba, p),
	)
}
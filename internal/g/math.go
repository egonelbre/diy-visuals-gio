// SPDX-License-Identifier: Unlicense OR MIT

package g

import (
	"math"
)

// Sat8 converts 0..1 float to 0..255 uint8
func Sat8(v float32) uint8 {
	v *= 255.0
	if v >= 255 {
		return 255
	} else if v <= 0 {
		return 0
	}
	return uint8(v)
}

func LerpClamp(p, min, max float32) float32 {
	if p < 0 {
		return min
	} else if p > 1 {
		return max
	}
	return min + (max-min)*p
}

func Mod(x, y float32) float32 { return float32(math.Mod(float64(x), float64(y))) }

func Sin(v float32) float32 {
	return float32(math.Sin(float64(v)))
}

func Cos(v float32) float32 {
	return float32(math.Cos(float64(v)))
}

func Sincos(v float32) (sn, cs float32) {
	s, c := math.Sincos(float64(v))
	return float32(s), float32(c)
}
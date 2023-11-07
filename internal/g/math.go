// SPDX-License-Identifier: Unlicense OR MIT

package g

import (
	"math"

	"gioui.org/f32"
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

func Lerp(p, min, max float32) float32 {
	return min + (max-min)*p
}

func Clamp(p, min, max float32) float32 {
	if p < 0 {
		return min
	} else if p > 1 {
		return max
	}
	return p
}

func LerpClamp(p, min, max float32) float32 {
	if p < 0 {
		return min
	} else if p > 1 {
		return max
	}
	return min + (max-min)*p
}

func PtLerp(p float32, a, b f32.Point) f32.Point {
	return f32.Point{
		X: Lerp(p, a.X, b.X),
		Y: Lerp(p, a.Y, b.Y),
	}
}

func PtLerpClamp(p float32, a, b f32.Point) f32.Point {
	return f32.Point{
		X: LerpClamp(p, a.X, b.X),
		Y: LerpClamp(p, a.Y, b.Y),
	}
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

func Sqrt(v float32) float32 {
	return float32(math.Sqrt(float64(v)))
}

func Len(p f32.Point) float32 {
	return Sqrt(p.X*p.X + p.Y*p.Y)
}

func Unit(p f32.Point) f32.Point {
	return p.Div(Len(p))
}

func Dot(a, b f32.Point) float32 {
	return a.X*b.X + a.Y*b.Y
}

func Normal(a f32.Point) f32.Point {
	return Unit(f32.Pt(-a.Y, a.X))
}

func Map(v, min, max, toMin, toMax float32) float32 {
	p := (v - min) / (max - min)
	return p*(toMax-toMin) + toMin
}

func Sign(v float32) float32 {
	switch {
	case v < 0:
		return -1
	case v > 0:
		return 1
	default:
		return 0
	}
}

func Abs(v float32) float32 {
	if v < 0 {
		return -v
	}
	return v
}

func Round(v float32) float32 {
	return float32(math.Round(float64(v)))
}

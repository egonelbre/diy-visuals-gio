// SPDX-License-Identifier: Unlicense OR MIT

package ease

import (
	"math"

	"github.com/egonelbre/diy-visuals-gio/internal/g"
)

type Func struct {
	Name string
	Ease func(float32) float32
}

var Funcs = []Func{
	Func{Name: "Linear", Ease: Linear},
	Func{Name: "EaseInQuad", Ease: EaseInQuad},
	Func{Name: "EaseOutQuad", Ease: EaseOutQuad},
	Func{Name: "EaseInOutQuad", Ease: EaseInOutQuad},
	Func{Name: "EaseInCubic", Ease: EaseInCubic},
	Func{Name: "EaseOutCubic", Ease: EaseOutCubic},
	Func{Name: "EaseInOutCubic", Ease: EaseInOutCubic},
	Func{Name: "EaseInQuart", Ease: EaseInQuart},
	Func{Name: "EaseOutQuart", Ease: EaseOutQuart},
	Func{Name: "EaseInOutQuart", Ease: EaseInOutQuart},
	Func{Name: "EaseInQuint", Ease: EaseInQuint},
	Func{Name: "EaseOutQuint", Ease: EaseOutQuint},
	Func{Name: "EaseInOutQuint", Ease: EaseInOutQuint},
	Func{Name: "Logarithm", Ease: Logarithm},
	Func{Name: "Sin", Ease: Sin},
	Func{Name: "Sqr", Ease: Sqr},
	Func{Name: "Sqrt", Ease: Sqrt},
	Func{Name: "Wobble", Ease: Wobble},
	Func{Name: "Cube", Ease: Cube},
	Func{Name: "Cube2", Ease: Cube2},
}

// no easing, no acceleration
func Linear(t float32) float32 {
	return t
}

// accelerating from zero velocity
func EaseInQuad(t float32) float32 {
	return t * t
}

// decelerating to zero velocity
func EaseOutQuad(t float32) float32 {
	return t * (2 - t)
}

// acceleration until halfway, then deceleration
func EaseInOutQuad(t float32) float32 {
	if t < .5 {
		return 2 * t * t
	}
	return -1 + (4-2*t)*t
}

// accelerating from zero velocity
func EaseInCubic(t float32) float32 {
	return t * t * t
}

// decelerating to zero velocity
func EaseOutCubic(t float32) float32 {
	t--
	return t*t*t + 1
}

// acceleration until halfway, then deceleration
func EaseInOutCubic(t float32) float32 {
	if t < .5 {
		return 4 * t * t * t
	}
	return (t-1)*(2*t-2)*(2*t-2) + 1
}

// accelerating from zero velocity
func EaseInQuart(t float32) float32 {
	return t * t * t * t
}

// decelerating to zero velocity
func EaseOutQuart(t float32) float32 {
	t--
	return 1 - t*t*t*t
}

// acceleration until halfway, then deceleration
func EaseInOutQuart(t float32) float32 {
	if t < .5 {
		return 8 * t * t * t * t
	}
	t--
	return 1 - 8*t*t*t*t
}

// accelerating from zero velocity
func EaseInQuint(t float32) float32 {
	return t * t * t * t * t
}

// decelerating to zero velocity
func EaseOutQuint(t float32) float32 {
	t--
	return 1 + t*t*t*t*t
}

// acceleration until halfway, then deceleration
func EaseInOutQuint(t float32) float32 {
	if t < .5 {
		return 16 * t * t * t * t * t
	}
	t--
	return 1 + 16*t*t*t*t*t
}

// logarithm
func Logarithm(t float32) float32 {
	return float32(math.Log2(float64(t + 1)))
}

// sin
func Sin(t float32) float32 {
	return t + g.Sin(t*2*math.Pi)/2
}

// sqr
func Sqr(t float32) float32 {
	return t * t
}

// sqrt
func Sqrt(t float32) float32 {
	return g.Sqrt(t)
}

// wobble
func Wobble(t float32) float32 {
	return 0.0005 * (t - 1) * (t - 1) * t * (44851 - 224256*t + 224256*t*t)
}

func Cube(t float32) float32 {
	s := g.Sin(t * math.Pi * 2)
	return t + s*s*s/2
}

func Cube2(t float32) float32 {
	s := g.Sin(t * math.Pi * 2)
	return t + s*s*s*s/2
}

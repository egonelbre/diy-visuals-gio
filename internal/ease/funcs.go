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
	Func{Name: "InQuad", Ease: InQuad},
	Func{Name: "OutQuad", Ease: OutQuad},
	Func{Name: "InOutQuad", Ease: InOutQuad},
	Func{Name: "InCubic", Ease: InCubic},
	Func{Name: "OutCubic", Ease: OutCubic},
	Func{Name: "InOutCubic", Ease: InOutCubic},
	Func{Name: "InQuart", Ease: InQuart},
	Func{Name: "OutQuart", Ease: OutQuart},
	Func{Name: "InOutQuart", Ease: InOutQuart},
	Func{Name: "InQuint", Ease: InQuint},
	Func{Name: "OutQuint", Ease: OutQuint},
	Func{Name: "InOutQuint", Ease: InOutQuint},
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
func InQuad(t float32) float32 {
	return t * t
}

// decelerating to zero velocity
func OutQuad(t float32) float32 {
	return t * (2 - t)
}

// acceleration until halfway, then deceleration
func InOutQuad(t float32) float32 {
	if t < .5 {
		return 2 * t * t
	}
	return -1 + (4-2*t)*t
}

// accelerating from zero velocity
func InCubic(t float32) float32 {
	return t * t * t
}

// decelerating to zero velocity
func OutCubic(t float32) float32 {
	t--
	return t*t*t + 1
}

// acceleration until halfway, then deceleration
func InOutCubic(t float32) float32 {
	if t < .5 {
		return 4 * t * t * t
	}
	return (t-1)*(2*t-2)*(2*t-2) + 1
}

// accelerating from zero velocity
func InQuart(t float32) float32 {
	return t * t * t * t
}

// decelerating to zero velocity
func OutQuart(t float32) float32 {
	t--
	return 1 - t*t*t*t
}

// acceleration until halfway, then deceleration
func InOutQuart(t float32) float32 {
	if t < .5 {
		return 8 * t * t * t * t
	}
	t--
	return 1 - 8*t*t*t*t
}

// accelerating from zero velocity
func InQuint(t float32) float32 {
	return t * t * t * t * t
}

// decelerating to zero velocity
func OutQuint(t float32) float32 {
	t--
	return 1 + t*t*t*t*t
}

// acceleration until halfway, then deceleration
func InOutQuint(t float32) float32 {
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

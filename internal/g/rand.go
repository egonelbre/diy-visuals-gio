// SPDX-License-Identifier: Unlicense OR MIT

package g

import (
	"math/rand"

	"gioui.org/f32"
)

func Rand() float32 {
	return rand.Float32()
}

func RandPt() f32.Point {
	return f32.Pt(Rand(), Rand())
}

func RandUnitPt() f32.Point {
	return f32.Pt(Rand()-0.5, Rand()-0.5).Mul(2)
}

func RandRange(min, max float32) float32 {
	return Rand()*(max-min) + min
}

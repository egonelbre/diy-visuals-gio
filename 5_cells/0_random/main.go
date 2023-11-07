// SPDy-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"time"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"

	"github.com/egonelbre/diy-visuals-gio/internal/g"
	"github.com/egonelbre/diy-visuals-gio/internal/qapp"
)

var Theme = material.NewTheme()
var Start time.Time
var Update float32
var LastTime time.Time

func main() {
	grid := NewGrid(20, 20)

	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		paint.Fill(gtx.Ops, g.Black)
		op.InvalidateOp{}.Add(gtx.Ops)

		if LastTime.IsZero() {
			Start = gtx.Now
			LastTime = gtx.Now
		}
		deltaTime := float32(gtx.Now.Sub(LastTime).Seconds())
		LastTime = gtx.Now

		Update -= deltaTime
		if Update < 0 {
			grid.Randomize()
			Update = 1
		}

		grid.Layout(gtx)

		return layout.Dimensions{}
	}, app.Size(1024, 1024))
}

type Grid struct {
	Cells []Cell
	Max   image.Point
}

func NewGrid(width, height int) *Grid {
	return &Grid{
		Cells: make([]Cell, width*height),
		Max:   image.Point{X: width, Y: height},
	}
}

func (grid *Grid) At(x, y int) *Cell {
	if x < 0 || x >= grid.Max.X {
		return nil
	}
	if y < 0 || y >= grid.Max.X {
		return nil
	}
	return &grid.Cells[grid.Offset(x, y)]
}

func (grid *Grid) Offset(x, y int) int {
	return y*grid.Max.X + x
}

func (grid *Grid) Randomize() {
	for i := range grid.Cells {
		cell := &grid.Cells[i]
		cell.Alive = g.Rand() > 0.5
	}
}

func (grid *Grid) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Center.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			size := gtx.Constraints.Max

			cellSize := min(
				size.X/grid.Max.X,
				size.Y/grid.Max.Y,
			)

			size = image.Point{
				X: cellSize * grid.Max.X,
				Y: cellSize * grid.Max.Y,
			}

			for i := range grid.Cells {
				cell := &grid.Cells[i]
				x := i % grid.Max.X
				y := i / grid.Max.X

				cell.Draw(gtx, image.Rectangle{
					Min: image.Pt(x*cellSize, y*cellSize),
					Max: image.Pt((x+1)*cellSize, (y+1)*cellSize),
				}, cellSize)
			}

			return layout.Dimensions{Size: size}
		})
}

type Cell struct {
	Alive bool
}

func (cell *Cell) IsAlive() bool {
	if cell == nil {
		return false
	}
	return cell.Alive
}

func (cell *Cell) Draw(gtx layout.Context, at image.Rectangle, cellSize int) {
	if cell.Alive {
		center := at.Inset(cellSize / 5)
		paint.FillShape(gtx.Ops, g.White, clip.Rect(center).Op())
	}
}

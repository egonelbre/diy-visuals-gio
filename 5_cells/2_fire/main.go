// SPDy-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"slices"
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

func main() {
	grid := NewGrid(25, 25)
	grid.Randomize()

	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		paint.Fill(gtx.Ops, g.Black)
		op.InvalidateOp{}.Add(gtx.Ops)

		if Start.IsZero() {
			Start = gtx.Now
		}
		now := float32(gtx.Now.Sub(Start).Seconds())

		dx := g.Sin(now*10) * 3
		grid.AddHeat(
			float32(grid.Max.X)/2+dx,
			float32(grid.Max.Y)-2,
			5)
		grid.Update()
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

func (grid *Grid) Clone() *Grid {
	return &Grid{
		Cells: slices.Clone(grid.Cells),
		Max:   grid.Max,
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

func (grid *Grid) IndexToPos(i int) (int, int) {
	x := i % grid.Max.X
	y := i / grid.Max.X
	return x, y
}

func (grid *Grid) Offset(x, y int) int {
	return y*grid.Max.X + x
}

func (grid *Grid) Randomize() {
	for i := range grid.Cells {
		cell := &grid.Cells[i]
		cell.Randomize()
	}
}

func (grid *Grid) AddHeat(x, y float32, amount float32) {
	lowX := g.Round(x)
	lx := int(lowX)
	lowY := g.Round(y)
	ly := int(lowY)

	dx := int(g.Sign(x - lowX))
	dy := int(g.Sign(y - lowY))

	contribLowX := 1 - g.Abs(x-lowX)
	contribLowY := 1 - g.Abs(y-lowY)

	grid.At(lx, ly).Add(amount * contribLowX * contribLowY)
	grid.At(lx+dx, ly).Add(amount * (1 - contribLowX) * contribLowY)
	grid.At(lx, ly+dy).Add(amount * contribLowX * (1 - contribLowY))
	grid.At(lx+dx, ly+dy).Add(amount * (1 - contribLowX) * (1 - contribLowY))
}

func (grid *Grid) Update() {
	snap := grid.Clone()

	for i := range grid.Cells {
		cell := &grid.Cells[i]
		x, y := grid.IndexToPos(i)
		cell.Update(image.Pt(x, y), snap)
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
				x, y := grid.IndexToPos(i)

				cell.Draw(gtx, image.Rectangle{
					Min: image.Pt(x*cellSize, y*cellSize),
					Max: image.Pt((x+1)*cellSize, (y+1)*cellSize),
				}, cellSize)
			}

			return layout.Dimensions{Size: size}
		})
}

type Cell struct {
	Temp float32
	Dead bool
}

func (cell *Cell) Randomize() {
	cell.Temp = g.Rand()
	cell.Temp *= cell.Temp
	cell.Temp *= cell.Temp
	cell.Temp *= cell.Temp
}

func (cell *Cell) Add(heat float32) {
	if cell == nil || cell.Dead {
		return
	}
	cell.Temp += heat
}

func (cell *Cell) Heat() float32 {
	if cell == nil {
		return 0
	}
	return cell.Temp
}

func (cell *Cell) Update(at image.Point, lastGrid *Grid) {
	t := float32(0)
	t += lastGrid.At(at.X-1, at.Y+1).Heat()
	t += lastGrid.At(at.X, at.Y+1).Heat()
	t += lastGrid.At(at.X+1, at.Y+1).Heat()
	t += lastGrid.At(at.X, at.Y+2).Heat()
	cell.Temp = t / 4.1
}

func (cell *Cell) Draw(gtx layout.Context, at image.Rectangle, cellSize int) {
	center := at.Inset(cellSize / 5)
	red := cell.Temp
	green := cell.Temp * cell.Temp
	color := g.RGB(red, green, 0)
	paint.FillShape(gtx.Ops, color, clip.Rect(center).Op())
}

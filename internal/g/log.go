// SPDX-License-Identifier: Unlicense OR MIT

package g

import (
	"fmt"
	"strings"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type Log struct {
	Theme *material.Theme
	Size  unit.Sp
	Lines []string
}

func (log *Log) Print(a ...any) {
	log.Lines = append(log.Lines, fmt.Sprint(a...))
}

func (log *Log) Printf(format string, a ...any) {
	log.Lines = append(log.Lines, fmt.Sprintf(format, a...))
}

func (log *Log) Layout(gtx layout.Context) layout.Dimensions {
	lbl := material.Label(log.Theme, log.Size, strings.Join(log.Lines, "\n"))
	lbl.Font.Typeface = "monospace"
	return lbl.Layout(gtx)
}

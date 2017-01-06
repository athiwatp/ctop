package widgets

import (
	"fmt"
	"math"
	"strconv"

	ui "github.com/gizak/termui"
)

const (
	kb = 1024
	mb = kb * 1024
	gb = mb * 1024
)

func byteFormat(n int64) string {
	if n < kb {
		return fmt.Sprintf("%sB", strconv.FormatInt(n, 10))
	}
	if n < mb {
		n = n / kb
		return fmt.Sprintf("%sK", strconv.FormatInt(n, 10))
	}
	if n < gb {
		n = n / mb
		return fmt.Sprintf("%sM", strconv.FormatInt(n, 10))
	}
	n = n / gb
	return fmt.Sprintf("%sG", strconv.FormatInt(n, 10))
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func compactPar(s string) *ui.Par {
	p := ui.NewPar(s)
	p.Border = false
	p.Height = 1
	p.Width = 20
	p.TextFgColor = ui.ColorWhite
	return p
}

func mkGauge() *ui.Gauge {
	g := ui.NewGauge()
	g.Height = 1
	g.Border = false
	g.Percent = 0
	g.PaddingBottom = 0
	g.BarColor = ui.ColorGreen
	g.Label = "-"
	return g
}

func colorScale(n int) ui.Attribute {
	if n > 70 {
		return ui.ColorRed
	}
	if n > 30 {
		return ui.ColorYellow
	}
	return ui.ColorGreen
}
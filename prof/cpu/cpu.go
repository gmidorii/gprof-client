package cpu

import (
	"fmt"

	ui "github.com/gizak/termui"
	"github.com/midorigreen/gprof-client/prof"
)

type CPUWidget struct {
	Widget []ui.GridBufferer
}

func CreateWidget() *CPUWidget {
	return &CPUWidget{}
}

func (c *CPUWidget) Update(prof prof.Prof) {
	cpu := prof.Data.CPU
	for i := range c.Widget {
		gauge, ok := c.Widget[i].(*ui.Gauge)
		if !ok {
			continue
		}
		gauge.Percent = int(cpu.Cores[i].Percent)
	}
}

func (c *CPUWidget) Create(prof prof.Prof) []ui.GridBufferer {
	cpu := prof.Data.CPU
	c.Widget = make([]ui.GridBufferer, len(cpu.Cores))
	for i, v := range cpu.Cores {
		g := ui.NewGauge()
		g.Percent = int(v.Percent)
		g.BorderLabel = fmt.Sprintf("cpu %d", i)
		c.Widget[i] = g
	}
	return c.Widget
}

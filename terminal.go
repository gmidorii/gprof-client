package main

import (
	"fmt"

	ui "github.com/gizak/termui"
)

var cpuWidget []ui.GridBufferer

type CPUWidget struct {
	Widget []ui.GridBufferer
}

func (c *CPUWidget) updateCPU(prof Prof) {
	cpu := prof.Data.CPU
	for i := range c.Widget {
		gauge, ok := c.Widget[i].(*ui.Gauge)
		if !ok {
			continue
		}
		gauge.Percent = int(cpu.Cores[i].Percent)
	}
}

func (c *CPUWidget) createCPU(prof Prof) []ui.GridBufferer {
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

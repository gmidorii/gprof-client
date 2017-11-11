package main

import (
	"fmt"

	ui "github.com/gizak/termui"
)

func display(prof Prof) error {
	gauges := cpuWidget(prof.Data.CPU)
	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, gauges...),
		),
	)
	return nil
}

func cpuWidget(cpu CPU) []ui.GridBufferer {
	gauges := make([]ui.GridBufferer, len(cpu.Cores))
	for i, v := range cpu.Cores {
		g := ui.NewGauge()
		g.Percent = int(v.Percent)
		g.BorderLabel = fmt.Sprintf("cpu %d", i)
		gauges[i] = g
	}
	return gauges
}

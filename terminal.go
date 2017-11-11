package main

import (
	"fmt"

	ui "github.com/gizak/termui"
)

func display(prof Prof) error {
	if err := ui.Init(); err != nil {
		return err
	}
	defer ui.Close()

	gauges := cpuWidget(prof.Data.CPU)

	fmt.Println(gauges)
	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, gauges...),
		),
	)

	ui.Body.Align()

	ui.Render(ui.Body)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Loop()
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

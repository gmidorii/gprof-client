package disk

import (
	"log"

	ui "github.com/gizak/termui"
	"github.com/midorigreen/gprof-client/prof"
)

type DiskWidget struct {
	Widget ui.GridBufferer
}

func CreateWidget() *DiskWidget {
	return &DiskWidget{}
}

func (d *DiskWidget) Update(prof prof.Prof) {
	disk := prof.Data.Disk
	mbar, ok := d.Widget.(*ui.MBarChart)
	if !ok {
		return
	}
	mbar.Data[0] = []int{mbToGB(disk.Usage.Used), mbToGB(disk.Usage.Total)}
	mbar.Data[1] = []int{mbToGB(disk.Usage.Free), 0}
}

func (d *DiskWidget) Create(prof prof.Prof) []ui.GridBufferer {
	disk := prof.Data.Disk
	log.Println(prof.Data.Disk)
	mbar := ui.NewMBarChart()
	mbar.Data[0] = []int{mbToGB(disk.Usage.Used), mbToGB(disk.Usage.Total)}
	mbar.Data[1] = []int{mbToGB(disk.Usage.Free), 0}
	mbar.DataLabels = []string{"memory", "test"}
	mbar.Width = 50
	mbar.Height = 10
	mbar.BarWidth = 20
	mbar.BarGap = 10
	mbar.BarColor[0] = ui.ColorBlue
	mbar.BarColor[1] = ui.ColorYellow

	d.Widget = mbar
	return []ui.GridBufferer{d.Widget}
}

func mbToGB(b int) int {
	return b / 1024
}

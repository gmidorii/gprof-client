package disk

import (
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
	bar, ok := d.Widget.(*ui.BarChart)
	if !ok {
		return
	}
	bar.Data = []int{disk.Usage.Total, disk.Usage.Used, disk.Usage.Free}
}

func (d *DiskWidget) Create(prof prof.Prof) []ui.GridBufferer {
	disk := prof.Data.Disk
	bar := ui.NewBarChart()
	bar.DataLabels = []string{"T", "U", "F"}
	bar.BorderLabel = "Disk memory"
	bar.Data = []int{disk.Usage.Total, disk.Usage.Used, disk.Usage.Free}
	bar.Height = 10
	bar.Width = 40

	d.Widget = bar
	return []ui.GridBufferer{d.Widget}
}

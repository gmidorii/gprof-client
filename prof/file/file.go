package file

import (
	"fmt"

	ui "github.com/gizak/termui"
	"github.com/midorigreen/gprof-client/prof"
)

type FileWidget struct {
	Widget []ui.GridBufferer
}

func CreateWidget() *FileWidget {
	return &FileWidget{}
}

func (f *FileWidget) Create(prof prof.Prof) []ui.GridBufferer {
	file := prof.Data.File
	par := ui.NewPar(file.Content)
	par.BorderLabel = fmt.Sprintf("%s (%s)", file.Name, file.UpdatedTime)
	par.Height = 10
	f.Widget = []ui.GridBufferer{par}
	return f.Widget
}

func (f *FileWidget) Update(prof prof.Prof) {
	file := prof.Data.File
	for i := range f.Widget {
		par, ok := f.Widget[i].(*ui.Par)
		if !ok {
			continue
		}
		par.Text = file.Content
		par.BorderLabel = fmt.Sprintf("%s (%s)", file.Name, file.UpdatedTime)
	}
}

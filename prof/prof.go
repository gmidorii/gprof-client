package prof

import (
	ui "github.com/gizak/termui"
)

type ProfWidget interface {
	Create(Prof) []ui.GridBufferer
	Update(Prof)
}

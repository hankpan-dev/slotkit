package slotapp

import (
	"slotkit"
)

const TotalPalines = 9

var Paylines []*slotkit.Payline = make([]*slotkit.Payline, TotalPalines)

func init() {
	Paylines[0] = slotkit.NewPayline(0, 0, 0, 0, 0)
	Paylines[1] = slotkit.NewPayline(1, 1, 1, 1, 1)
	Paylines[2] = slotkit.NewPayline(2, 2, 2, 2, 2)
	Paylines[3] = slotkit.NewPayline(0, 0, 1, 0, 0)
	Paylines[4] = slotkit.NewPayline(2, 2, 1, 2, 2)
	Paylines[5] = slotkit.NewPayline(0, 0, 1, 2, 2)
	Paylines[6] = slotkit.NewPayline(2, 2, 1, 0, 0)
	Paylines[7] = slotkit.NewPayline(0, 1, 2, 1, 0)
	Paylines[8] = slotkit.NewPayline(2, 1, 0, 1, 2)
}

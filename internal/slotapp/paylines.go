package slotapp

import (
	"slotkit"
)

const TotalPalines = 9

var Paylines []*slotkit.Payline = make([]*slotkit.Payline, TotalPalines)

// 	0	3	6	9	12
// 	1	4	7	10	13
// 	2	5	8	11	14
func init() {
	Paylines[0] = slotkit.NewPayline(0, 3, 6, 9, 12)
	Paylines[1] = slotkit.NewPayline(1, 4, 7, 10, 13)
	Paylines[2] = slotkit.NewPayline(2, 5, 8, 11, 14)
	Paylines[3] = slotkit.NewPayline(0, 3, 7, 9, 12)
	Paylines[4] = slotkit.NewPayline(2, 5, 7, 11, 14)
	Paylines[5] = slotkit.NewPayline(0, 3, 7, 11, 14)
	Paylines[6] = slotkit.NewPayline(2, 5, 7, 9, 12)
	Paylines[7] = slotkit.NewPayline(0, 4, 8, 10, 12)
	Paylines[8] = slotkit.NewPayline(2, 4, 6, 10, 14)
}

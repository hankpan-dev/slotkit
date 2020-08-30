package slotapp

import (
	"slotlib"
)

const TotalPalines = 9

var Paylines []*slotlib.Payline = make([]*slotlib.Payline, TotalPalines)

// 	0	3	6	9	12
// 	1	4	7	10	13
// 	2	5	8	11	14
func init() {
	Paylines[0] = slotlib.NewPayline(0, 3, 6, 9, 12)
	Paylines[1] = slotlib.NewPayline(1, 4, 7, 10, 13)
	Paylines[2] = slotlib.NewPayline(2, 5, 8, 11, 14)
	Paylines[3] = slotlib.NewPayline(0, 3, 7, 9, 12)
	Paylines[4] = slotlib.NewPayline(2, 5, 7, 11, 14)
	Paylines[5] = slotlib.NewPayline(0, 3, 7, 11, 14)
	Paylines[6] = slotlib.NewPayline(2, 5, 7, 9, 12)
	Paylines[7] = slotlib.NewPayline(0, 4, 8, 10, 12)
	Paylines[8] = slotlib.NewPayline(2, 4, 6, 10, 14)
}

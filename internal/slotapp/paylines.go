package slotapp

import (
	"slotlib/slots"
)

const TotalPalines = 9

var Paylines []*slots.Payline = make([]*slots.Payline, TotalPalines)

// 	0	3	6	9	12
// 	1	4	7	10	13
// 	2	5	8	11	14
func init() {
	Paylines[0] = slots.NewPayline(0, 3, 6, 9, 12)
	Paylines[1] = slots.NewPayline(1, 4, 7, 10, 13)
	Paylines[2] = slots.NewPayline(2, 5, 8, 11, 14)
	Paylines[3] = slots.NewPayline(0, 3, 7, 9, 12)
	Paylines[4] = slots.NewPayline(2, 5, 7, 11, 14)
	Paylines[5] = slots.NewPayline(0, 3, 7, 11, 14)
	Paylines[6] = slots.NewPayline(2, 5, 7, 9, 12)
	Paylines[7] = slots.NewPayline(0, 4, 8, 10, 12)
	Paylines[8] = slots.NewPayline(2, 4, 6, 10, 14)
}

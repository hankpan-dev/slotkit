package slotapp

const PaytableLength = 8

var Paytable [][]int = make([][]int, PaytableLength)

func init() {
	Paytable[0] = []int{0, 0, 2, 3, 5}
	Paytable[1] = []int{0, 0, 3, 4, 8}
	Paytable[2] = []int{0, 0, 4, 5, 10}
	Paytable[3] = []int{0, 0, 5, 7, 14}
	Paytable[4] = []int{0, 0, 6, 9, 18}
	Paytable[5] = []int{0, 0, 7, 12, 25}
	Paytable[6] = []int{0, 0, 10, 15, 30}
	Paytable[7] = []int{0, 0, 20, 30, 60}
}

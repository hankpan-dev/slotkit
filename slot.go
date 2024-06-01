package slotkit

// Slot :
type Slot struct {
	cells Cells
	reels []Cells
}

// NewSlot : 建立相同滾輪長度的 slot
func NewSlot(reels int, rows int) *Slot {
	layout := make([]int, reels)
	for i := 0; i < reels; i++ {
		layout[i] = rows
	}
	return NewSlotFlex(layout)
}

// NewSlotFlex : 建立滾輪長度不同的 slot
func NewSlotFlex(rows []int) *Slot {

	length := 0
	for _, size := range rows {
		length += size
	}

	s := &Slot{}
	s.cells = NewCells(length)
	s.reels = make([]Cells, len(rows))
	index := 0
	for i := range s.reels {
		length = rows[i]
		s.reels[i] = s.cells[index : index+length] // s.Slice(index, length)
		index += length
	}
	return s
}

// Reel : 取得指定滾輪
func (s Slot) Reel(index int) Cells {
	return s.reels[index]
}

// Reels : 取得所有滾輪
func (s Slot) Reels() []Cells {
	return s.reels
}

// TotalReels : 取得滾輪數量
func (s Slot) TotalReels() int {
	return len(s.reels)
}

// Cells : 取得所有盤面內容
func (s Slot) Cells() Cells {
	return s.cells
}

func (s Slot) String() string {
	var str string = "["
	for _, reel := range s.reels {
		str += " " + reel.String()
	}
	str += " ]"
	return str
}

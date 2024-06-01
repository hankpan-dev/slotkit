package slotkit

type Payline struct {
	pos []int
}

// NewPayline : 建立 Payline 物件
func NewPayline(pos ...int) *Payline {
	p := &Payline{pos: pos}
	return p
}

// Positions : 取得 pos 陣列
func (p Payline) Positions() []int {
	return p.pos
}

// FetchSymbols : 取得 Payline 上的圖騰
func (p Payline) FetchSymbols(slot *Slot) Cells {
	strip := NewCells(len(p.pos))
	for i, pos := range p.pos {
		var cel *Cell
		if i < slot.TotalReels() && pos < len(slot.Reel(i)) {
			cel = slot.Reel(i)[pos]
		}
		strip[i] = cel
	}
	return strip
}

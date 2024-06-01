package slotkit

// Cell : 圖騰格
type Cell struct {
	Symbol Symbol
	Param  interface{}
}

// Cells :
type Cells []*Cell

// NewCells : 建立圖騰格陣列
func NewCells(length int) Cells {
	cells := make([]*Cell, length)
	for i := range cells {
		cells[i] = &Cell{}
	}
	return cells
}

// At : 以循環方式轉換 index
//
//	ex. index = -1, 取得最後一個位置
//	    index = length, 取得第一個位置.
func (c Cells) At(index int) (int, *Cell) {
	length := len(c)
	index = ((index % length) + length) % length
	return index, c[index]
}

// Bound : 檢查 index 是否在彩帶範圍內
func (c Cells) Bound(index int) bool {
	return (0 <= index && len(c) > index)
}

// Swap : 交換圖騰格內容
func (c *Cells) Swap(a int, b int) bool {
	if c.Bound(a) && c.Bound(b) {
		cell := (*c)[a]
		(*c)[a] = (*c)[b]
		(*c)[b] = cell
		return true
	}
	return false
}

// SetSymbols : 設定圖騰 返回次數 (略過 nil 的圖騰格)
func (c *Cells) SetSymbols(index int, symbols ...Symbol) int {
	for i, symbol := range symbols {
		if !c.Bound(index) {
			return i
		}
		if cell := (*c)[index]; cell != nil {
			cell.Symbol = symbol
		}
		index++
	}
	// 設定全部完成, 傳回數量
	return len(symbols)
}

// CountID : 計算相同 id 的數量
func (c Cells) CountID(id int) int {
	count := 0
	for _, cell := range c {
		if cell.Symbol.Id() == id {
			count++
		}
	}
	return count
}

// CountMatch : 計算相符圖騰 flag 的數量
func (c Cells) CountMatch(mask uint32) int {
	count := 0
	for _, cell := range c {
		if cell.Symbol.Match(mask) {
			count++
		}
	}
	return count
}

// CountEqual : 計算相同圖騰數量 (id 與 flag 皆相等)
func (c Cells) CountEqual(symbol Symbol) int {
	count := 0
	for _, cell := range c {
		if cell.Symbol.Equal(symbol) {
			count++
		}
	}
	return count
}

// FindMatch : 尋找相符合的圖騰, 並返回索引位置 (失敗 -1)
func (c Cells) FindMatch(mask uint32, start int) int {
	return c.find(mask, start, true)
}

// FindUnmatch : 尋找相不符合的圖騰, 並返回索引位置 (失敗 -1)
func (c Cells) FindUnmatch(mask uint32, start int) int {
	return c.find(mask, start, false)
}

// Combinations : 檢查連線組合, 傳回連線長度
func (c Cells) Combinations(mask uint32) int {
	count := 0
	for _, cell := range c {
		if !cell.Symbol.Match(mask) {
			return count
		}
		count++
	}
	return count
}

// Reverse : 反轉彩帶
func (c Cells) Reverse() Cells {
	size := len(c)
	cells := make(Cells, size)
	for i, cell := range c {
		cells[size-i-1] = cell
	}
	return cells
}

// GetSymbolsId : 取得彩帶上的圖騰
func (c Cells) GetSymbolsId() []int {
	ids := make([]int, len(c))
	for i, cell := range c {
		ids[i] = cell.Symbol.Id()
	}
	return ids
}

func (c Cells) find(mask uint32, start int, match bool) int {
	for i := start; i < len(c); i++ {
		cell := c[i]

		if nil == cell || nil == cell.Symbol {
			break
		}
		if match == cell.Symbol.Match(mask) {
			return i
		}
	}
	return -1
}

// ResetParams : 清除 Param
func (c *Cells) ResetParams() {
	for _, pos := range *c {
		pos.Param = nil
	}
}

func (c Cells) String() string {
	var str string = "["
	for _, pos := range c {
		str += " " + pos.Symbol.String()
	}
	str += " ]"
	return str
}

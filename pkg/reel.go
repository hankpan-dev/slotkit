package slotlib

// Reel :
type Reel struct {
	pos []*Pos
}

// Pos :
type Pos struct {
	Symbol *Symbol
	Param  interface{}
}

// NewReel :
func NewReel(size int) *Reel {
	r := &Reel{}
	r.Create(size)
	return r
}

// Create :
func (r *Reel) Create(size int) {
	r.pos = make([]*Pos, size)

	for i := range r.pos {
		r.pos[i] = &Pos{}
	}
}

// Pos : 取得 pos 陣列
func (r *Reel) Pos() []*Pos {
	return r.pos
}

// Length : 取得 reel 長度
func (r *Reel) Length() int {
	return len(r.pos)
}

// Index : 取得滾輪的 Pos
func (r *Reel) Index(index int) *Pos {
	if 0 <= index && len(r.pos) > index {
		return r.pos[index]
	}
	return nil
}

// SetSymbols :
func (r *Reel) SetSymbols(start int, symbols ...*Symbol) {

	index := start
	for _, symbol := range symbols {

		pos := r.Index(index)
		index++

		if nil != pos {
			pos.Symbol = symbol
		} else {
			break
		}
	}
}

// CountId : 計算相同 id 的數量
func (r *Reel) CountId(id int) int {
	count := 0
	for _, pos := range r.pos {
		if pos.Symbol.Id() == id {
			count++
		}
	}
	return count
}

// CountMatch : 計算相符圖騰 flag 的數量
func (r *Reel) CountMatch(mask uint32) int {
	count := 0
	for _, pos := range r.pos {
		if true == pos.Symbol.Match(mask) {
			count++
		}
	}
	return count
}

// CountEqual : 計算相同圖騰數量 (id 與 flag 皆相等)
func (r *Reel) CountEqual(symbol *Symbol) int {
	count := 0
	for _, pos := range r.pos {
		if true == pos.Symbol.Equal(symbol) {
			count++
		}
	}
	return count
}

// Switch : 將滾輪上的兩個 Pos 交換
func (r *Reel) Switch(a int, b int) bool {
	// 檢查交換位置是否超出範圍
	if (0 <= a && len(r.pos) > a) && (0 <= b && len(r.pos) > b) {
		if a != b {
			pos := r.pos[a]
			r.pos[a] = r.pos[b]
			r.pos[b] = pos
		}
		return true
	}
	return false
}

// ResetParams : 清除 Param
func (r *Reel) ResetParams() {
	for _, pos := range r.pos {
		pos.Param = nil
	}
}

func (r *Reel) String() string {
	var str string = "["
	for _, pos := range r.pos {
		str += " " + pos.Symbol.String()
	}
	str += " ]"
	return str
}

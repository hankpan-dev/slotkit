package slotlib

type Payline struct {
	list []int
}

// NewPayline : 建立 Payline 物件
func NewPayline(positions ...int) *Payline {
	p := &Payline{}
	p.Create(positions...)
	return p
}

// Create : 建立 Payline 物件
func (p *Payline) Create(positions ...int) {
	p.list = positions
}

// List : 取得 pos 陣列
func (p *Payline) List() []int {
	return p.list
}

// CountId : 計算相同 id 的數量
func (p *Payline) CountId(slot *Slot, id int) int {
	count := 0
	for _, index := range p.list {
		pos := slot.Index(index)
		if nil != pos && pos.Symbol.Id() == id {
			count++
		}
	}
	return count
}

// CountMatch : 計算相符圖騰 flag 的數量
func (p *Payline) CountMatch(slot *Slot, mask uint32) int {
	count := 0
	for _, index := range p.list {
		pos := slot.Index(index)
		if nil != pos && true == pos.Symbol.Match(mask) {
			count++
		}
	}
	return count
}

// CountEqual : 計算相同圖騰數量 (id 與 flag 皆相等)
func (p *Payline) CountEqual(slot *Slot, symbol *Symbol) int {
	count := 0
	for _, index := range p.list {
		pos := slot.Index(index)
		if nil != pos && true == pos.Symbol.Equal(symbol) {
			count++
		}
	}
	return count
}

// FindMatch : 尋找相符合的圖騰
// start : 起始位置
// dir : false = 反向搜尋
func (p *Payline) FindMatch(slot *Slot, mask uint32, start int, dir bool) (int, *Pos) {

	index := start
	enum := Enum(len(p.list), start, dir)
	for nil != enum {
		idx := p.list[enum.Index()]
		pos := slot.Index(idx)

		// 檢查錯誤
		if nil == pos || nil == pos.Symbol {
			break
		}
		if true == pos.Symbol.Match(mask) {
			return index, pos
		}
		// 列舉下一個
		index++
		enum = enum.Next()
	}

	return -1, nil
}

// Combinations : 檢查連線組合, 傳回連線長度
func (p *Payline) Combinations(slot *Slot, mask uint32, dir bool) int {

	count := 0
	enum := Enum(len(p.list), 0, dir)

	for nil != enum {
		idx := p.list[enum.Index()]
		pos := slot.Index(idx)

		// 檢查位置是否超出範圍
		if nil == pos || nil == pos.Symbol {
			return 0
		}

		if false == pos.Symbol.Match(mask) {
			return count
		}
		count++
		enum = enum.Next()
	}

	return count
}

type enum struct {
	index  int
	length int
	step   int
}

func Enum(length int, start int, direction bool) *enum {
	// 檢查邊界
	if 0 <= start && length > start {
		// 初始化 enum
		e := &enum{index: start, length: length, step: 1}

		if false == direction {
			e.index = e.length - e.index - 1
			e.step = -1
		}
		return e
	}
	// 索引超出範圍時 傳回 nil
	return nil
}

func (e *enum) Index() int {
	return e.index
}

func (e *enum) Next() *enum {
	next := e.index + e.step
	if 0 <= next && e.length > next {
		e.index = next
		return e
	}
	return nil
}

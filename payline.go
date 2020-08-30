package slotlib

type Payline struct {
	pos   []int
	strip *Strip
}

// NewPayline : 建立 Payline 物件
func NewPayline(positions ...int) *Payline {
	p := &Payline{}
	p.Create(positions...)
	return p
}

// Create : 建立 Payline 物件
func (p *Payline) Create(positions ...int) {
	p.pos = positions
	p.strip = NewStrip(len(positions))
}

// Positions : 取得 pos 陣列
func (p *Payline) Positions() []int {
	return p.pos
}

// SetSymbols : 設定 Payline 上的圖騰
func (p *Payline) SetSymbols(slot *Slot) {
	for i, pos := range p.pos {
		stop := slot.Index(pos)
		p.strip.stops[i] = stop
	}
}

// CountId : 計算相同 id 的數量
func (p *Payline) CountID(id int) int {
	return p.strip.CountID(id)
}

// CountMatch : 計算相符圖騰 flag 的數量
func (p *Payline) CountMatch(mask uint32) int {
	return p.strip.CountMatch(mask)
}

// CountEqual : 計算相同圖騰數量 (id 與 flag 皆相等)
func (p *Payline) CountEqual(symbol *Symbol) int {
	return p.strip.CountEqual(symbol)
}

// FindMatch : 尋找相符合的圖騰
// start : 起始位置
// dir : false = 反向搜尋
func (p *Payline) FindMatch(mask uint32, start int, dir bool) (int, *Stop) {
	return p.strip.FindMatch(mask, start, dir)
}

// FindUnmatch : 尋找不相符的圖騰
func (p *Payline) FindUnmatch(mask uint32, start int, dir bool) (int, *Stop) {
	return p.strip.FindUnmatch(mask, start, dir)
}

// Combinations : 檢查連線組合, 傳回連線長度
func (p *Payline) Combinations(mask uint32, dir bool) int {

	index := 0
	count := 0
	if false == dir {
		index = p.strip.Length() - 1
	}

	for enum := p.strip.Enum(index, dir); nil != enum.Index(); enum.Next() {
		stop := enum.Index()

		// 檢查位置是否超出範圍
		if nil == stop || nil == stop.Symbol {
			return 0
		}

		if false == stop.Symbol.Match(mask) {
			return count
		}
		count++
	}

	return count
}

func (p *Payline) String() string {
	return p.strip.String()
}

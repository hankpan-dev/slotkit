package slotlib

// Stop
type Stop struct {
	Symbol *Symbol
	Params interface{}
}

// Strip :
type Strip struct {
	stops []*Stop
}

// NewStrip : 建立彩帶
func NewStrip(length int) *Strip {
	s := &Strip{}
	s.Create(length)
	return s
}

// Create : 建立彩帶
func (s *Strip) Create(length int) {
	s.stops = make([]*Stop, length)
	for i := range s.stops {
		s.stops[i] = &Stop{}
	}
}

// Slice : 取得彩帶切片
func (s *Strip) Slice(start int, length int) *Strip {
	strip := &Strip{
		stops: s.stops[start : start+length],
	}
	return strip
}

// Index : 索引彩帶上的 Stop.
func (s *Strip) Index(index int) *Stop {
	if true == s.Bound(index) {
		return s.stops[index]
	}
	return nil
}

// Cycler : 以循環方式轉換 index.
// 當 index = -1 時, 返回最後一個位置的索引值.
// 當 index = length 時, 返回第一個位置的索引值.
func (s *Strip) Cycler(index int) int {
	length := s.Length()
	return ((index % length) + length) % length
}

// Length : 取得彩帶長度
func (s *Strip) Length() int {
	return len(s.stops)
}

// Bound : 檢查 index 是否在彩帶範圍內
func (s *Strip) Bound(index int) bool {
	return (0 <= index && s.Length() > index)
}

// Swap : 交換 Stop
func (s *Strip) Swap(a int, b int) bool {
	if true == s.Bound(a) && true == s.Bound(b) {
		stop := s.stops[a]
		s.stops[a] = s.stops[b]
		s.stops[b] = stop
		return true
	}
	return false
}

// SetSymbols : 設定彩帶圖騰
func (s *Strip) SetSymbols(index int, symbols ...*Symbol) int {
	for i, symbol := range symbols {
		// 取得 Stop
		stop := s.Index(index)
		if nil != stop {
			stop.Symbol = symbol
			index++
		} else {
			// 傳回設定的數量
			return i
		}
	}
	// 設定全部完成, 傳回數量
	return s.Length()
}

// CountID : 計算相同 id 的數量
func (s *Strip) CountID(id int) int {
	count := 0
	for _, stop := range s.stops {
		if stop.Symbol.Id() == id {
			count++
		}
	}
	return count
}

// CountMatch : 計算相符圖騰 flag 的數量
func (s *Strip) CountMatch(mask uint32) int {
	count := 0
	for _, stop := range s.stops {
		if true == stop.Symbol.Match(mask) {
			count++
		}
	}
	return count
}

// CountEqual : 計算相同圖騰數量 (id 與 flag 皆相等)
func (s *Strip) CountEqual(symbol *Symbol) int {
	count := 0
	for _, stop := range s.stops {
		if true == stop.Symbol.Equal(symbol) {
			count++
		}
	}
	return count
}

// FindMatch : 尋找相符合的圖騰.
// start : 起始位置.
// dir : false = 反向搜尋.
// 成功 : 當找到符合條件的圖騰時, 傳回 index 與 Stop.
// 失敗 : 傳回 start 與 nil
func (s *Strip) FindMatch(mask uint32, start int, dir bool) (int, *Stop) {
	return s.find(mask, start, dir, true)
}

// FindUnmatch : 尋找不相符的圖騰.
// start : 起始位置.
// dir : false = 反向搜尋.
// 成功 : 當找到符合條件的圖騰時, 傳回 index 與 Stop.
// 失敗 : 傳回 start 與 nil
func (s *Strip) FindUnmatch(mask uint32, start int, dir bool) (int, *Stop) {
	return s.find(mask, start, dir, false)
}

func (s *Strip) find(mask uint32, start int, dir bool, match bool) (int, *Stop) {

	enum := s.Enum(start, dir)
	for nil != enum {
		stop := s.stops[enum.Index()]

		// 檢查錯誤
		if nil == stop || nil == stop.Symbol {
			break
		}
		if match == stop.Symbol.Match(mask) {
			return enum.Index(), stop
		}
		// 列舉下一個
		enum = enum.Next()
	}

	return start, nil
}

func (s *Strip) Enum(index int, dir bool) *enum {

	// 檢查邊界
	if true == s.Bound(index) {
		enum := &enum{index: index, length: s.Length(), step: 1}
		if false == dir {
			enum.step = -1
		}
		return enum
	}

	// 索引超出範圍時 傳回 nil
	return nil
}

func (s *Strip) String() string {
	var str string = "["
	for _, stop := range s.stops {
		str += " " + stop.Symbol.String()
	}
	str += " ]"
	return str
}

type enum struct {
	index  int
	length int
	step   int
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

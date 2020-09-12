package slots

// Symbol :
type Symbol struct {
	id   int    // 圖騰編號
	flag uint32 // 圖騰類型
	str  string // 圖騰描述文字
}

// NewSymbol : 建立 Symbol 物件
func NewSymbol(id int, flag uint32, str string) *Symbol {
	return &Symbol{
		id:   id,
		flag: flag,
		str:  str,
	}
}

// Create : 建立 Symbol 物件
func (s *Symbol) Create(id int, flag uint32, str string) {
	s.id = id
	s.flag = flag
	s.str = str
}

// ID : 圖騰識別
func (s *Symbol) Id() int {
	if nil != s {
		return s.id
	}
	return -1
}

// Flag : 圖騰類型檢查
func (s *Symbol) Flag() uint32 {
	if nil != s {
		return s.flag
	}
	return 0
}

// String : 圖騰文字描述
func (s *Symbol) String() string {
	if nil != s {
		return s.str
	}
	return "nil"
}

// Match : 檢查圖騰是否符合指定的 flag.
// 成立條件 : (symbol.flag & mask) == symbol.flag
func (s *Symbol) Match(mask uint32) bool {
	return nil != s && s.flag == (s.flag&mask)
}

// Equal : 檢查兩個圖騰是否相等 (id 與 flag 必須完全相同)
func (s *Symbol) Equal(s2 *Symbol) bool {
	return (nil != s && nil != s2) && (s.id == s2.id && s.flag == s2.flag)
}

package slotkit

import (
	"fmt"
	"strings"
)

type Symbol interface {

	// ID : 圖騰識別
	Id() int

	// Flag : 圖騰類型
	Flag() uint32

	// String : 圖騰文字描述
	String() string

	// Match : 檢查圖騰是否符合指定的 flag.
	//  成立條件 : (symbol.flag & mask) == symbol.flag
	Match(uint32) bool

	// Equal : 檢查兩個圖騰是否相等 (id 與 flag 必須完全相同)
	Equal(Symbol) bool
}

// symbol :
type symbol struct {
	id   int    // 圖騰編號
	flag uint32 // 圖騰類型
	str  string // 圖騰描述文字
}

// NewSymbol : 建立 Symbol 物件
func NewSymbol(id int, flag uint32, str ...string) Symbol {

	s := strings.Join(str, "")
	if len(s) == 0 {
		s = fmt.Sprintf("%02d", id)
	}

	return &symbol{
		id:   id,
		flag: flag,
		str:  s,
	}
}

func (s *symbol) Id() int {
	return s.id
}

func (s *symbol) Flag() uint32 {
	return s.flag
}

func (s symbol) String() string {
	return s.str
}

func (s *symbol) Match(mask uint32) bool {
	return nil != s && s.flag == (s.flag&mask)
}

func (s *symbol) Equal(s2 Symbol) bool {
	return (nil != s && nil != s2) && (s.id == s2.Id() && s.flag == s2.Flag())
}

// SymbolMap : 圖騰對照表
type SymbolMap map[int]Symbol

func (sm *SymbolMap) Add(symbol Symbol) {
	(*sm)[symbol.Id()] = symbol
}

// Get : 取得圖騰陣列.
func (sm SymbolMap) Get(ids ...int) []Symbol {
	return sm.GetRange(ids, 0, len(ids))
}

// GetRange : 取得圖騰陣列 (起始 ids[start % len(ids)]).
func (sm SymbolMap) GetRange(ids []int, start int, count int) []Symbol {
	symbols := make([]Symbol, count)

	for i := 0; i < count; i++ {
		id := ids[(start+i)%len(ids)]
		if s, ok := sm[id]; ok {
			symbols[i] = s
		}
	}
	return symbols
}

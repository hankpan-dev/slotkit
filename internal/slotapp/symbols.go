package slotapp

import (
	"slotlib/pkg"
)

var Symbols map[string]*slotlib.Symbol = make(map[string]*slotlib.Symbol)

func init() {
	Symbols["1"] = Symbol10
	Symbols["2"] = SymbolJ
	Symbols["3"] = SymbolQ
	Symbols["4"] = SymbolK
	Symbols["5"] = SymbolA
	Symbols["S"] = SymbolScatter
	Symbols["W"] = SymbolWild
}

var (
	NonSymbol     = slotlib.NewSymbol(0, 0x0000, "?")
	Symbol10      = slotlib.NewSymbol(1, 0x0001, "10")
	SymbolJ       = slotlib.NewSymbol(2, 0x0002, "J")
	SymbolQ       = slotlib.NewSymbol(3, 0x0004, "Q")
	SymbolK       = slotlib.NewSymbol(4, 0x0008, "K")
	SymbolA       = slotlib.NewSymbol(5, 0x0010, "A")
	SymbolScatter = slotlib.NewSymbol(9, 0x4000, "Satter")
	SymbolWild    = slotlib.NewSymbol(10, 0x8000, "Wild")
)

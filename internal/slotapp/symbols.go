package slotapp

import (
	"slotlib"
)

var Symbols map[int]*slotlib.Symbol = make(map[int]*slotlib.Symbol)

func init() {
	AddSymbol(Symbol10)
	AddSymbol(Symbol10)
	AddSymbol(SymbolJ)
	AddSymbol(SymbolQ)
	AddSymbol(SymbolK)
	AddSymbol(SymbolA)
	AddSymbol(SymbolScatter)
	AddSymbol(SymbolWild)
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

func AddSymbol(symbol *slotlib.Symbol) {
	Symbols[symbol.Id()] = symbol
}

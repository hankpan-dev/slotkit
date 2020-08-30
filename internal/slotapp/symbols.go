package slotapp

import (
	"slotlib"
)

const normal uint32 = 0x00ff

var Symbols map[int]*slotlib.Symbol = make(map[int]*slotlib.Symbol)

func init() {
	addSymbol(SymbolJ)
	addSymbol(SymbolQ)
	addSymbol(SymbolK)
	addSymbol(SymbolA)
	addSymbol(SymbolStone)
	addSymbol(SymbolGem)
	addSymbol(SymbolSpell)
	addSymbol(SymbolRing)
	addSymbol(SymbolScatter)
	addSymbol(SymbolWild)
}

var (
	NonSymbol     = slotlib.NewSymbol(0, 0x0000, "?")
	SymbolJ       = slotlib.NewSymbol(2, 0x0001, "J")
	SymbolQ       = slotlib.NewSymbol(3, 0x0002, "Q")
	SymbolK       = slotlib.NewSymbol(4, 0x0004, "K")
	SymbolA       = slotlib.NewSymbol(5, 0x0008, "A")
	SymbolStone   = slotlib.NewSymbol(5, 0x0010, "Stone")
	SymbolGem     = slotlib.NewSymbol(5, 0x0020, "Gem")
	SymbolSpell   = slotlib.NewSymbol(5, 0x0040, "Spell")
	SymbolRing    = slotlib.NewSymbol(1, 0x0080, "Ring")
	SymbolScatter = slotlib.NewSymbol(9, 0x4000, "Satter")
	SymbolWild    = slotlib.NewSymbol(10, 0x8000, "Wild")
)

func addSymbol(symbol *slotlib.Symbol) {
	Symbols[symbol.Id()] = symbol
}

package slotapp

import (
	"slotlib/slots"
)

const normal uint32 = 0x00ff

var Symbols map[int]*slots.Symbol = make(map[int]*slots.Symbol)

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
	NonSymbol     = slots.NewSymbol(0, 0x0000, "?")
	SymbolJ       = slots.NewSymbol(2, 0x0001, "J")
	SymbolQ       = slots.NewSymbol(3, 0x0002, "Q")
	SymbolK       = slots.NewSymbol(4, 0x0004, "K")
	SymbolA       = slots.NewSymbol(5, 0x0008, "A")
	SymbolStone   = slots.NewSymbol(5, 0x0010, "Stone")
	SymbolGem     = slots.NewSymbol(5, 0x0020, "Gem")
	SymbolSpell   = slots.NewSymbol(5, 0x0040, "Spell")
	SymbolRing    = slots.NewSymbol(1, 0x0080, "Ring")
	SymbolScatter = slots.NewSymbol(9, 0x4000, "Satter")
	SymbolWild    = slots.NewSymbol(10, 0x8000, "Wild")
)

func addSymbol(symbol *slots.Symbol) {
	Symbols[symbol.Id()] = symbol
}

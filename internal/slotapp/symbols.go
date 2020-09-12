package slotapp

import (
	"slotkit"
)

const normal uint32 = 0x00ff

var Symbols map[int]*slotkit.Symbol = make(map[int]*slotkit.Symbol)

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
	NonSymbol     = slotkit.NewSymbol(0, 0x0000, "?")
	SymbolJ       = slotkit.NewSymbol(2, 0x0001, "J")
	SymbolQ       = slotkit.NewSymbol(3, 0x0002, "Q")
	SymbolK       = slotkit.NewSymbol(4, 0x0004, "K")
	SymbolA       = slotkit.NewSymbol(5, 0x0008, "A")
	SymbolStone   = slotkit.NewSymbol(5, 0x0010, "Stone")
	SymbolGem     = slotkit.NewSymbol(5, 0x0020, "Gem")
	SymbolSpell   = slotkit.NewSymbol(5, 0x0040, "Spell")
	SymbolRing    = slotkit.NewSymbol(1, 0x0080, "Ring")
	SymbolScatter = slotkit.NewSymbol(9, 0x4000, "Satter")
	SymbolWild    = slotkit.NewSymbol(10, 0x8000, "Wild")
)

func addSymbol(symbol *slotkit.Symbol) {
	Symbols[symbol.Id()] = symbol
}

package slotapp

import (
	"slotkit"
)

const normal uint32 = 0x00ff

var Symbols slotkit.SymbolMap = make(slotkit.SymbolMap)

func init() {
	Symbols.Add(SymbolJ)
	Symbols.Add(SymbolQ)
	Symbols.Add(SymbolK)
	Symbols.Add(SymbolA)
	Symbols.Add(SymbolStone)
	Symbols.Add(SymbolGem)
	Symbols.Add(SymbolSpell)
	Symbols.Add(SymbolRing)
	Symbols.Add(SymbolScatter)
	Symbols.Add(SymbolWild)
}

var (
	SymbolWild    = slotkit.NewSymbol(0, 0x8000, "WD")
	SymbolJ       = slotkit.NewSymbol(1, 0x0001, "m1")
	SymbolQ       = slotkit.NewSymbol(2, 0x0002, "m2")
	SymbolK       = slotkit.NewSymbol(3, 0x0004, "m3")
	SymbolA       = slotkit.NewSymbol(4, 0x0008, "m4")
	SymbolStone   = slotkit.NewSymbol(5, 0x0010, "m5")
	SymbolGem     = slotkit.NewSymbol(6, 0x0020, "m6")
	SymbolSpell   = slotkit.NewSymbol(7, 0x0040, "m7")
	SymbolRing    = slotkit.NewSymbol(8, 0x0080, "m8")
	SymbolScatter = slotkit.NewSymbol(9, 0x4000, "SC")
)

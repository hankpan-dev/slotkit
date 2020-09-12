package slotapp

import (
	"slotlib/slots"
)

var Strips *slots.Strip = slots.NewStrip(20)

func init() {
	Strips.SetSymbols(0,
		SymbolJ,
		SymbolK,
		SymbolQ,
		SymbolWild,
		SymbolStone,
		SymbolA,
		SymbolK,
		SymbolQ,
		SymbolRing,
		SymbolQ,
		SymbolStone,
		SymbolGem,
		SymbolSpell,
		SymbolGem,
		SymbolA,
		SymbolWild,
		SymbolSpell,
		SymbolStone,
		SymbolJ,
		SymbolScatter,
	)
}

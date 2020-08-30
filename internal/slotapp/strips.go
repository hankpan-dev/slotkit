package slotapp

import (
	"slotlib"
)

var Strips *slotlib.Strip = slotlib.NewStrip(20)

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

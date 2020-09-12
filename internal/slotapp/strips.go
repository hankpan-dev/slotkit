package slotapp

import (
	"slotkit"
)

var Strips *slotkit.Strip = slotkit.NewStrip(20)

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

package slotapp

import (
	"log"
	"slotlib"
)

func Run() {

	slot := slotlib.NewSlot(3, 3, 3, 3, 3) // 3x5
	slot.Reel(0).SetSymbols(0, SymbolA, SymbolQ, SymbolJ)
	slot.Reel(1).SetSymbols(0, SymbolA, SymbolK, Symbol10)
	slot.Reel(2).SetSymbols(0, SymbolWild, SymbolA, SymbolScatter)
	slot.Reel(3).SetSymbols(0, SymbolQ, SymbolA, SymbolA)
	slot.Reel(4).SetSymbols(0, SymbolK, SymbolScatter, SymbolK)
	log.Println(slot)

	// 定義 payline
	payline := slotlib.NewPayline(0, 3, 6, 9, 12)
	payline.Update(slot)

	// 計算 normal symbol 數量
	var normal uint32 = 0x000FF
	count := slot.CountMatch(normal)
	log.Println("slot has normal symbols :", count)

	// 尋找第一個 normal symbol
	// 計算連線數 (包含 wild symbol)
	dir := false
	idx, stop := payline.FindMatch(normal, 4, dir)
	if nil != stop {
		log.Println("payline - symbol", stop.Symbol, "at index", idx)
		combins := payline.Combinations(stop.Symbol.Flag()|SymbolWild.Flag(), dir)
		log.Println("payline - symbol", stop.Symbol, "combinations", combins)

		// 計算 normal symbol Ways
		ways := 0
		base := 1
		combins = 0
		for i := 0; i < slot.Reels(); i++ {
			reel := slot.Reel(i)
			count := reel.CountMatch(stop.Symbol.Flag() | SymbolWild.Flag())

			if 0 == count {
				break
			}
			// 計算 ways 與 combinations
			ways = base * count
			base = ways
			combins++
		}
		log.Println("payline - symbol", stop.Symbol, ways, "ways, combinations", combins)
	}

	symbol := SymbolWild
	count = slot.CountEqual(symbol)
	log.Println("slot has", count, symbol)

	count = slot.CountMatch(normal)
	log.Println("slot has", count, "normal symbols")

}

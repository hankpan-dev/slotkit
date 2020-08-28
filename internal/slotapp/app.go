package slotapp

import (
	"log"
	slotlib "slotlib/pkg"
)

func Run() {
	slot := slotlib.NewSlot(3, 3, 3, 3, 3) // 3x5
	slot.Reels()[0].SetSymbols(0, SymbolA, SymbolQ, SymbolJ)
	slot.Reels()[1].SetSymbols(0, SymbolA, Symbol10, Symbol10)
	slot.Reels()[2].SetSymbols(0, SymbolWild, SymbolA, SymbolScatter)
	slot.Reels()[3].SetSymbols(0, SymbolQ, SymbolA, SymbolA)
	slot.Reels()[4].SetSymbols(0, SymbolScatter, SymbolK, SymbolK)
	log.Println(slot)

	// 定義 payline
	payline := slotlib.NewPayline(0, 3, 6, 9, 12)

	// 計算 normal symbol 數量
	var normal uint32 = 0x000FF
	count := payline.CountMatch(slot, normal)
	log.Println("total normal symbols :", count)

	// 尋找第一個 normal symbol
	// 計算連線數 (包含 wild symbol)
	idx, pos := payline.FindMatch(slot, normal, 0, true)
	if 0 <= idx && nil != pos {
		combins := payline.Combinations(slot, pos.Symbol.Flag()|SymbolWild.Flag(), true)
		log.Println("symbol", pos.Symbol, "combinations", combins)

		ways := 0
		combins = 0
		for _, reel := range slot.Reels() {
			count := reel.CountMatch(pos.Symbol.Flag() | SymbolWild.Flag())

			if 0 == count {
				break
			}
			if 0 == ways {
				ways = 1
			}

			ways *= count
			combins++
		}
		log.Println("symbol", pos.Symbol, ways, "ways, combinations", combins)
	}

	symbol := SymbolWild
	count = slot.CountEqual(symbol)
	log.Println("slot has", count, symbol)

	count = slot.CountMatch(normal)
	log.Println("slot has", count, "normal symbols")


}
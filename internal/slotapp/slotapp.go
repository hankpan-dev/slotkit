package slotapp

import (
	"fmt"
	"log"
	"math/rand"
	"slotkit"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func Run() {
	slot := slotkit.NewSlot(5, 3) // 5x3
	spin(slot)
	checkPaylines(slot)
	checkWays(slot)
}

func spin(slot *slotkit.Slot) {

	for i, reel := range slot.Reels() {
		strip := Strips[i]

		// 設定滾輪圖騰, 等同於
		// reel.SetSymbols(0, Symbols.GetRange(strip, idx, len(reel))...)
		idx := rnd.Intn(len(strip))
		for _, row := range reel {
			id := strip[idx%len(strip)]
			row.Symbol = Symbols[id]
			idx++
		}
	}

	log.Println("spin result :", slot)
}

func checkPaylines(slot *slotkit.Slot) {

	log.Println("check paylines...")

	win := 0
	for i, payline := range Paylines {
		// 取得 payline 對應的圖騰
		line := payline.FetchSymbols(slot)

		var symbol slotkit.Symbol
		if n := line.FindMatch(normal, 0); n >= 0 {
			symbol = line[n].Symbol
		}

		if symbol == nil {
			continue
		}

		score := 0
		combins := line.Combinations(symbol.Flag() | SymbolWild.Flag())
		if 0 < combins {
			score = Paytable[symbol.Id()-1][combins-1]
		}

		details := ""
		if 0 < score {
			details = fmt.Sprintf("'%s' combinations %d, win %d.", symbol, combins, score)
		}

		log.Println("payline", i, line, details)
		win += score
	}

	log.Println("payline total win", win)

	sct := slot.Cells().CountID(SymbolScatter.Id())
	log.Println("total", sct, "scatters")
}

func checkWays(slot *slotkit.Slot) {
	log.Println("check ways...")

	// 略過 scatter & wild
	mask := SymbolScatter.Flag() | SymbolWild.Flag()

	// 列舉第一輪
	for _, pos := range slot.Reel(0) {
		// 計算 normal symbol Ways
		base := 1
		ways := 0
		combins := 0

		// 取得第一輪的圖騰
		symbol := pos.Symbol
		if symbol.Match(mask) {
			continue
		}
		mask |= symbol.Flag()

		for _, reel := range slot.Reels() {
			count := reel.CountMatch(symbol.Flag() | SymbolWild.Flag())
			if count == 0 {
				break
			}

			// 計算 ways 與 combinations
			ways = base * count
			base = ways
			combins++
		}

		detial := fmt.Sprintf("'%s' has %v ways, %v combinations", symbol, ways, combins)
		log.Println(detial)
	}
}

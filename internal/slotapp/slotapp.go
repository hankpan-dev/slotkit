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
	slot := slotkit.NewSlot(3, 3, 3, 3, 3) // 3x5
	spin(slot)
	checkPaylines(slot)
	checkWays(slot)
}

func spin(slot *slotkit.Slot) {

	for index := 0; index < slot.Reels(); index++ {
		reel := slot.Reel(index)

		// 取得亂數, 指向彩帶的位置
		idx := rnd.Intn(Strips.Length())

		for iter := reel.Iterator(0, true); iter.HasNext(); iter.Next() {
			iter.Current().Symbol = Strips.Cycler(idx).Symbol
			idx++
		}
	}

	log.Println("spin result :", slot)
}

func checkPaylines(slot *slotkit.Slot) {
	log.Println("check paylines...")
	win := 0
	for i, payline := range Paylines {
		// 設置 payline 圖騰
		payline.SetSymbols(slot)
		_, stop := payline.FindMatch(normal, 0, true)

		combins := 0
		if nil != stop {
			combins = payline.Combinations(stop.Symbol.Flag()|SymbolWild.Flag(), true)
		}

		score := 0
		if 0 < combins {
			score = Paytable[stop.Symbol.Id()-1][combins-1]
		}
		detail := ""
		if 0 < score {
			detail = fmt.Sprintf("has %v '%s' combinations, win %v.", combins, stop.Symbol, score)
		}
		log.Println(i, ":", payline, detail)
		win += score
	}

	log.Println("payline total win", win)

	sct := slot.CountID(SymbolScatter.Id())
	if 3 <= sct {
		log.Println("total", sct, "scatters")
	}
}

func checkWays(slot *slotkit.Slot) {
	log.Println("check ways...")

	// 略過 scatter & wild
	mask := SymbolScatter.Flag() | SymbolWild.Flag()

	// 列舉第一輪
	for iter := slot.Reel(0).Iterator(0, true); iter.HasNext(); iter.Next() {
		// 計算 normal symbol Ways
		base := 1
		ways := 0
		combins := 0

		// 取得第一輪上的圖騰
		symbol := iter.Current().Symbol
		if true == symbol.Match(mask) {
			continue
		}
		mask |= symbol.Flag()

		for i := 0; i < slot.Reels(); i++ {
			reel := slot.Reel(i)
			count := reel.CountMatch(symbol.Flag() | SymbolWild.Flag())
			if 0 == count {
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

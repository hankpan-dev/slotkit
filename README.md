# slotkit
Slot machine library

### 範例 : 建立一個 3x5 的滾輪
#### 滾輪位置編號<BR>
| R0 | R1 | R2 | R3 | R4 |
|----|----|----|----|----|
| 00 | 03 | 06 | 09 | 12 |
| 01 | 04 | 07 | 10 | 13 |
| 02 | 05 | 08 | 11 | 14 |

```
slot := slotkit.NewSlot(3, 3, 3, 3, 3)
```



### 範例 : Payline
#### 建立 Payline 樣式為 - 位置(01, 04, 07, 10, 13)

```
// 建立 Payline 樣式為 - 01, 04, 07, 10, 13 號位置
payline := slotkit.NewPayline(1, 4, 7, 10, 13)

// 更新 slot 圖騰到 payline 上
payline.SetSymbols(slot)

// 尋找符合條件的圖騰 (Mask=0x00ff, 尋找 type = 0x0000 ~ 0x00ff 之間的任意一個圖騰)
start, stop := payline.FindMatch(0xff, 0, true)

// 計算連線長度
combins := payline.Combinations(stop.Symbol.Flag()|SymbolWild.Flag(), true)


```

### 範例 : 建立 Symbols <BR>
```
var (
  NonSymbol     = slotkit.NewSymbol(0, 0x0000, "?")       // Normal symbol
  SymbolJ       = slotkit.NewSymbol(1, 0x0001, "J")       // Normal symbol
  SymbolQ       = slotkit.NewSymbol(2, 0x0002, "Q")       // Normal symbol
  SymbolK       = slotkit.NewSymbol(3, 0x0004, "K")       // Normal symbol
  SymbolA       = slotkit.NewSymbol(4, 0x0008, "A")       // Normal symbol
  SymbolScatter = slotkit.NewSymbol(9, 0x4000, "Satter")  // Scatter
  SymbolWild    = slotkit.NewSymbol(10, 0x8000, "Wild")   // Wild
)
```

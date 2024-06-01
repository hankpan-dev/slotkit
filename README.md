# slotkit
> Slot machine library

### 範例 : 建立一個 5x3 的滾輪
```
// 創建 slot 物件 (5x3)
slot := slotkit.NewSlot(5, 3)
slot := slotkit.NewSlotFlex(3, 3, 3, 3, 3)

// 設定 slot 上的圖騰
slot.Cells()[0].Symbol = SymbolA
slot.Reel(0)[0].Symbol = SymbolA

// 計算 slot 相同 id 的圖騰數量
slot.CountID(SymbolWild.Id)

// 取得滾輪
reel := slot.Reel(0)

// 設定滾輪上的圖騰
reel.SetSymbols(0, SymbolA, SymbolWild, ...)

// 計算滾輪上相同 id 的圖騰數量
reel.CountID(SymbolWild.Id)
```
#### < Slot 排列規則 >
| R0 | R1 | R2 | R3 | R4 |
|----|----|----|----|----|
| 00 | 03 | 06 | 09 | 12 |
| 01 | 04 | 07 | 10 | 13 |
| 02 | 05 | 08 | 11 | 14 |
```
// slot 最後一個位置, 相當於最後一個滾輪的最後位置, 以下幾個寫法結果相同
slot.SetSymbol(14, Symbol) 
slot.Reel(4)[2].Symbol = Symbol
```

---
### 範例 : Payline
##### 建立 Payline 樣式為 - 位置(0, 0, 1, 2, 2)

```
// 建立 Payline 樣式為 - Slot 的 00, 03, 07, 11, 14 號位置
payline := slotkit.NewPayline(0, 0, 1, 2, 2)

// 取得 payline 對應於 slot 上的圖騰到
line := payline.FetchSymbols(slot)

// 取得反向圖騰
line := payline.FetchSymbols(slot).Reverse()

// 尋找符合條件的圖騰 (Mask=0x00ff, 尋找 type=0x0000 ~ 0x00ff 之間的任意一個圖騰)
index := line.FindMatch(0xff, 0)

// 計算連線長度
combins := line.Combinations(SymbolA.Flag()|SymbolWild.Flag())

// 計算 Payline 上相同 id 的圖騰數量
count := line.CountID((SymbolWild.Id))
```

---
### 範例 : Symbol
##### Symbol 包含 Id, Type, Name
Id : 定義圖騰唯一編號
Type : 定義圖騰屬性, 例如 normal, scatter, wild...
Name : 圖騰名稱

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
以上圖騰範例  
可用使用 CountMatch(0x000f) 計算 Normal symbols 數量  
可以使用 CountMatch(0x800f) 計算 Normal & Wild symbol 總和

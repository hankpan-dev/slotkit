package slotlib

// Slot :
type Slot struct {
	length int
	reels  []*Reel
}

// NewSlot : 建立 slot
func NewSlot(reelsize ...int) *Slot {
	s := &Slot{}
	s.Create(reelsize...)
	return s
}

// Create :
func (s *Slot) Create(reelsize ...int) {
	s.length = 0
	s.reels = make([]*Reel, len(reelsize))

	for i, size := range reelsize {
		s.reels[i] = NewReel(size)
		s.length += size
	}
}

// Reels : 取得滾輪陣列
func (s *Slot) Reels() []*Reel {
	return s.reels
}

// TotalReels : 取得滾輪數量
func (s *Slot) TotalReels() int {
	return len(s.reels)
}

// Length : 取得 Pos 數量
func (s *Slot) Length() int {
	return s.length
}

// Index :
func (s *Slot) Index(index int) *Pos {
	// 檢查 index 是否超出範圍
	if 0 <= index && index < s.length {
		for _, reel := range s.reels {
			
			item := reel.Index(index)
			
			// 檢查是否取得 item
			if nil == item {
				index -= reel.Length()
			} else {
				return item
			}
		}
	}
	return nil
}

// CountId : 計算相同 id 的數量
func (s *Slot) CountId(id int) int {
	count := 0
	for _, reel := range s.reels {
		count += reel.CountId(id)
	}
	return count
}

// CountMatch : 計算相符圖騰 flag 的數量
func (s *Slot) CountMatch(mask uint32) int {
	count := 0
	for _, reel := range s.reels {
		count += reel.CountMatch(mask)
	}
	return count
}

// CountEqual : 計算相同圖騰數量 (id 與 flag 皆相等)
func (s *Slot) CountEqual(symbol *Symbol) int {
	count := 0
	for _, reel := range s.reels {
		count += reel.CountEqual(symbol)
	}
	return count
}

// SwitchReel : 將 slot 上的兩個 reel 交換
func (s *Slot) SwitchReel(a int, b int) bool {
	// 檢查交換位置是否超出範圍
	if (0 <= a && len(s.reels) > a) && (0 <= b && len(s.reels) > b) {
		if a != b {
			reel := s.reels[a]
			s.reels[a] = s.reels[b]
			s.reels[b] = reel
		}
		return true
	}
	return false
}

func (s *Slot) String() string {
	var str string = "["
	for _, reel := range s.reels {
		str += " " + reel.String()
	}
	str += " ]"
	return str
}

package slotkit

// Slot :
type Slot struct {
	Strip
	reels []*Strip
}

// NewSlot : 建立 slot
func NewSlot(reelsize ...int) *Slot {
	s := &Slot{}
	s.Create(reelsize...)
	return s
}

// Create :
func (s *Slot) Create(reelsize ...int) {
	length := 0
	for _, size := range reelsize {
		length += size
	}

	s.Strip.Create(length)
	s.reels = make([]*Strip, len(reelsize))
	index := 0
	for i := range s.reels {
		length = reelsize[i]
		s.reels[i] = s.Slice(index, length)
		index += length
	}
}

// Reel : 取得滾輪陣列
func (s *Slot) Reel(index int) *Strip {
	return s.reels[index]
}

// Reels : 取得滾輪數量
func (s *Slot) Reels() int {
	return len(s.reels)
}

func (s *Slot) String() string {
	var str string = "["
	for _, reel := range s.reels {
		str += " " + reel.String()
	}
	str += " ]"
	return str
}

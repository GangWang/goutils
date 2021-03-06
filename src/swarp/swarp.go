package swrap

type SWrap []byte

func New(a []byte) SWrap {
	return SWrap(a)
}

func (sw *SWrap) Bytes() []byte {
	return []byte(*sw)
}

func (sw *SWrap) Len() int {
	return len(*sw)
}

func (sw *SWrap) Add(a byte) {
	*sw = append(*sw, a)
}

func (sw *SWrap) Merge(a []byte) {
	s := *sw
	l := len(s) + len(a)
	ss := make([]byte, l, l)
	copy(ss[0:], s[:])
	copy(ss[len(s):], a)
	*sw = ss
}

func (sw *SWrap) Delete(i int) {
	s := *sw
	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0 // GC
	*sw = s[:len(s)-1]
}

func (sw *SWrap) Compare(b []byte) bool {
	s := *sw
	if len(s) != len(b) {
		return false
	}

	for i, v := range b {
		if s[i] != v {
			return false
		}
	}
	return true
}

func (sw *SWrap) Push(b byte) {
	*sw = append(*sw, b)
}

func (sw *SWrap) Pop() byte {
	s := *sw
	last := s[len(s)-1]
	s[len(s)-1] = 0 // GC
	*sw = s[:len(s)-1]
	return last
}

func (sw *SWrap) Shift(b byte) {
	s := *sw
	l := len(s) + 1
	ss := make([]byte, l, l)
	ss[0] = b
	copy(ss[1:], s[:])
	*sw = ss
}

func (sw *SWrap) UnShift() byte {
	s := *sw
	top := s[0]
	s[0] = 0 // GC
	*sw = s[1:]
	return top
}

func (sw *SWrap) Replace(i int, b byte) {
	s := *sw
	over := i - len(s)
	if over > -1 {
		ss := make([]byte, i+1)
		copy(ss[0:], s[:])
		s = ss
	}
	s[i] = b
	*sw = s
}

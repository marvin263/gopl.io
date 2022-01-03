package intset

// OtherMethod reports whether the set contains the non-negative value x.
func (s *IntSet) OtherMethod(x int) bool {
	word, bit := x/64, uint(x%64)
	return uint(word) == bit
}

package chezmoi

type BiDiMap[K comparable] struct {
	forward map[K]K
	reverse map[K]K
}

func NewBiDiMap[K comparable](n int) *BiDiMap[K] {
	return &BiDiMap[K]{
		forward: make(map[K]K, n),
		reverse: make(map[K]K, n),
	}
}

func NewBiDiMapFromMap[M ~map[K]K, K comparable](m M) (*BiDiMap[K], bool) {
	biDiMap := NewBiDiMap[K](len(m))
	for forwardKey, reverseKey := range m {
		if ok := biDiMap.Insert(forwardKey, reverseKey); !ok {
			return nil, false
		}
	}
	return biDiMap, true
}

func (b *BiDiMap[K]) Contains(k K) bool {
	_, forwardOK := b.forward[k]
	_, reverseOK := b.reverse[k]
	return forwardOK || reverseOK
}

func (b *BiDiMap[K]) Insert(forwardKey, reverseKey K) bool {
	if b.Contains(forwardKey) || b.Contains(reverseKey) {
		return false
	}
	b.forward[forwardKey] = reverseKey
	b.reverse[reverseKey] = forwardKey
	return true
}

func (b *BiDiMap[K]) LookupForward(forwardKey K) (K, bool) {
	reverseKey, ok := b.forward[forwardKey]
	return reverseKey, ok
}

func (b *BiDiMap[K]) LookupReverse(reverseKey K) (K, bool) {
	forwardKey, ok := b.reverse[reverseKey]
	return forwardKey, ok
}

func (b *BiDiMap[K]) Verify() bool {
	if len(b.forward) != len(b.reverse) {
		return false
	}
	for forwardKey, reverseKey := range b.forward {
		if key, ok := b.reverse[reverseKey]; !ok || key != forwardKey {
			return false
		}
	}
	return true
}

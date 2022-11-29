package chezmoi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBiMap(t *testing.T) {
	m := NewBiDiMap[int](0)
	assert.True(t, m.Insert(1, -1))
	assert.True(t, m.Contains(1))
	assert.True(t, m.Contains(-1))
	assert.False(t, m.Contains(0))
	assert.True(t, m.Verify())

	reverseKey, ok := m.LookupForward(1)
	assert.True(t, ok)
	assert.Equal(t, -1, reverseKey)

	forwardKey, ok := m.LookupReverse(-1)
	assert.True(t, ok)
	assert.Equal(t, 1, forwardKey)
	assert.False(t, m.Insert(1, 0))
}

func TestNewBiDiMap(t *testing.T) {
	m, ok := NewBiDiMapFromMap(map[string]string{
		"forward": "reverse",
	})
	assert.True(t, ok)
	assert.True(t, m.Verify())

	reverseKey, ok := m.LookupForward("forward")
	assert.True(t, ok)
	assert.Equal(t, "reverse", reverseKey)

	forwardKey, ok := m.LookupReverse("reverse")
	assert.True(t, ok)
	assert.Equal(t, "forward", forwardKey)
}

func TestNewBiDiMapError(t *testing.T) {
	m, ok := NewBiDiMapFromMap(map[bool]bool{
		true:  true,
		false: true,
	})
	assert.False(t, ok)
	assert.Nil(t, m)
}

func TestNewBiDiMapVerify(t *testing.T) {
	assert.False(t, (&BiDiMap[int]{
		forward: map[int]int{
			0: 1,
		},
	}).Verify())
	assert.False(t, (&BiDiMap[int]{
		forward: map[int]int{
			0: 0,
		},
		reverse: map[int]int{
			1: 1,
		},
	}).Verify())
}

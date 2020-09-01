package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMap(t *testing.T) {
	m := make(map[string]int)
	m["e1m1"] = 123

	mm := map[string]int{
		"e3m1": -123,
	}
	mm["e3m1"] = 0x1233
}

func TestKeyHitKeyMiss(t *testing.T) {
	m := make(map[string]int)
	m["e1m1"] = 123

	// key miss results in a zero-value
	assert.Equal(t, m["e1m2"], 0)

	if v, ok := m["e2m2"]; v == 0 {
		assert.Equal(t, ok, false)
	}
}

func TestIterateKeyValue(t *testing.T) {
	m := make(map[string]int)
	m["e1m1"] = 123
	for k, v := range m {
		assert.Equal(t, "e1m1", k)
		assert.Equal(t, 123, v)
	}
}

func TestDeleteKey(t *testing.T) {
	m := make(map[string]int)
	m["e1m1"] = 123
	delete(m, "e1m1")
	assert.Equal(t, 0, m["e1m1"])
}

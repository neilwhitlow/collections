package linkedhashmap_test

import (
	"testing"

	lhm "github.com/neilwhitlow/collections/linkedhashmap"
	"github.com/stretchr/testify/assert"
)

func TestPut(t *testing.T) {
	t.Run("", func(t *testing.T) {
		m := lhm.New[string, string]()
		m.Put("a", "Alpha")
		m.Put("z", "Zeta")
		value := m.Get("a")
		assert.Equal(t, "Alpha", value, "Alpha should exist at key a")

		value = m.Get("z")
		assert.Equal(t, "Zeta", value, "Zeta should exist at key z")
	})
}

func TestPutReplace(t *testing.T) {
	t.Run("", func(t *testing.T) {
		m := lhm.New[string, string]()
		m.Put("a", "Alpha")
		value := m.Get("a")
		assert.Equal(t, "Alpha", value, "Alpha should exist at key a")

		priorValue, exists := m.Put("a", "Agammemnon")
		assert.True(t, exists, "Key a should already exist")
		assert.Equal(t, "Alpha", priorValue, "Alpha should be the prior value at key a")

		value = m.Get("a")
		assert.Equal(t, "Agammemnon", value, "Agammemnon should now exist at key a")
	})
}

func TestGet(t *testing.T) {
	t.Run("", func(t *testing.T) {
		m := lhm.New[string, string]()
		m.Put("d", "Delta")
		m.Put("e", "Epsilon")
		value := m.Get("d")
		assert.Equal(t, "Delta", value, "Delta should exist at key d")

		value = m.Get("e")
		assert.Equal(t, "Epsilon", value, "Epsilon should exist at key e")
	})
}

func TestGetNotExists(t *testing.T) {
	t.Run("", func(t *testing.T) {
		m := lhm.New[string, string]()
		m.Put("d", "Delta")
		value := m.Get("e")
		assert.Equal(t, "", value, "Key e should not exist")
	})
}

func TestFirst(t *testing.T) {
	t.Run("", func(t *testing.T) {
		m := lhm.New[string, string]()
		m.Put("d", "Delta")
		m.Put("a", "Alpha")

		kvp := m.First()
		assert.Equal(t, "d", kvp.Key, "key should be d")
		assert.Equal(t, "Delta", kvp.Value, "Value should be Delta")
	})
}

func TestFirstNil(t *testing.T) {
	t.Run("", func(t *testing.T) {
		m := lhm.New[string, string]()
		kvp := m.First()
		assert.Nil(t, kvp, "Should be nil")
	})
}

func TestNext(t *testing.T) {
	t.Run("", func(t *testing.T) {
		m := lhm.New[string, string]()
		m.Put("d", "Delta")
		m.Put("a", "Alpha")

		kvp := m.First().Next()
		assert.Equal(t, "a", kvp.Key, "key should be a")
		assert.Equal(t, "Alpha", kvp.Value, "Value should be Alpha")
	})
}

func TestNextNil(t *testing.T) {
	t.Run("", func(t *testing.T) {
		m := lhm.New[string, string]()
		m.Put("d", "Delta")
		m.Put("e", "Epsilon")
		kvp := m.Last().Next()
		assert.Nil(t, kvp, "Should be nil")
	})
}

func TestLast(t *testing.T) {
	t.Run("", func(t *testing.T) {
		m := lhm.New[string, string]()
		m.Put("a", "Alpha")
		m.Put("d", "Delta")

		kvp := m.Last()
		assert.Equal(t, "d", kvp.Key, "key should be d")
		assert.Equal(t, "Delta", kvp.Value, "Value should be Delta")
	})
}

func TestLastNil(t *testing.T) {
	t.Run("", func(t *testing.T) {
		m := lhm.New[string, string]()
		kvp := m.Last()
		assert.Nil(t, kvp, "Should be nil")
	})
}

func TestPrev(t *testing.T) {
	t.Run("", func(t *testing.T) {
		m := lhm.New[string, string]()
		m.Put("a", "Alpha")
		m.Put("d", "Delta")

		kvp := m.Last().Prev()
		assert.Equal(t, "a", kvp.Key, "key should be a")
		assert.Equal(t, "Alpha", kvp.Value, "Value should be Alpha")
	})
}

func TestPrevNil(t *testing.T) {
	t.Run("", func(t *testing.T) {
		m := lhm.New[string, string]()
		m.Put("d", "Delta")
		kvp := m.First().Prev()
		assert.Nil(t, kvp, "Should be nil")
	})
}

func TestNewWithCapacity(t *testing.T) {
	t.Run("", func(t *testing.T) {
		m := lhm.New[string, string](5)
		m.Put("a", "Alpha")
		m.Put("z", "Zeta")
		value := m.Get("a")
		assert.Equal(t, "Alpha", value, "Alpha should exist at key a")

		value = m.Get("z")
		assert.Equal(t, "Zeta", value, "Zeta should exist at key z")
	})
}

func TestKeys(t *testing.T) {
	t.Run("", func(t *testing.T) {
		m := lhm.New[string, string]()
		m.Put("a", "Alpha")
		m.Put("z", "Zeta")

		assert.Len(t, m.Keys(), 2, "length should be 2")
	})
}

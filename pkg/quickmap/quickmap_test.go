package quickmap

import (
	"strconv"
	"testing"
)

func TestQuickMap(t *testing.T) {
	m := New()

	// Test Insert and Get
	t.Run("Insert and Get", func(t *testing.T) {
		m.Insert("key1", "value1")
		value, exists := m.Get("key1")
		if !exists {
			t.Errorf("Get(\"key1\") returned false, expected true")
		}
		if value != "value1" {
			t.Errorf("Get(\"key1\" = %v, expected \"value1\"", value)
		}
	})

	// Test overwriting existing key
	t.Run("Overwrite existing key", func(t *testing.T) {
		m.Insert("key1", "new_value")
		value, _ := m.Get("key1")
		if value != "new_value" {
			t.Errorf("Get(\"key1\" = %v, expected \"new_value\"", value)
		}
	})

	// Test Delete
	t.Run("Delete", func(t *testing.T) {
		m.Delete("key1")
		_, exists := m.Get("key1")
		if exists {
			t.Errorf("Get(\"key1\") returned true after deletion, expected false")
		}
	})

	// Test resize
	t.Run("Resize", func(t *testing.T) {
		initialCap := len(m.buckets)
		for i := 0; i < 100; i++ {
			m.Insert(strconv.Itoa(i), i)
		}
		if len(m.buckets) <= initialCap {
			t.Errorf("Expected resize to occur, but capacity remained at %d", len(m.buckets))
		}
	})
}

func BenchmarkQuickMap(b *testing.B) {
	m := New()

	b.Run("Insert", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Insert(strconv.Itoa(i), i)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Get(strconv.Itoa(i % 1000))
		}
	})

	b.Run("Delete", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Delete(strconv.Itoa(i % 1000))
		}
	})
}

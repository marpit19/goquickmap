package quickmap

import (
	"fmt"
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

	// Test NewWithCapacity
	t.Run("NewWithCapacity", func(t *testing.T) {
		m := NewWithCapacity(100)
		if len(m.buckets) < 100 {
			t.Errorf("NewWithCapacity(100) created a map with capacity %d, expected at least 100", len(m.buckets))
		}
	})

	// Test InsertMany
	t.Run("InsertMany", func(t *testing.T) {
		m := New()
		pairs := map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}
		m.InsertMany(pairs)
		if m.Size() != 3 {
			t.Errorf("After InsertMany, Size() = %d, expected 3", m.Size())
		}
		for k, v := range pairs {
			if val, exists := m.Get(k); !exists || val != v {
				t.Errorf("After InsertMany, Get(%s) = %v, %t; expected %v, true", k, val, exists, v)
			}
		}
	})

	// Test DeleteMany
	t.Run("DeleteMany", func(t *testing.T) {
		m := New()
		pairs := map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}
		m.InsertMany(pairs)
		m.DeleteMany([]string{"key1", "key3"})
		if m.Size() != 1 {
			t.Errorf("After DeleteMany, Size() = %d, expected 1", m.Size())
		}
		if _, exists := m.Get("key1"); exists {
			t.Errorf("After DeleteMany, key1 still exists")
		}
		if _, exists := m.Get("key3"); exists {
			t.Errorf("After DeleteMany, key3 still exists")
		}
		if _, exists := m.Get("key2"); !exists {
			t.Errorf("After DeleteMany, key2 no longer exists")
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

	// Benchmark InsertMany
	b.Run("InsertMany", func(b *testing.B) {
		m := New()
		for i := 0; i < b.N; i++ {
			pairs := make(map[string]interface{})
			for j := 0; j < 1000; j++ {
				pairs[strconv.Itoa(j)] = j
			}
			m.InsertMany(pairs)
		}
	})

	// Benchmark DeleteMany
	b.Run("DeleteMany", func(b *testing.B) {
		m := New()
		pairs := make(map[string]interface{})
		keys := make([]string, 1000)
		for i := 0; i < 1000; i++ {
			key := strconv.Itoa(i)
			pairs[key] = i
			keys[i] = key
		}
		m.InsertMany(pairs)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.DeleteMany(keys)
		}
	})

	// Benchmark with different initial capacities
    initialCapacities := []int{16, 128, 1024, 8192}
    for _, cap := range initialCapacities {
        b.Run(fmt.Sprintf("InsertWithCapacity%d", cap), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                m := NewWithCapacity(cap)
                for j := 0; j < 1000; j++ {
                    m.Insert(strconv.Itoa(j), j)
                }
            }
        })
    }
}

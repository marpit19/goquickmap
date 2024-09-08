package quickdict

import (
	"fmt"
	"strconv"
	"testing"
)

func TestQuickDict(t *testing.T) {
	d := New()

	// Test Set and Get
	t.Run("Set and Get", func(t *testing.T) {
		d.Set("key1", "value1")
		value, exists := d.Get("key1")
		if !exists {
			t.Errorf("Get(\"key1\") returned false, expected true")
		}
		if value != "value1" {
			t.Errorf("Get(\"key1\") = %v, expected \"value1\"", value)
		}
	})

	// Test Delete
	t.Run("Delete", func(t *testing.T) {
		d.Delete("key1")
		_, exists := d.Get("key1")
		if exists {
			t.Errorf("Get(\"key1\") returned true after deletion, expected false")
		}
	})

	// Test Size, Keys, and Values
	t.Run("Size, Keys, and Values", func(t *testing.T) {
		d.Set("key2", 2)
		d.Set("key3", 3)
		if d.Size() != 2 {
			t.Errorf("Size() = %d, expected 2", d.Size())
		}
		keys := d.Keys()
		if len(keys) != 2 {
			t.Errorf("len(Keys()) = %d, expected 2", len(keys))
		}
		values := d.Values()
		if len(values) != 2 {
			t.Errorf("len(Values()) = %d, expected 2", len(values))
		}
	})

	// Test NewWithCapacity
	t.Run("NewWithCapacity", func(t *testing.T) {
		d := NewWithCapacity(100)
		if d.Size() != 0 {
			t.Errorf("NewWithCapacity(100) created a dict with size %d, expected 0", d.Size())
		}
		// Add 100 elements to ensure it doesn't resize immediately
		for i := 0; i < 100; i++ {
			d.Set(strconv.Itoa(i), i)
		}
		if d.Size() != 100 {
			t.Errorf("After adding 100 elements, Size() = %d, expected 100", d.Size())
		}
	})

	// Test SetMany
	t.Run("SetMany", func(t *testing.T) {
		d := New()
		pairs := map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}
		d.SetMany(pairs)
		if d.Size() != 3 {
			t.Errorf("After SetMany, Size() = %d, expected 3", d.Size())
		}
		for k, v := range pairs {
			if val, exists := d.Get(k); !exists || val != v {
				t.Errorf("After SetMany, Get(%s) = %v, %t; expected %v, true", k, val, exists, v)
			}
		}
	})

	// Test DeleteMany
	t.Run("DeleteMany", func(t *testing.T) {
		d := New()
		pairs := map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
			"key4": "value4",
		}
		d.SetMany(pairs)
		d.DeleteMany([]string{"key1", "key3"})
		if d.Size() != 2 {
			t.Errorf("After DeleteMany, Size() = %d, expected 2", d.Size())
		}
		if _, exists := d.Get("key1"); exists {
			t.Errorf("After DeleteMany, key1 still exists")
		}
		if _, exists := d.Get("key3"); exists {
			t.Errorf("After DeleteMany, key3 still exists")
		}
		if _, exists := d.Get("key2"); !exists {
			t.Errorf("After DeleteMany, key2 no longer exists")
		}
		if _, exists := d.Get("key4"); !exists {
			t.Errorf("After DeleteMany, key4 no longer exists")
		}
	})
}

func BenchmarkQuickDict(b *testing.B) {
	d := New()

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d.Set(strconv.Itoa(i), i)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d.Get(strconv.Itoa(i % 1000))
		}
	})

	b.Run("Delete", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d.Delete(strconv.Itoa(i % 1000))
		}
	})

	// Benchmark SetMany
    b.Run("SetMany", func(b *testing.B) {
        d := New()
        for i := 0; i < b.N; i++ {
            pairs := make(map[string]interface{})
            for j := 0; j < 1000; j++ {
                pairs[strconv.Itoa(j)] = j
            }
            d.SetMany(pairs)
        }
    })

    // Benchmark DeleteMany
    b.Run("DeleteMany", func(b *testing.B) {
        d := New()
        pairs := make(map[string]interface{})
        keys := make([]string, 1000)
        for i := 0; i < 1000; i++ {
            key := strconv.Itoa(i)
            pairs[key] = i
            keys[i] = key
        }
        d.SetMany(pairs)
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
            d.DeleteMany(keys)
        }
    })

    // Benchmark with different initial capacities
    initialCapacities := []int{16, 128, 1024, 8192}
    for _, cap := range initialCapacities {
        b.Run(fmt.Sprintf("SetWithCapacity%d", cap), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                d := NewWithCapacity(cap)
                for j := 0; j < 1000; j++ {
                    d.Set(strconv.Itoa(j), j)
                }
            }
        })
    }
}

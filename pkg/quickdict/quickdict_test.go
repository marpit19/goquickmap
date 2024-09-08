package quickdict

import (
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
}

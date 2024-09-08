package quickset

import (
	"strconv"
	"testing"
)

func TestQuickSet(t *testing.T) {
    s := New()

    // Test Add and Contains
    t.Run("Add and Contains", func(t *testing.T) {
        s.Add("item1")
        if !s.Contains("item1") {
            t.Errorf("Contains(\"item1\") returned false, expected true")
        }
    })

    // Test Remove
    t.Run("Remove", func(t *testing.T) {
        s.Remove("item1")
        if s.Contains("item1") {
            t.Errorf("Contains(\"item1\") returned true after removal, expected false")
        }
    })

    // Test Size and Elements
    t.Run("Size and Elements", func(t *testing.T) {
        s.Add("item2")
        s.Add("item3")
        if s.Size() != 2 {
            t.Errorf("Size() = %d, expected 2", s.Size())
        }
        elements := s.Elements()
        if len(elements) != 2 {
            t.Errorf("len(Elements()) = %d, expected 2", len(elements))
        }
    })
}

func BenchmarkQuickSet(b *testing.B) {
    s := New()

    b.Run("Add", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s.Add(strconv.Itoa(i))
        }
    })

    b.Run("Contains", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s.Contains(strconv.Itoa(i % 1000))
        }
    })

    b.Run("Remove", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s.Remove(strconv.Itoa(i % 1000))
        }
    })
}

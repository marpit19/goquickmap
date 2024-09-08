package quickset

import (
	"fmt"
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

	// Test NewWithCapacity
	t.Run("NewWithCapacity", func(t *testing.T) {
		s := NewWithCapacity(100)
		if s.Size() != 0 {
			t.Errorf("NewWithCapacity(100) created a set with size %d, expected 0", s.Size())
		}
		// Add 100 elements to ensure it doesn't resize immediately
		for i := 0; i < 100; i++ {
			s.Add(strconv.Itoa(i))
		}
		if s.Size() != 100 {
			t.Errorf("After adding 100 elements, Size() = %d, expected 100", s.Size())
		}
	})

	// Test AddMany
	t.Run("AddMany", func(t *testing.T) {
		s := New()
		elements := []string{"elem1", "elem2", "elem3"}
		s.AddMany(elements)
		if s.Size() != 3 {
			t.Errorf("After AddMany, Size() = %d, expected 3", s.Size())
		}
		for _, elem := range elements {
			if !s.Contains(elem) {
				t.Errorf("After AddMany, set does not contain %s", elem)
			}
		}
	})

	// Test RemoveMany
	t.Run("RemoveMany", func(t *testing.T) {
		s := New()
		elements := []string{"elem1", "elem2", "elem3", "elem4"}
		s.AddMany(elements)
		s.RemoveMany([]string{"elem1", "elem3"})
		if s.Size() != 2 {
			t.Errorf("After RemoveMany, Size() = %d, expected 2", s.Size())
		}
		if s.Contains("elem1") || s.Contains("elem3") {
			t.Errorf("After RemoveMany, set still contains removed elements")
		}
		if !s.Contains("elem2") || !s.Contains("elem4") {
			t.Errorf("After RemoveMany, set is missing elements that should remain")
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

	// Benchmark AddMany
	b.Run("AddMany", func(b *testing.B) {
		s := New()
		for i := 0; i < b.N; i++ {
			elements := make([]string, 1000)
			for j := 0; j < 1000; j++ {
				elements[j] = strconv.Itoa(j)
			}
			s.AddMany(elements)
		}
	})

	// Benchmark RemoveMany
	b.Run("RemoveMany", func(b *testing.B) {
		s := New()
		elements := make([]string, 1000)
		for i := 0; i < 1000; i++ {
			elements[i] = strconv.Itoa(i)
		}
		s.AddMany(elements)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.RemoveMany(elements)
		}
	})

	// Benchmark with different initial capacities
	initialCapacities := []int{16, 128, 1024, 8192}
	for _, cap := range initialCapacities {
		b.Run(fmt.Sprintf("AddWithCapacity%d", cap), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s := NewWithCapacity(cap)
				for j := 0; j < 1000; j++ {
					s.Add(strconv.Itoa(j))
				}
			}
		})
	}
}

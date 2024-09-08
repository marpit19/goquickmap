package quickset

import (
	"github.com/marpit19/goquickmap/pkg/quickmap"
)

// QuickSet represents a set data structure
type QuickSet struct {
	data *quickmap.QuickMap
}

func New() *QuickSet {
	return &QuickSet{
		data: quickmap.New(),
	}
}

// NewWithCapacity creates and returns a new QuickSet with the specified initial capacity
func NewWithCapacity(initialCapacity int) *QuickSet {
	return &QuickSet{
		data: quickmap.NewWithCapacity(initialCapacity),
	}
}

// Add inserts an element into the set
func (s *QuickSet) Add(element string) {
	s.data.Insert(element, struct{}{})
}

// Contains checks if an element exists in the set
func (s *QuickSet) Contains(element string) bool {
	_, exists := s.data.Get(element)
	return exists
}

// Remove deletes an element from the set
func (s *QuickSet) Remove(element string) {
	s.data.Delete(element)
}

// Size return sthe number of elements in the set
func (s *QuickSet) Size() int {
	return s.data.Size()
}

// Elements return a slice of all elements in the set
func (s *QuickSet) Elements() []string {
	elements := make([]string, 0, s.Size())
	s.data.ForEach(func(key string, value interface{}) {
		elements = append(elements, key)
	})
	return elements
}

// AddMany adds multiple elements to the set
func (s *QuickSet) AddMany(elements []string) {
	pairs := make(map[string]interface{}, len(elements))
	for _, elem := range elements {
		pairs[elem] = struct{}{}
	}
	s.data.InsertMany(pairs)
}

// RemoveMany removes multiple elements from the set
func (s *QuickSet) RemoveMany(elements []string) {
	s.data.DeleteMany(elements)
}

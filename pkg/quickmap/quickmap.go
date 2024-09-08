package quickmap

import "github.com/marpit19/goquickmap/internal/hash"

const (
	initalSize = 16
	loadFactor = 0.75
)

type node struct {
	key   string
	value interface{}
	next  *node
}

// QuickMap represents a hash table
type QuickMap struct {
	buckets []*node
	size    int
}

// creates and returns  a new QuickMap
func New() *QuickMap {
	return &QuickMap{
		buckets: make([]*node, initalSize),
		size:    0,
	}
}

// Insert adds a new key-value pair to our map
func (m *QuickMap) Insert(key string, value interface{}) {
	index := hash.Hash(key) % uint64(len(m.buckets))
	newNode := &node{key: key, value: value}

	if m.buckets[index] == nil {
		m.buckets[index] = newNode
	} else {
		current := m.buckets[index]
		for current.next != nil {
			if current.key == key {
				current.value = value
				return
			}
			current = current.next
		}
		if current.key == key {
			current.value = value
		} else {
			current.next = newNode
		}
	}
	m.size++
	
	if float64(m.size)/float64(len(m.buckets))>loadFactor {
		m.resize()
	}
}

// Get retrieves a value by key
func (m *QuickMap) Get(key string) (interface{}, bool) {
	index := hash.Hash(key) % uint64(len(m.buckets))
	current := m.buckets[index]
	
	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	
	return nil, false
}

// Delete removes a key-value pair from the map
func (m *QuickMap) Delete(key string) {
	index := hash.Hash(key) % uint64(len(m.buckets))
	if m.buckets[index] == nil {
		return
	}
	
	if m.buckets[index].key == key {
		m.buckets[index] = m.buckets[index].next
		m.size--
		return
	}
	
	current := m.buckets[index]
	for current.next != nil {
		if current.next.key == key {
			current.next = current.next.next
			m.size--
			return
		}
		
		current = current.next
	}
}

// resize increases the size of the hash table and reshases all the elements
func (m *QuickMap) resize() {
	newBuckets := make([]*node, len(m.buckets)*2)
	for _, bucket := range m.buckets {
		for bucket != nil {
			index := hash.Hash(bucket.key) % uint64(len(newBuckets))
			next := bucket.next
			bucket.next = newBuckets[index]
			newBuckets[index] = bucket
			bucket = next
		}
	}
	m.buckets = newBuckets
}
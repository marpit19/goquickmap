package quickdict

import (
	"github.com/marpit19/goquickmap/pkg/quickmap"
)

// QuickDict represnts a dictionary data strucutre
type QuickDict struct {
	data *quickmap.QuickMap
}

// New creates and returns a new QuickDict
func New() *QuickDict {
	return &QuickDict{
		data: quickmap.New(),
	}
}

// NewWithCapacity creates and returns a new QuickDict with the specified initial capacity
func NewWithCapacity(initialCapacity int) *QuickDict {
	return &QuickDict{
		data: quickmap.NewWithCapacity(initialCapacity),
	}
}

// Set inserts or updates a key-value pair in the dictionary
func (d *QuickDict) Set(key string, value interface{}) {
	d.data.Insert(key, value)
}

// Get retrieves a value by key from the dictionary
func (d *QuickDict) Get(key string) (interface{}, bool) {
	return d.data.Get(key)
}

// Delete removes a key-value pair from the dictionary
func (d *QuickDict) Delete(key string) {
	d.data.Delete(key)
}

// Size returns the number of key-value pairs in the dictionary
func (d *QuickDict) Size() int {
	return d.data.Size()
}

// Keys returns a slice of all keys in the dictionary
func (d *QuickDict) Keys() []string {
	keys := make([]string, 0, d.Size())
	d.data.ForEach(func(key string, value interface{}) {
		keys = append(keys, key)
	})
	return keys
}

// Values returns a slice of all values in the dictionary
func (d *QuickDict) Values() []interface{} {
	values := make([]interface{}, 0, d.Size())
	d.data.ForEach(func(key string, value interface{}) {
		values = append(values, value)
	})
	return values
}

// SetMany inserts or updates multiple key-value pairs in the dictionary
func (d *QuickDict) SetMany(pairs map[string]interface{}) {
	d.data.InsertMany(pairs)
}

// DeleteMany removes multiple key-value pairs from the dictionary
func (d *QuickDict) DeleteMany(keys []string) {
	d.data.DeleteMany(keys)
}

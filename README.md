# GoQuickMap

GoQuickMap is a high-performance, flexible hash table library for Go, featuring implementations of hash tables, sets, and maps.

## Features

- QuickMap: Efficient hash table implementation
- QuickSet: Set data structure built on QuickMap
- QuickDict: Dictionary/map data structure built on QuickMap
- Configurable initial capacity for optimized performance
- Batch operations for efficient bulk insertions and deletions

## Installation

To install GoQuickMap, use the following command:

```
go get github.com/marpit19/goquickmap
```

## Usage

### QuickMap

```go
import "github.com/marpit19/goquickmap/pkg/quickmap"

// Create a new QuickMap with default capacity
m := quickmap.New()

// Create a QuickMap with a specific initial capacity
m := quickmap.NewWithCapacity(1000)

// Insert a key-value pair
m.Insert("key", "value")

// Get a value
value, exists := m.Get("key")

// Delete a key-value pair
m.Delete("key")

// Batch insert
pairs := map[string]interface{}{
    "key1": "value1",
    "key2": "value2",
}
m.InsertMany(pairs)

// Batch delete
keys := []string{"key1", "key2"}
m.DeleteMany(keys)
```

### QuickSet

```go
import "github.com/marpit19/goquickmap/pkg/quickset"

// Create a new QuickSet with default capacity
s := quickset.New()

// Create a QuickSet with a specific initial capacity
s := quickset.NewWithCapacity(1000)

// Add an element
s.Add("element")

// Check if an element exists
exists := s.Contains("element")

// Remove an element
s.Remove("element")

// Batch add
elements := []string{"elem1", "elem2"}
s.AddMany(elements)

// Batch remove
s.RemoveMany(elements)
```

### QuickDict

```go
import "github.com/marpit19/goquickmap/pkg/quickdict"

// Create a new QuickDict with default capacity
d := quickdict.New()

// Create a QuickDict with a specific initial capacity
d := quickdict.NewWithCapacity(1000)

// Set a key-value pair
d.Set("key", "value")

// Get a value
value, exists := d.Get("key")

// Delete a key-value pair
d.Delete("key")

// Batch set
pairs := map[string]interface{}{
    "key1": "value1",
    "key2": "value2",
}
d.SetMany(pairs)

// Batch delete
keys := []string{"key1", "key2"}
d.DeleteMany(keys)
```

## Performance

GoQuickMap is designed for high performance. Here are some benchmark results on an Apple M3 Pro:

- QuickMap Insert: ~383 ns/op
- QuickMap Get: ~26 ns/op
- QuickMap Delete: ~26 ns/op
- QuickMap InsertMany (1000 items): ~172 µs/op
- QuickMap DeleteMany (1000 items): ~4.2 µs/op

- QuickSet Add: ~316 ns/op
- QuickSet Contains: ~25 ns/op
- QuickSet Remove: ~26 ns/op
- QuickSet AddMany (1000 items): ~115 µs/op
- QuickSet RemoveMany (1000 items): ~4.2 µs/op

- QuickDict Set: ~360 ns/op
- QuickDict Get: ~26 ns/op
- QuickDict Delete: ~25 ns/op
- QuickDict SetMany (1000 items): ~171 µs/op
- QuickDict DeleteMany (1000 items): ~4.2 µs/op

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
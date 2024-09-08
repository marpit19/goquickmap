# GoQuickMap

<center>

[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://pkg.go.dev/github.com/marpit19/goquickmap)
[![GitHub Actions](https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white)](https://github.com/marpit19/goquickmap/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/marpit19/goquickmap)](https://goreportcard.com/report/github.com/marpit19/goquickmap)

</center>

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

Note: All keys in the current implementation must be strings.

### QuickMap
```go
m := quickmap.New()
m.Insert("key", "value")  // key must be a string
value, exists := m.Get("key")
```

## Performance

GoQuickMap offers significant performance improvements over built-in Go maps and popular third-party set implementations. Here's a comparison based on 1,000,000 operations:

### Map Operations

| Operation    | Built-in Map | QuickMap    | Improvement |
|--------------|--------------|-------------|-------------|
| Insert       | 271.17ms     | 236.95ms    | 12.6% faster |
| Get          | 153.89ms     | 83.68ms     | 45.6% faster |
| Delete       | 188.82ms     | 70.47ms     | 62.7% faster |

QuickMap also supports efficient batch operations:
- Batch Insert (10,000 items): 638.54µs
- Batch Delete (10,000 items): 106.5µs

### Set Operations

| Operation    | golang-set   | QuickSet    | Improvement |
|--------------|--------------|-------------|-------------|
| Add          | 231.33ms     | 211.49ms    | 8.6% faster |
| Contains     | 258.16ms     | 88.20ms     | 65.8% faster |
| Remove       | 232.36ms     | 78.90ms     | 66.0% faster |

QuickSet also supports efficient batch operations:
- Batch Add (10,000 items): 1.26ms
- Batch Remove (10,000 items): 111.67µs

### Analysis

1. **Superior Performance**: GoQuickMap consistently outperforms built-in maps and popular set implementations across all operations.

2. **Significant Speedup for Lookups and Deletions**: QuickMap and QuickSet show dramatic improvements in Get/Contains (45-65% faster) and Delete/Remove operations (62-66% faster).

3. **Efficient Insertions**: Both QuickMap and QuickSet demonstrate faster insertion times compared to their counterparts.

4. **Batch Operations**: The library offers highly efficient batch operations, allowing for rapid insertion and deletion of multiple items simultaneously.

5. **Consistent Advantage**: The performance advantage is maintained across different types of operations, indicating a well-optimized underlying structure.

These results demonstrate that GoQuickMap is an excellent choice for applications requiring high-performance hash tables, maps, or sets, especially those dealing with large datasets or frequent lookup and deletion operations.

## Current Limitations and Future Plans

### String Keys Only
The current implementation of GoQuickMap, including QuickMap, QuickSet, and QuickDict, only supports string keys. This design choice was made to optimize performance for string-based keys, which are common in many applications.

#### Implications:
1. **Use Case Focus**: The library is currently best suited for applications that primarily use string identifiers or textual data as keys.
2. **Performance Optimization**: The string-specific implementation allows for optimizations that may not be possible with a more generic approach.
3. **Benchmark Context**: The performance comparisons provided are specifically for string keys and may vary for other types of keys.

### Future Considerations
We acknowledge that supporting only string keys is a limitation. Potential future enhancements may include:

1. **Generic Key Support**: Implementing support for generic types as keys, allowing for greater flexibility.
2. **Numeric Key Optimization**: Adding specialized implementations for common numeric types (int, int64, etc.) that could potentially offer even better performance for these types.
3. **Custom Hash Functions**: Allowing users to provide custom hash functions for their specific key types.

We welcome feedback and contributions from the community regarding these potential improvements. If you have specific use cases that require non-string keys, please open an issue to discuss your needs.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

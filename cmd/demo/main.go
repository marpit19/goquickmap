package main

import (
	"fmt"

	"github.com/marpit19/goquickmap/pkg/quickdict"
	"github.com/marpit19/goquickmap/pkg/quickmap"
	"github.com/marpit19/goquickmap/pkg/quickset"
)

func main() {
    fmt.Println("GoQuickMap Demo")

    // QuickMap demo
    fmt.Println("\n--- QuickMap Demo ---")
    demoQuickMap()

    // QuickSet demo
    fmt.Println("\n--- QuickSet Demo ---")
    demoQuickSet()

    // QuickDict demo
    fmt.Println("\n--- QuickDict Demo ---")
    demoQuickDict()
}

func demoQuickMap() {
    // Create a new QuickMap with default capacity
    m := quickmap.New()

    // Insert some key-value pairs
    m.Insert("name", "Alice")
    m.Insert("age", 30)
    m.Insert("city", "New York")

    // Get and print values
    name, _ := m.Get("name")
    age, _ := m.Get("age")
    fmt.Printf("Name: %s, Age: %d\n", name, age)

    // Delete a key-value pair
    m.Delete("city")

    // Check if a key exists
    _, exists := m.Get("city")
    fmt.Printf("City exists: %v\n", exists)

    // Batch insert
    m.InsertMany(map[string]interface{}{
        "country": "USA",
        "language": "English",
    })

    // Batch delete
    m.DeleteMany([]string{"name", "age"})

    // Print the size of the map
    fmt.Printf("Map size: %d\n", m.Size())

    // Create a QuickMap with a specific initial capacity
    largeMap := quickmap.NewWithCapacity(1000)
    fmt.Printf("Large map initial size: %d\n", largeMap.Size())
}

func demoQuickSet() {
    // Create a new QuickSet
    s := quickset.New()

    // Add some elements
    s.Add("apple")
    s.Add("banana")
    s.Add("cherry")

    // Check if an element exists
    fmt.Printf("Contains 'banana': %v\n", s.Contains("banana"))

    // Remove an element
    s.Remove("cherry")

    // Batch add
    s.AddMany([]string{"date", "elderberry", "fig"})

    // Batch remove
    s.RemoveMany([]string{"apple", "banana"})

    // Print the size of the set
    fmt.Printf("Set size: %d\n", s.Size())

    // Create a QuickSet with a specific initial capacity
    largeSet := quickset.NewWithCapacity(1000)
    fmt.Printf("Large set initial size: %d\n", largeSet.Size())
}

func demoQuickDict() {
    // Create a new QuickDict
    d := quickdict.New()

    // Set some key-value pairs
    d.Set("name", "Bob")
    d.Set("age", 25)
    d.Set("city", "San Francisco")

    // Get and print values
    name, _ := d.Get("name")
    age, _ := d.Get("age")
    fmt.Printf("Name: %s, Age: %d\n", name, age)

    // Delete a key-value pair
    d.Delete("city")

    // Batch set
    d.SetMany(map[string]interface{}{
        "country": "USA",
        "language": "English",
    })

    // Batch delete
    d.DeleteMany([]string{"name", "age"})

    // Print the size of the dictionary
    fmt.Printf("Dictionary size: %d\n", d.Size())

    // Create a QuickDict with a specific initial capacity
    largeDict := quickdict.NewWithCapacity(1000)
    fmt.Printf("Large dictionary initial size: %d\n", largeDict.Size())
}

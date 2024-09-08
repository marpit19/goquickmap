// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"runtime"
// 	"strconv"
// 	"time"

// 	mapset "github.com/deckarep/golang-set/v2"
// 	"github.com/marpit19/goquickmap/pkg/quickmap"
// 	"github.com/marpit19/goquickmap/pkg/quickset"
// )

// const (
//     numOperations = 1000000
//     numBatchOperations = 10000
// )

// func main() {
//     fmt.Println("Performance Comparison")
//     fmt.Printf("Number of operations: %d\n", numOperations)
//     fmt.Printf("Number of batch operations: %d\n", numBatchOperations)

//     compareMap()
//     compareSet()
// }

// func compareMap() {
//     fmt.Println("\n--- Map Comparison ---")

//     // Built-in map
//     start := time.Now()
//     m := make(map[string]int)
//     for i := 0; i < numOperations; i++ {
//         key := strconv.Itoa(i)
//         m[key] = i
//     }
//     builtinInsertTime := time.Since(start)

//     start = time.Now()
//     for i := 0; i < numOperations; i++ {
//         key := strconv.Itoa(i)
//         _ = m[key]
//     }
//     builtinGetTime := time.Since(start)

//     start = time.Now()
//     for i := 0; i < numOperations; i++ {
//         key := strconv.Itoa(i)
//         delete(m, key)
//     }
//     builtinDeleteTime := time.Since(start)

//     // QuickMap
//     start = time.Now()
//     qm := quickmap.New()
//     for i := 0; i < numOperations; i++ {
//         key := strconv.Itoa(i)
//         qm.Insert(key, i)
//     }
//     quickmapInsertTime := time.Since(start)

//     start = time.Now()
//     for i := 0; i < numOperations; i++ {
//         key := strconv.Itoa(i)
//         _, _ = qm.Get(key)
//     }
//     quickmapGetTime := time.Since(start)

//     start = time.Now()
//     for i := 0; i < numOperations; i++ {
//         key := strconv.Itoa(i)
//         qm.Delete(key)
//     }
//     quickmapDeleteTime := time.Since(start)

//     // Batch operations
//     batchKeys := make([]string, numBatchOperations)
//     batchMap := make(map[string]interface{}, numBatchOperations)
//     for i := 0; i < numBatchOperations; i++ {
//         key := strconv.Itoa(rand.Intn(numOperations))
//         batchKeys[i] = key
//         batchMap[key] = i
//     }

//     start = time.Now()
//     qm.InsertMany(batchMap)
//     quickmapBatchInsertTime := time.Since(start)

//     start = time.Now()
//     qm.DeleteMany(batchKeys)
//     quickmapBatchDeleteTime := time.Since(start)

//     // Print results
//     fmt.Println("Built-in map:")
//     fmt.Printf("  Insert: %v\n", builtinInsertTime)
//     fmt.Printf("  Get: %v\n", builtinGetTime)
//     fmt.Printf("  Delete: %v\n", builtinDeleteTime)

//     fmt.Println("QuickMap:")
//     fmt.Printf("  Insert: %v\n", quickmapInsertTime)
//     fmt.Printf("  Get: %v\n", quickmapGetTime)
//     fmt.Printf("  Delete: %v\n", quickmapDeleteTime)
//     fmt.Printf("  Batch Insert (%d items): %v\n", numBatchOperations, quickmapBatchInsertTime)
//     fmt.Printf("  Batch Delete (%d items): %v\n", numBatchOperations, quickmapBatchDeleteTime)
// }

// func compareSet() {
//     fmt.Println("\n--- Set Comparison ---")

//     // golang-set
//     start := time.Now()
//     s := mapset.NewSet[string]()
//     for i := 0; i < numOperations; i++ {
//         s.Add(strconv.Itoa(i))
//     }
//     mapsetAddTime := time.Since(start)

//     start = time.Now()
//     for i := 0; i < numOperations; i++ {
//         s.Contains(strconv.Itoa(i))
//     }
//     mapsetContainsTime := time.Since(start)

//     start = time.Now()
//     for i := 0; i < numOperations; i++ {
//         s.Remove(strconv.Itoa(i))
//     }
//     mapsetRemoveTime := time.Since(start)

//     // QuickSet
//     start = time.Now()
//     qs := quickset.New()
//     for i := 0; i < numOperations; i++ {
//         qs.Add(strconv.Itoa(i))
//     }
//     quicksetAddTime := time.Since(start)

//     start = time.Now()
//     for i := 0; i < numOperations; i++ {
//         qs.Contains(strconv.Itoa(i))
//     }
//     quicksetContainsTime := time.Since(start)

//     start = time.Now()
//     for i := 0; i < numOperations; i++ {
//         qs.Remove(strconv.Itoa(i))
//     }
//     quicksetRemoveTime := time.Since(start)

//     // Batch operations
//     batchElements := make([]string, numBatchOperations)
//     for i := 0; i < numBatchOperations; i++ {
//         batchElements[i] = strconv.Itoa(rand.Intn(numOperations))
//     }

//     start = time.Now()
//     qs.AddMany(batchElements)
//     quicksetBatchAddTime := time.Since(start)

//     start = time.Now()
//     qs.RemoveMany(batchElements)
//     quicksetBatchRemoveTime := time.Since(start)

//     // Print results
//     fmt.Println("golang-set:")
//     fmt.Printf("  Add: %v\n", mapsetAddTime)
//     fmt.Printf("  Contains: %v\n", mapsetContainsTime)
//     fmt.Printf("  Remove: %v\n", mapsetRemoveTime)

//     fmt.Println("QuickSet:")
//     fmt.Printf("  Add: %v\n", quicksetAddTime)
//     fmt.Printf("  Contains: %v\n", quicksetContainsTime)
//     fmt.Printf("  Remove: %v\n", quicksetRemoveTime)
//     fmt.Printf("  Batch Add (%d items): %v\n", numBatchOperations, quicksetBatchAddTime)
//     fmt.Printf("  Batch Remove (%d items): %v\n", numBatchOperations, quicksetBatchRemoveTime)
// }

// func printMemUsage() {
//     var m runtime.MemStats
//     runtime.ReadMemStats(&m)
//     fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
//     fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
//     fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
//     fmt.Printf("\tNumGC = %v\n", m.NumGC)
// }

// func bToMb(b uint64) uint64 {
//     return b / 1024 / 1024
// }

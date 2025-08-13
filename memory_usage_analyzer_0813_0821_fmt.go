// 代码生成时间: 2025-08-13 08:21:40
package main

import (
    "fmt"
    "os"
    "runtime"
    "time"
    
    "github.com/revel/revel"
)

// MemoryUsage contains the information about memory usage
type MemoryUsage struct {
    Alloc       uint64 `json:"alloc"`       // bytes allocated and not yet freed
    TotalAlloc uint64 `json:"total_alloc"` // bytes allocated (even if freed)
    Sys         uint64 `json:"sys"`         // bytes obtained from system (includes stack)
    Looks       uint64 `json:"looks"`       // number of pointer lookups
    Mallocs     uint64 `json:"mallocs"`     // number of mallocs
    Frees       uint64 `json:"frees"`       // number of frees
    HeapAlloc   uint64 `json:"heap_alloc"`  // bytes allocated in heap
    HeapSys     uint64 `json:"heap_sys"`    // heap system bytes
    HeapIdle    uint64 `json:"heap_idle"`   // heap idle bytes
    HeapInuse   uint64 `json:"heap_inuse"`  // heap in-use bytes
    HeapReleased uint64 `json:"heap_released"` // bytes released to the OS
    HeapObjects  uint64 `json:"heap_objects
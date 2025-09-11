// 代码生成时间: 2025-09-12 07:54:56
package main

import (
    "fmt"
    "os"
    "runtime"
    "runtime/pprof"
    "log"
    "time"
)

// MemoryAnalyzer 结构体包含内存分析器
type MemoryAnalyzer struct {
    CPUProfile    string
    MemoryProfile string
}

// NewMemoryAnalyzer 创建一个新的内存分析器，设置默认的分析文件路径
func NewMemoryAnalyzer() *MemoryAnalyzer {
    return &MemoryAnalyzer{
        CPUProfile:    "cpu.prof",
        MemoryProfile: "mem.prof",
    }
}

// StartCPUProfile 开始CPU分析
func (ma *MemoryAnalyzer) StartCPUProfile() error {
    f, err := os.Create(ma.CPUProfile)
    if err != nil {
        return err
    }
    defer f.Close()
    if err := pprof.StartCPUProfile(f); err != nil {
        return err
    }
    fmt.Println("CPU profiling started.")
    return nil
}

// StopCPUProfile 停止CPU分析
func (ma *MemoryAnalyzer) StopCPUProfile() {
    pprof.StopCPUProfile()
    fmt.Println("CPU profiling stopped.")
}

// StartMemoryProfile 开始内存分析
func (ma *MemoryAnalyzer) StartMemoryProfile() error {
    f, err := os.Create(ma.MemoryProfile)
    if err != nil {
        return err
    }
    defer f.Close()
    if err := pprof.WriteHeapProfile(f); err != nil {
        return err
    }
    fmt.Println("Memory profiling started.")
    return nil
}

// AnalyzeMemory 使用pprof分析内存使用情况
func (ma *MemoryAnalyzer) AnalyzeMemory() {
    // 获取当前内存使用情况
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("Alloc = %v MiB
", m.Alloc/1024/1024)
    fmt.Printf("TotalAlloc = %v MiB
", m.TotalAlloc/1024/1024)
    // 其他内存统计...
}

func main() {
    ma := NewMemoryAnalyzer()
    if err := ma.StartCPUProfile(); err != nil {
        log.Fatalf("Failed to start CPU profile: %v", err)
    }
    defer ma.StopCPUProfile()

    if err := ma.StartMemoryProfile(); err != nil {
        log.Fatalf("Failed to start memory profile: %v", err)
    }
    defer func() {
        pprof.Lookup("heap").WriteTo(os.Stdout, 0)
    }()

    // 模拟一些内存分配
    for i := 0; i < 1024; i++ {
        memory := make([]byte, 1024)
        fmt.Println(string(memory))
    }

    // 内存分析
    ma.AnalyzeMemory()

    // 等待10秒以确保pprof有足够的时间写入文件
    time.Sleep(10 * time.Second)
}

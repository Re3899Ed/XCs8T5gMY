// 代码生成时间: 2025-09-06 12:39:44
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "syscall"
    "time"
    "revel"
)

// TestSuite 结构体用于存储测试参数
type TestSuite struct {
    TestName string
    Command  string
}

// RunTest 执行一个测试命令
func (t *TestSuite) RunTest() error {
    // 创建命令
    cmd := revel.Command(t.Command)
    
    // 执行命令
    if err := cmd.Exec(); err != nil {
        return fmt.Errorf("error running test command: %v", err)
    }
    
    // 检查命令退出状态
    if cmd.ProcessState.ExitCode() != 0 {
        return fmt.Errorf("test command exited with non-zero status: %v", cmd.ProcessState.ExitCode())
    }
    
    return nil
}

// DiscoverTests 发现测试文件
func DiscoverTests(directory string) ([]TestSuite, error) {
    var testSuites []TestSuite
    
    // 遍历目录中的文件
    err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && filepath.Ext(path) == ".test" {
            testSuites = append(testSuites, TestSuite{
                TestName: filepath.Base(path),
                Command:  fmt.Sprintf("go test %s", path),
            })
        }
        return nil
    })
    
    if err != nil {
        return nil, fmt.Errorf("error discovering test files: %v", err)
    }
    
    return testSuites, nil
}

func main() {
    // 默认测试目录
    testDir := "./tests"
    
    // 发现测试
    testSuites, err := DiscoverTests(testDir)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    
    // 执行测试
    for _, test := range testSuites {
        fmt.Printf("Running test: %s
", test.TestName)
        if err := test.RunTest(); err != nil {
            fmt.Fprintf(os.Stderr, "Error running test %s: %v
", test.TestName, err)
            os.Exit(1)
        }
    }
    
    fmt.Println("All tests passed successfully.")
}

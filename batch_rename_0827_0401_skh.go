// 代码生成时间: 2025-08-27 04:01:32
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "revel"
)

// BatchRenameApp 结构体，用于应用的配置和操作
type BatchRenameApp struct {
    *revel.Controller
}

// Index 路由到重命名页面
func (c BatchRenameApp) Index() revel.Result {
    return c.Render()
}

// RenameFiles 处理批量文件重命名请求
func (c BatchRenameApp) RenameFiles() revel.Result {
    var result struct {
        Success bool   "json:"success"
        Message string "json:"message"
    }

    // 获取表单参数
    pattern := c.Params.Get("pattern")
    dir := c.Params.Get("dir")
    newBaseName := c.Params.Get("newBaseName")

    // 参数校验
    if pattern == "" || dir == "" || newBaseName == "" {
        result.Success = false
        result.Message = "All parameters are required."
        return c.RenderJSON(result)
    }

    // 读取目录中的文件
    files, err := os.ReadDir(dir)
    if err != nil {
        result.Success = false
        result.Message = "Failed to read directory: " + err.Error()
        return c.RenderJSON(result)
    }

    // 重命名文件
    for _, file := range files {
        if file.IsDir() {
            continue
        }
        oldPath := filepath.Join(dir, file.Name())
        newPath := filepath.Join(dir, fmt.Sprintf(pattern, newBaseName, file.Name()))
        err := os.Rename(oldPath, newPath)
        if err != nil {
            result.Success = false
            result.Message = "Failed to rename file: " + err.Error()
            return c.RenderJSON(result)
        }
    }

    // 如果所有文件都成功重命名
    result.Success = true
    result.Message = "Files renamed successfully."
    return c.RenderJSON(result)
}

func init() {
    revel.Filter(func(c *revel.Controller, fc []revel.Filter) {
        // Before action filters
    }, revel.BEFORE, *BatchRenameApp)

    revel.Filter(func(c *revel.Controller, fc []revel.Filter) {
        // After action filters
    }, revel.AFTER, *BatchRenameApp)

    // Register routes
    revel.Router.
        PathPrefix("/").Name("batch-rename").
        Use(BatchRenameApp{}, func(c *BatchRenameApp, fc []revel.Filter) revel.Result {
            c.Index()
        })

    revel.Router.
        Path("/rename").
        Name("rename-files").
        Use(BatchRenameApp{}, func(c *BatchRenameApp, fc []revel.Filter) revel.Result {
            c.RenameFiles()
        })
}

// 代码生成时间: 2025-08-05 14:12:11
package main

import (
    "archive/zip"
    "io"
    "io/fs"
    "log"
    "os"
    "path/filepath"
    "revel"
)

// UnzipService 结构体定义了解压文件服务
type UnzipService struct {
    revel.Controller
}

// UnzipAction 定义了解压文件的动作
func (c UnzipService) UnzipAction(zipFilePath, targetDir string) (string, error) {
    // 检查文件是否存在
    if _, err := os.Stat(zipFilePath); os.IsNotExist(err) {
        return "", err
    }

    // 读取zip文件
    reader, err := zip.OpenReader(zipFilePath)
    if err != nil {
        return "", err
    }
    defer reader.Close()

    // 创建目标目录
    if err := os.MkdirAll(targetDir, 0755); err != nil {
        return "", err
    }

    // 解压文件
    for _, file := range reader.File {
        targetFilePath := filepath.Join(targetDir, file.Name)
        if file.FileInfo().IsDir() {
            // 创建目录
            if err := os.MkdirAll(targetFilePath, 0755); err != nil {
                return "", err
            }
        } else {
            // 创建文件
            outFile, err := os.OpenFile(targetFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
            if err != nil {
                return "", err
            }
            defer outFile.Close()
            
            // 读取并写入文件内容
            rc, err := file.Open()
            if err != nil {
                return "", err
            }
            defer rc.Close()
            _, err = io.Copy(outFile, rc)
            if err != nil {
                return "", err
            }
        }
    }

    return "Unzipped successfully", nil
}

func main() {
    // Revel初始化
    revel.Init()
    defer revel.Close()

    // 设置路由
    revel.Router(
        &UnzipService{},
        []string{"GET", "POST"},
        "/unzip",
        "UnzipAction",
    )

    // 启动Revel服务器
    revel.Run(8080)
}
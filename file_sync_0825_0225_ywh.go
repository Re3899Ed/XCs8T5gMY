// 代码生成时间: 2025-08-25 02:25:33
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "revel"
    "strings"
)

// BackupService 处理文件备份和同步的服务
type BackupService struct {
    *revel.Controller
}

// NewBackupService 创建一个新的BackupService实例
func NewBackupService() *BackupService {
    return &BackupService{
        Controller: &revel.Controller{},
    }
}

// SyncFiles 同步源和目标目录中的文件
func (b *BackupService) SyncFiles(source, destination string) error {
    // 确保源路径存在
    if _, err := os.Stat(source); os.IsNotExist(err) {
        return fmt.Errorf("source directory does not exist: %s", source)
    }

    // 确保目标路径存在
    if _, err := os.Stat(destination); os.IsNotExist(err) {
        if err := os.MkdirAll(destination, 0755); err != nil {
            return fmt.Errorf("failed to create destination directory: %s", destination)
        }
    }

    // 遍历源目录
    err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // 跳过目录
        if info.IsDir() {
            return nil
        }

        // 构造目标文件路径
        relativePath, err := filepath.Rel(source, path)
        if err != nil {
            return err
        }
        destPath := filepath.Join(destination, relativePath)

        // 检查目标文件是否已存在
        if _, err := os.Stat(destPath); os.IsNotExist(err) {
            // 复制文件到目标路径
            if err := b.copyFile(path, destPath); err != nil {
                return err
            }
        } else {
            // 比较文件大小和最后修改时间
            if err := b.compareFiles(path, destPath); err != nil {
                return err
            }
        }

        return nil
    })

    return err
}

// copyFile 复制文件从源路径到目标路径
func (b *BackupService) copyFile(sourcePath, destinationPath string) error {
    sourceFile, err := os.Open(sourcePath)
    if err != nil {
        return err
    }
    defer sourceFile.Close()

    destinationFile, err := os.Create(destinationPath)
    if err != nil {
        return err
    }
    defer destinationFile.Close()

    _, err = io.Copy(destinationFile, sourceFile)
    return err
}

// compareFiles 比较两个文件的大小和最后修改时间
func (b *BackupService) compareFiles(path1, path2 string) error {
    info1, err := os.Stat(path1)
    if err != nil {
        return err
    }

    info2, err := os.Stat(path2)
    if err != nil {
        return err
    }

    if info1.Size() != info2.Size() || info1.ModTime() != info2.ModTime() {
        if err := b.copyFile(path1, path2); err != nil {
            return err
        }
    }
    return nil
}

func main() {
    // 创建新的BackupService实例
    service := NewBackupService()

    // 设置源和目标目录路径
    source := "/path/to/source/directory"
    destination := "/path/to/destination/directory"

    // 同步文件
    if err := service.SyncFiles(source, destination); err != nil {
        log.Fatal(err)
    } else {
        fmt.Println("Files have been synced successfully.")
    }
}

// 代码生成时间: 2025-09-14 09:42:53
package main

import (
    "archive/zip"
    "io"
    "log"
    "os"
    "path/filepath"
    "revel"
)

// UnzipToolController is the controller for unzipping files.
type UnzipToolController struct {
    *revel.Controller
}

// UnzipAction handles the unzipping of a zip file.
// It takes a file path and destination as parameters.
func (c UnzipToolController) UnzipAction(zipFilePath, destination string) revel.Result {
    // Check if the provided zip file path is valid.
    if _, err := os.Stat(zipFilePath); os.IsNotExist(err) {
        c.Flash.Error("The provided zip file does not exist.")
        return c.Redirect(UnzipToolController.Unzip)
    }

    // Open the zip file.
    reader, err := zip.OpenReader(zipFilePath)
    if err != nil {
        c.Flash.Error("Error opening zip file: " + err.Error())
        return c.Redirect(UnzipToolController.Unzip)
    }
    defer reader.Close()

    // Create the destination directory if it does not exist.
    if err := os.MkdirAll(destination, 0755); err != nil {
        c.Flash.Error("Error creating destination directory: " + err.Error())
        return c.Redirect(UnzipToolController.Unzip)
    }

    // Iterate through the files in the zip archive.
    for _, file := range reader.File {
        // Set the path for the file in the destination directory.
        filePath := filepath.Join(destination, file.Name)

        // Check if the file is a directory, if so create the directory.
        if file.FileInfo().IsDir() {
            if err := os.MkdirAll(filePath, file.Mode()); err != nil {
                c.Flash.Error("Error creating directory: " + err.Error())
                return c.Redirect(UnzipToolController.Unzip)
            }
            continue
        }

        // Create the file in the destination directory.
        outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
        if err != nil {
            c.Flash.Error("Error opening file: " + err.Error())
            return c.Redirect(UnzipToolController.Unzip)
        }
        defer outFile.Close()

        // Copy the file from the zip archive to the destination directory.
        rc, err := file.Open()
        if err != nil {
            c.Flash.Error("Error opening file in zip archive: " + err.Error())
            return c.Redirect(UnzipToolController.Unzip)
        }
        defer rc.Close()

        _, err = io.Copy(outFile, rc)
        if err != nil {
            c.Flash.Error("Error copying file: " + err.Error())
            return c.Redirect(UnzipToolController.Unzip)
        }
    }

    c.Flash.Success("The file has been successfully unzipped.")
    return c.Redirect(UnzipToolController.Unzip)
}

func init() {
    revel.Router("(unzip)/(zipFilePath string)/(destination string)", &UnzipToolController{}, "UnzipAction")
}

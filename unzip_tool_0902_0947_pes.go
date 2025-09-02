// 代码生成时间: 2025-09-02 09:47:08
 * This tool provides a basic interface to upload and decompress files.
 */

package main

import (
    "fmt"
    "os"
    "net/http"
    "path/filepath"
    "archive/zip"
    "io"
    "log"
    "revel"
)

// DecompressFile decompresses a zip file to the specified destination.
func DecompressFile(src, dest string) error {
    r, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer r.Close()
    
    for _, f := range r.File {
        // Create directory if it doesn't exist.
        if f.FileInfo().IsDir() {
            os.MkdirAll(filepath.Join(dest, f.Name), os.ModePerm)
            continue
        }
        
        // Create the file.
        out, err := os.OpenFile(
            filepath.Join(dest, f.Name),
            os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
            f.Mode())
        if err != nil {
            return err
        }
        defer out.Close()
        
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer rc.Close()
        
        _, err = io.Copy(out, rc)
        if err != nil {
            return err
        }
    }
    return nil
}

// Handler for the file upload.
type FileUpload struct {
    *revel.Controller
}

// Action to handle the file upload.
func (c FileUpload) Action() revel.Result {
    if c.Request.Method == "POST" {
        // Get the uploaded file.
        file, _, err := c.Request.FormFile("file")
        if err != nil {
            c.Flash.Error("Error reading the uploaded file.")
            return c.RenderError(err)
        }
        defer file.Close()
        
        // Specify the destination folder.
        dest := "./uploads/"
        os.MkdirAll(dest, os.ModePerm)
        
        // Save the file to disk.
        destFile, err := os.Create(dest + file.Filename)
        if err != nil {
            c.Flash.Error("Error saving the file.")
            return c.RenderError(err)
        }
        defer destFile.Close()
        
        _, err = io.Copy(destFile, file)
        if err != nil {
            c.Flash.Error("Error writing the file.")
            return c.RenderError(err)
        }
        
        // Decompress the file.
        err = DecompressFile(dest+file.Filename, dest)
        if err != nil {
            c.Flash.Error("Error decompressing the file.")
            return c.RenderError(err)
        }
        
        c.Flash.Success("File uploaded and decompressed successfully.")
        return c.Render()
    }
    return c.Render()
}

func init() {
    revel.Filter("unzip.Tool", func(c *revel.Controller, fc func(*revel.Controller) revel.Result) revel.Result {
        // You can add filters here to handle specific pre and post actions.
        return fc(c)
    }, revel.BEFORE)
}

func main() {
    revel.Run(
        // Set the app properties here.
        // For example, set the port, mode, etc.
        revel.DevMode,
        0,
    )
}
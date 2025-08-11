// 代码生成时间: 2025-08-11 11:34:34
package main

import (
    "image"
    "image/jpeg"
    "image/png"
    "os"
    "path/filepath"
    "revel"
    "strings"
)

// ImageResizer is a controller for resizing images.
type ImageResizer struct {
    revel.Controller
}

// Resize action resizes images in a given directory to a specified width and height.
func (c ImageResizer) Resize(dir string, width int, height int) revel.Result {
    // Get the list of files in the directory.
    files, err := os.ReadDir(dir)
    if err != nil {
        return c.RenderError(err)
    }

    for _, file := range files {
        if !file.IsDir() {
            // Construct the full path to the file.
            filePath := filepath.Join(dir, file.Name())

            // Open the image file.
            img, err := os.Open(filePath)
            if err != nil {
                // Log the error and continue to the next file.
                revel.ERROR.Printf("Failed to open image: %s", err)
                continue
            }
            defer img.Close()

            // Decode the image.
            imgConfig, _, err := image.DecodeConfig(img)
            if err != nil {
                revel.ERROR.Printf("Failed to decode image: %s", err)
                continue
            }

            // Check if the image is JPEG or PNG.
            switch imgConfig.Format {
            case "jpeg", "png":
                // Resize the image.
                resizedImg := resizeImage(img, width, height)
                if resizedImg == nil {
                    revel.ERROR.Printf("Failed to resize image: %s", filePath)
                    continue
                }

                // Save the resized image.
                if err := saveResizedImage(resizedImg, filePath, imgConfig.Format); err != nil {
                    revel.ERROR.Printf("Failed to save resized image: %s", err)
                }
            default:
                revel.ERROR.Printf("Unsupported image format: %s", imgConfig.Format)
            }
        }
    }

    return c.RenderJSON(revel.Map{
        "status": "success",
        "message": "Images resized successfully",
    })
}

// resizeImage resizes the image to the specified width and height.
func resizeImage(imgFile *os.File, width int, height int) image.Image {
    // Decode the image.
    img, _, err := image.Decode(imgFile)
    if err != nil {
        return nil
    }

    // Create a new image with the specified width and height.
    imgResized := image.NewRGBA(image.Rect(0, 0, width, height))

    // Calculate the scaling factor.
    scalingFactor := width / float64(img.Bounds().Dx())

    // Draw the original image onto the resized image.
    imgResized = resize.Resize(width, height, img, resize.Lanczos3)

    return imgResized
}

// saveResizedImage saves the resized image to a file.
func saveResizedImage(img image.Image, filePath string, format string) error {
    // Create a new file for the resized image.
    outFile, err := os.Create(filePath + ".resized")
    if err != nil {
        return err
    }
    defer outFile.Close()

    // Save the resized image to the file.
    switch format {
    case "jpeg":
        return jpeg.Encode(outFile, img, nil)
    case "png":
        return png.Encode(outFile, img)
    default:
        return fmt.Errorf("unsupported image format: %s", format)
    }
}

// 代码生成时间: 2025-08-08 06:15:12
package excel

import (
    "os"
    "path/filepath"
    "encoding/csv"
    "log"
    "time"
    "revel"
    "github.com/tealeg/xlsx/v3"
)

// ExcelGenerator struct contains fields to manage Excel file creation
type ExcelGenerator struct {
    File *xlsx.File
}

// NewExcelGenerator creates a new instance of ExcelGenerator
func NewExcelGenerator() *ExcelGenerator {
    return &ExcelGenerator{
        File: xlsx.NewFile(),
    }
}

// CreateExcelSheet creates a new sheet in the Excel file
func (eg *ExcelGenerator) CreateExcelSheet(name string) *xlsx.Sheet {
    sheet, err := eg.File.AddSheet(name)
    if err != nil {
        revel.ERROR.Printf("Error creating sheet: %s", err)
        return nil
    }
    return sheet
}

// AddRow adds a new row to the specified sheet
func (eg *ExcelGenerator) AddRow(sheetName string, row []string) error {
    sheet, ok := eg.File.Sheet[sheetName]
    if !ok {
        return fmt.Errorf("sheet not found: %s", sheetName)
    }
    newRow := sheet.AddRow()
    for _, cell := range row {
        newRow.AddCell().Value = cell
    }
    return nil
}

// SaveExcelFile saves the Excel file to the specified location
func (eg *ExcelGenerator) SaveExcelFile(filePath string) error {
    if err := eg.File.Save(filePath); err != nil {
        return err
    }
    return nil
}

// Controller for handling HTTP requests
type ExcelController struct {
    *revel.Controller
}

// Generate action generates an Excel file based on provided parameters
func (c ExcelController) Generate() revel.Result {
    filePath := filepath.Join(revel.BasePath, "tmp", "generated_excel.xlsx")
    eg := NewExcelGenerator()
    sheet := eg.CreateExcelSheet("Sheet1")
    if sheet == nil {
        return c.RenderError(revel.StatusInternalServerError)
    }
    if err := eg.AddRow("Sheet1", []string{"Header1", "Header2", time.Now().Format(time.RFC3339)}); err != nil {
        return c.RenderError(revel.StatusInternalServerError)
    }
    if err := eg.SaveExcelFile(filePath); err != nil {
        return c.RenderError(revel.StatusInternalServerError)
    }
    return c.RenderFile(filePath)
}

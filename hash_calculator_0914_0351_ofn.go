// 代码生成时间: 2025-09-14 03:51:20
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "github.com/revel/revel"
    "strings"
)

// HashCalculator is the controller for hash value calculation.
type HashCalculator struct {
    revel.Controller
}

// CalculateSHA256 calculates the SHA-256 hash of the given input string.
func (c HashCalculator) CalculateSHA256() revel.Result {
    input := c.Params.Form["input"]string{}
    if len(input) == 0 {
        // Return error if no input is provided.
        return c.RenderError(revel.NewErrorResult(revel.NewRequestError(
            "errors.input_required",
            "Input is required for hash calculation.",
            nil,
            revel.NewHTTPResponse(400),
        ))
    }
    
    // Calculate the SHA-256 hash.
    h := sha256.New()
    h.Write([]byte(input))
    hashed := h.Sum(nil)
    hashString := hex.EncodeToString(hashed)
    
    // Return the result.
    return c.RenderJSON(map[string]string{
        "input": input,
        "hash": hashString,
    })
}

func init() {
    // Filters is a list of global filters.
    revel.Filters = []revel.Filter{
        // Add global filters here, e.g.,
        revel.RouterFilter,
        revel.SessionFilter,
        revel.FlashFilter,
        revel.ParamFilter,
        revel.ActionInvoker,
    }

    // Register startup functions with OnAppStart.
    revel.OnAppStart(func() {
        // Initialize the Revel configuration here.
    })
}

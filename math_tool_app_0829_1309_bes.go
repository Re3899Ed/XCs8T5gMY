// 代码生成时间: 2025-08-29 13:09:06
package main

import (
    "encoding/json"
    "github.com/revel/revel"
)

// Application structure is used by Revel to handle HTTP requests and responses.
type Application struct {
    *revel.Controller
}

// Home is the action for the application's root route.
func (c Application) Home() revel.Result {
    return c.Render()
}

// Add handles addition operation.
func (c Application) Add(a, b float64) revel.Result {
    sum := a + b
    return c.RenderJSON(SumResult{Sum: sum})
}

// Subtract handles subtraction operation.
func (c Application) Subtract(a, b float64) revel.Result {
    difference := a - b
    return c.RenderJSON(SumResult{Result: difference})
}

// Multiply handles multiplication operation.
func (c Application) Multiply(a, b float64) revel.Result {
    product := a * b
    return c.RenderJSON(SumResult{Result: product})
}

// Divide handles division operation.
func (c Application) Divide(a, b float64) revel.Result {
    if b == 0 {
        // Handle division by zero error.
        return c.RenderError(400, "Cannot divide by zero.")
    }
    quotient := a / b
    return c.RenderJSON(SumResult{Result: quotient})
}

// SumResult is a structure to return results from math operations.
type SumResult struct {
    Result float64 `json:"result"`
    Sum    float64 `json:"sum"`
}

// main function initializes the Revel framework and starts the server.
func main() {
    revel.Import Routes // Import routes defined by the application.
    revel.Run() // Start the Revel application.
}

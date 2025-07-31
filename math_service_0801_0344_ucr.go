// 代码生成时间: 2025-08-01 03:44:26
package controllers

import (
    "encoding/json"
    "math"
    "revel"
)

// MathService controller provides basic mathematical operations.
type MathService struct {
    revel.Controller
}

// Add performs addition of two numbers.
func (c MathService) Add(a, b float64) revel.Result {
    result, err := add(a, b)
    if err != nil {
        return c.RenderError(err)
    }
    return c.RenderJSON(map[string]float64{
        "result": result,
    })
}

// Subtract performs subtraction of two numbers.
func (c MathService) Subtract(a, b float64) revel.Result {
    result, err := subtract(a, b)
    if err != nil {
        return c.RenderError(err)
    }
    return c.RenderJSON(map[string]float64{
        "result": result,
    })
}

// Multiply performs multiplication of two numbers.
func (c MathService) Multiply(a, b float64) revel.Result {
    result, err := multiply(a, b)
    if err != nil {
        return c.RenderError(err)
    }
    return c.RenderJSON(map[string]float64{
        "result": result,
    })
}

// Divide performs division of two numbers.
func (c MathService) Divide(a, b float64) revel.Result {
    result, err := divide(a, b)
    if err != nil {
        return c.RenderError(err)
    }
    return c.RenderJSON(map[string]float64{
        "result": result,
    })
}

// add adds two numbers and returns the result, or an error if one occurs.
func add(a, b float64) (float64, error) {
    return a + b, nil
}

// subtract subtracts b from a and returns the result, or an error if one occurs.
func subtract(a, b float64) (float64, error) {
    if b > a {
        return 0, fmt.Errorf("subtraction result cannot be negative")
    }
    return a - b, nil
}

// multiply multiplies two numbers and returns the result, or an error if one occurs.
func multiply(a, b float64) (float64, error) {
    return a * b, nil
}

// divide divides a by b and returns the result, or an error if one occurs.
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero is undefined")
    }
    return a / b, nil
}
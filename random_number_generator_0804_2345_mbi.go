// 代码生成时间: 2025-08-04 23:45:09
package controllers

import (
    "encoding/json"
    "math/rand"
    "time"

    "github.com/revel/revel"
)

// RandomNumberController handles requests for generating random numbers.
type RandomNumberController struct {
    revel.Controller
}

// NewRandom generates a random number between 1 and 100.
func (c RandomNumberController) NewRandom() revel.Result {
    // Seed the random number generator.
    rand.Seed(time.Now().UnixNano())

    // Generate a random number between 1 and 100.
    randomNumber := rand.Intn(100) + 1

    // Create a response object.
    response := struct {
        Number int `json:"number"`
    }{
        randomNumber,
    }

    // Marshal the response object to JSON.
    responseBytes, err := json.Marshal(response)
    if err != nil {
        // Handle any marshaling errors.
        c.RenderError(revel.StatusInternalServerError, err)
    }

    // Return the JSON response.
    return c.RenderJSON(responseBytes)
}

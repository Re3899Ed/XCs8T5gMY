// 代码生成时间: 2025-09-01 02:25:39
package controllers

import (
    "crypto/rand"
    "encoding/binary"
    "math/big"
    
    "github.com/revel/revel"
)

// RandomNumber is a controller for generating random numbers.
type RandomNumber struct {
    revel.Controller
}

// Generate returns a random number between 1 and the given max value.
// @Title Generate Random Number
// @Description Generates a random number between 1 and the given max value.
// @Param max query int true "The maximum value for the random number."
// @Success 200 {string} string "The generated random number."
// @Failure 400 {string} string "Invalid input."
// @Failure 500 {string} string "Internal server error."
// @Router /random-number/generate [get]
func (c RandomNumber) Generate(max int) revel.Result {
    if max < 1 {
        // If max is less than 1, return a bad request error.
        return c.RenderError(revel.StatusBadRequest, "Invalid input: max must be greater than 0.")
    }

    // Generate a random number between 1 and max.
    randomNumber, err := generateRandom(max)
    if err != nil {
        // If an error occurs, return an internal server error.
        return c.RenderError(revel.StatusInternalServerError, err.Error())
    }

    // Return the random number as a string.
    return c.RenderJSON(RandomNumberResult{
        Number: int64(randomNumber),
    })
}

// generateRandom generates a random number between 1 and max using crypto/rand.
// It returns the generated number and any potential error.
func generateRandom(max int) (int, error) {
    // Create a big integer with max+1 bits to store the random number.
    maxInt := big.NewInt(int64(max) + 1)
    randInt, err := rand.Int(rand.Reader, maxInt)
    if err != nil {
        return 0, err
    }

    // Since we want a number between 1 and max, we subtract 1 from the result.
    return int(randInt.Int64() - 1), nil
}

// RandomNumberResult is the JSON response for the random number generation.
type RandomNumberResult struct {
    Number int64 `json:"number"`
}

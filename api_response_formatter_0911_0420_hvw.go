// 代码生成时间: 2025-09-11 04:20:55
package main

import (
    "fmt"
    "revel"
)

// ApiResponseFormatter is a struct that holds the data for API responses.
type ApiResponseFormatter struct {
    Data    interface{} `json:"data"`
    Message string    `json:"message"`
    Status  int       `json:"status"`
}

// NewApiResponseFormatter creates a new ApiResponseFormatter instance with default values.
func NewApiResponseFormatter(data interface{}, message string, status int) ApiResponseFormatter {
    return ApiResponseFormatter{
        Data:    data,
        Message: message,
        Status:  status,
    }
}

// Controller is a Revel controller that uses ApiResponseFormatter.
type Controller struct {
    revel.Controller
}

// OkResponse responds with a successful status and data.
func (c Controller) OkResponse(data interface{}) revel.Result {
    return c.RenderJson(NewApiResponseFormatter(data, "OK", 200))
}

// ErrorResponse responds with an error status and message.
func (c Controller) ErrorResponse(message string, status int) revel.Result {
    return c.RenderJson(NewApiResponseFormatter(nil, message, status))
}

// Index is the action that serves the index page.
func (c Controller) Index() revel.Result {
    // Example usage of OkResponse
    return c.OkResponse(map[string]string{"hello": "world"})
}

func init() {
    // Filters are run before each action
    revel.Filters = []revel.Filter{
        revel.PanicFilter, // Recover from panics and display an error page
        revel.ActionInvoker,
    }
    // Registration of filters and results
    revel.RegisterAction(Controller.Index, revel.MethodGet)
}

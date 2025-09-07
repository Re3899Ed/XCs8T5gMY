// 代码生成时间: 2025-09-08 06:28:06
package main

import (
    "encoding/json"
    "github.com/revel/revel"
)

// ThemeSwitcherController is the controller for handling theme switching.
type ThemeSwitcherController struct {
    *revel.Controller
}

// SetTheme action sets the theme and saves it to the session.
func (c ThemeSwitcherController) SetTheme(theme string) revel.Result {
    // Check if the theme is valid
    validThemes := []string{"light", "dark"}
    if !contains(validThemes, theme) {
        return c.RenderError(revel.NewError("Invalid theme."), 400)
    }

    // Save the theme to the session
    c.Session["theme"] = theme

    // Return a success message with the new theme
    return c.RenderJson(map[string]string{"message": "Theme set successfully.", "theme": theme})
}

// GetCurrentTheme action retrieves the current theme from the session.
func (c ThemeSwitcherController) GetCurrentTheme() revel.Result {
    // Retrieve the theme from the session, default to "light" if not set
    theme, ok := c.Session["theme"].(string)
    if !ok {
        theme = "light"
    }

    // Return the current theme
    return c.RenderJson(map[string]string{"theme": theme})
}

// contains checks if a string is present in a slice of strings.
func contains(s []string, str string) bool {
    for _, v := range s {
        if v == str {
            return true
        }
    }
    return false
}

// RenderError is a custom render method that sends an error response.
func (c *ThemeSwitcherController) RenderError(err error, statusCode int) revel.Result {
    resp := revel.NewResponse()
    resp.Status = statusCode
    resp.Header["Content-Type"] = "application/json"
    resp.WriteJson(map[string]string{"error": err.Error()})
    return resp
}

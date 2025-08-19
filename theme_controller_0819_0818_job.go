// 代码生成时间: 2025-08-19 08:18:15
package controllers

import (
    "encoding/json"
    "github.com/revel/revel"
)

// ThemeController is responsible for handling theme-related requests.
type ThemeController struct {
    App *revel.Controller
}

// SwitchTheme changes the user's theme preference.
// @Route GET /theme/switch
// @PathParam theme string
// @PathParam user_id int
// @Return  200
func (c ThemeController) SwitchTheme(theme string, user_id int) revel.Result {
    // Check for valid theme
    validThemes := []string{"light", "dark"}
    for _, v := range validThemes {
        if v == theme {
            // Store the theme preference in the user's session or cookie
            c.App.Session["theme"] = theme
            return c.RenderJSON(revel.Result{
                Code: 200,
                Message: "Theme switched successfully",
                Data: theme,
            })
        }
    }
    // Return an error if the theme is not valid
    return c.RenderJSON(revel.Result{
        Code: 400,
        Message: "Invalid theme provided",
        Data: "",
    })
}

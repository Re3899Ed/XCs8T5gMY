// 代码生成时间: 2025-09-06 21:34:47
package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/revel/revel"
    "github.com/revel/revel/modules/auth" // Revel's built-in authentication
)

// AuthService handles user authentication
type AuthService struct {
    *revel.Controller
    auth *auth.Auth
}

// Login attempts to authenticate a user and provide an authentication token
func (c *AuthService) Login() revel.Result {
    var loginData struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    // Decode the JSON body into loginData
    if err := json.NewDecoder(c.Request.Request.Body).Decode(&loginData); err != nil {
        return c.RenderError(err)
    }

    // Attempt to authenticate the user
    session, err := c.auth.Authenticate(c.Request, loginData.Username, loginData.Password)
    if err != nil {
        return c.RenderError(err)
    }

    // If authentication is successful, return a JSON response
    // with the authentication token
    return c.RenderJSON(map[string]string{
        "token": session.ID,
    })
}

// Logout invalidates the current session, effectively logging the user out
func (c *AuthService) Logout() revel.Result {
    c.auth.Logout(c.Request, c.Response)
    return c.RenderJSON(map[string]string{
        "message": "Logged out successfully",
    })
}

// RenderError provides a common error rendering function
func (c *AuthService) RenderError(err error) revel.Result {
    return c.RenderJSON(map[string]string{
        "error": err.Error(),
    })
}

// Register initializes the AuthService with authentication module
func init() {
    revel.Filters.Append(revel.PrepareFilter, func(c *revel.Controller, fc []revel.Filter) {
        c.Authenticate = &auth.Auth{
            // Configure authentication settings
            SessionName:  "revel.auth",
            SessionKey:   "revel.auth.session",
            Realm:        "revel",
            LoginUrl:     "/auth/login",
            LogoutUrl:    "/auth/logout",
            UserPrefix:   "auth.",
            LoginHandler:  &AuthService{Controller: c}.Login,
            LogoutHandler: &AuthService{Controller: c}.Logout,
        }
    })
}

// 代码生成时间: 2025-09-03 07:50:04
package controllers

import (
    "encoding/json"
    "errors"
    "github.com/revel/revel"
    "app/models" // 假设有一个models包包含用户和权限模型
)

// PermissionsController 处理用户权限相关请求
type PermissionsController struct {
    App
}

// AddUser 添加新用户
func (c PermissionsController) AddUser(user models.User) revel.Result {
    // 检查用户是否已存在
    if models.UserExists(user.Username) {
        return c.RenderError(409, "User already exists")
    }
    // 添加用户到数据库
    if err := models.AddUser(user); err != nil {
        return c.RenderError(500, err.Error())
    }
    return c.RenderJSON(user)
}

// UpdateUser 更新用户信息
func (c PermissionsController) UpdateUser(user models.User) revel.Result {
    // 检查用户是否存在
    if !models.UserExists(user.Username) {
        return c.RenderError(404, "User not found")
    }
    // 更新用户信息
    if err := models.UpdateUser(user); err != nil {
        return c.RenderError(500, err.Error())
    }
    return c.RenderJSON(user)
}

// DeleteUser 删除用户
func (c PermissionsController) DeleteUser(username string) revel.Result {
    // 检查用户是否存在
    if !models.UserExists(username) {
        return c.RenderError(404, "User not found")
    }
    // 删除用户
    if err := models.DeleteUser(username); err != nil {
        return c.RenderError(500, err.Error())
    }
    return c.RenderJSON(map[string]string{"message": "User deleted"})
}

// AddRole 为用户添加角色
func (c PermissionsController) AddRole(username string, role models.Role) revel.Result {
    // 检查用户是否存在
    if !models.UserExists(username) {
        return c.RenderError(404, "User not found")
    }
    // 为用户添加角色
    if err := models.AddRole(username, role); err != nil {
        return c.RenderError(500, err.Error())
    }
    return c.RenderJSON(map[string]string{"message": "Role added"})
}

// RemoveRole 从用户移除角色
func (c PermissionsController) RemoveRole(username string, role models.Role) revel.Result {
    // 检查用户是否存在
    if !models.UserExists(username) {
        return c.RenderError(404, "User not found")
    }
    // 从用户移除角色
    if err := models.RemoveRole(username, role); err != nil {
        return c.RenderError(500, err.Error())
    }
    return c.RenderJSON(map[string]string{"message": "Role removed"})
}

// RenderError 渲染错误信息
func (c PermissionsController) RenderError(code int, message string) revel.Result {
    c.Response.Status = code
    return c.RenderJSON(map[string]string{"error": message})
}

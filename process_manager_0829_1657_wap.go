// 代码生成时间: 2025-08-29 16:57:32
package main

import (
    "fmt"
    "os/exec"
    "revel"
    "strings"
)

// ProcessManager 用于管理进程
type ProcessManager struct {
    Name string
}

// StartProcess 开始一个新的进程
func (pm *ProcessManager) StartProcess(command string) (*exec.Cmd, error) {
    cmd := exec.Command("/bin/sh", "-c", command)
    if err := cmd.Start(); err != nil {
        return nil, err
    }
    return cmd, nil
}

// StopProcess 停止一个正在运行的进程
func (pm *ProcessManager) StopProcess(cmd *exec.Cmd) error {
    if err := cmd.Process.Kill(); err != nil {
        return err
    }
    return nil
}

// ProcessController 控制器处理进程管理请求
type ProcessController struct {
    *revel.Controller
}

// StartAction 开始一个新进程
func (c *ProcessController) StartAction() revel.Result {
    command := c.Params.Get("command")
    if command == "" {
        return c.RenderError(revel.NewError("No command provided"))
    }
    pm := &ProcessManager{}
    cmd, err := pm.StartProcess(command)
    if err != nil {
        return c.RenderError(revel.NewError("Failed to start process"))
    }
    return c.RenderJSON(map[string]interface{}{"pid": cmd.Process.Pid})
}

// StopAction 停止一个进程
func (c *ProcessController) StopAction() revel.Result {
    pidStr := c.Params.Get("pid")
    pid, err := strconv.Atoi(pidStr)
    if err != nil {
        return c.RenderError(revel.NewError("Invalid PID"))
    }
    process, err := os.FindProcess(pid)
    if err != nil {
        return c.RenderError(revel.NewError("Process not found"))
    }
    pm := &ProcessManager{}
    err = pm.StopProcess(process)
    if err != nil {
        return c.RenderError(revel.NewError("Failed to stop process"))
    }
    return c.RenderJSON(map[string]interface{}{"status": "Process stopped"})
}

func init() {
    revel.InterceptFunc(revel.PrepareGzip, revel.BEFORE, "*/*")
}

func main() {
    revel.Run()
}
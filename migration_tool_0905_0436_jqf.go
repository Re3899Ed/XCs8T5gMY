// 代码生成时间: 2025-09-05 04:36:23
package main

import (
    "fmt"
    "log"
    "os"
    "revel"
    "github.com/robfig/revel/modules/db/app/db"
    "github.com/mattes/migrate"
    _ "github.com/mattes/migrate/database/postgres" // 引入postgres驱动
    _ "github.com/mattes/migrate/source/file"            // 引入文件源
)

// 迁移工具结构
type MigrationTool struct {
    *revel.Controller
}

// ApplyMigrations 应用所有未应用的数据库迁移
func (c *MigrationTool) ApplyMigrations() revel.Result {
    driver, err := migrate.NewPostgresDriver(os.Getenv("DATABASE_URL"))
    if err != nil {
        return c.RenderError(err)
    }
    defer driver.Close()

    migrations := &migrate.FileMigrationSource{
        Dir: "migrations",
    }

    m, err := migrate.NewWithDatabaseInstance(
        "file://migrations",
        driver,
        driver)
    if err != nil {
        return c.RenderError(err)
    }

    err = m.Up()
    if err != nil {
        return c.RenderError(err)
    }

    return c.RenderJSON(map[string]string{"status": "migrations applied"})
}

// RollbackMigrations 回滚最后一批应用的数据库迁移
func (c *MigrationTool) RollbackMigrations() revel.Result {
    driver, err := migrate.NewPostgresDriver(os.Getenv("DATABASE_URL"))
    if err != nil {
        return c.RenderError(err)
    }
    defer driver.Close()

    migrations := &migrate.FileMigrationSource{
        Dir: "migrations",
    }

    m, err := migrate.NewWithDatabaseInstance(
        "file://migrations",
        driver,
        driver)
    if err != nil {
        return c.RenderError(err)
    }

    err = m.Down()
    if err != nil {
        return c.RenderError(err)
    }

    return c.RenderJSON(map[string]string{"status": "migrations rolled back"})
}

// RenderError 渲染错误信息
func (c *MigrationTool) RenderError(err error) revel.Result {
    return c.RenderJSON(map[string]string{"error": err.Error()})
}

func init() {
    revel.Filters = []revel.Filter{
        revel.PanicFilter,
        revel.ActionInvoker,
    }
}

func main() {
    revel.Run()
}
// 代码生成时间: 2025-09-15 20:39:06
package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "os"
    "path/filepath"
    "revel"
)

// ConfigManager is the struct that handles configuration file operations.
type ConfigManager struct {
    Path string // The path to the configuration directory.
}

// NewConfigManager creates a new instance of ConfigManager with the provided path.
func NewConfigManager(path string) *ConfigManager {
    return &ConfigManager{
        Path: path,
    }
}

// LoadConfig loads the configuration from the file at the specified path.
func (cm *ConfigManager) LoadConfig(filename string) (map[string]interface{}, error) {
    fullPath := filepath.Join(cm.Path, filename)
    file, err := os.Open(fullPath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var config map[string]interface{}
    if err := json.NewDecoder(file).Decode(&config); err != nil {
        return nil, err
    }

    return config, nil
}

// SaveConfig saves the provided configuration to a file at the specified path.
func (cm *ConfigManager) SaveConfig(filename string, config map[string]interface{}) error {
    fullPath := filepath.Join(cm.Path, filename)
    file, err := os.Create(fullPath)
    if err != nil {
        return err
    }
    defer file.Close()

    if err := json.NewEncoder(file).Encode(config); err != nil {
        return err
    }

    return nil
}

// main function to demonstrate the usage of ConfigManager.
func main() {
    // Define the path to the configuration directory.
    configPath := "./configs"
    // Create a new ConfigManager instance.
    cm := NewConfigManager(configPath)

    // Define a sample configuration.
    sampleConfig := map[string]interface{}{
        "database": "users.db",
        "host": "localhost",
        "port": 5432,
    }

    // Load the configuration from a file named "app_config.json".
    config, err := cm.LoadConfig("app_config.json")
    if err != nil {
        revel.ERROR.Printf("Failed to load configuration: %v", err)
        return
    }
    fmt.Printf("Loaded configuration: %+v
", config)

    // Save the sample configuration to a file named "new_config.json".
    if err := cm.SaveConfig("new_config.json", sampleConfig); err != nil {
        revel.ERROR.Printf("Failed to save configuration: %v", err)
        return
    }
    fmt.Println("Configuration saved successfully.")
}

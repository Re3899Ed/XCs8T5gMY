// 代码生成时间: 2025-08-26 19:55:07
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "regexp"
    "strings"

    "github.com/revel/revel"
)

// TextFileAnalyzer 是一个分析文本文件的工具
type TextFileAnalyzer struct{}

// 分析文件的函数
func (tfa *TextFileAnalyzer) AnalyzeFile(filePath string) (*AnalysisResult, error) {
    content, err := ioutil.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

    // 将文件内容转换为字符串
    contentStr := string(content)

    // 统计单词数量
    wordCount := CountWords(contentStr)

    // 提取文本中的URL
    urls := ExtractURLs(contentStr)

    // 创建分析结果对象
    result := &AnalysisResult{
        FilePath:    filePath,
        WordCount:   wordCount,
        Urls:        urls,
    }

    return result, nil
}

// 分析结果的结构体
type AnalysisResult struct {
    FilePath    string   `json:"filePath"`
    WordCount   int      `json:"wordCount"`
    Urls        []string `json:"urls"`
}

// CountWords 统计单词数量
func CountWords(text string) int {
    // 使用正则表达式分割单词
    words := regexp.MustCompile(`\W+`).Split(strings.ToLower(text), -1)
    // 忽略空字符串
    words = filterEmpty(strings.Fields(strings.Join(words, " ")))
    return len(words)
}

// ExtractURLs 从文本中提取URL
func ExtractURLs(text string) []string {
    var urls []string
    regex := regexp.MustCompile(`http[s]?://[^\s]+`)
    matches := regex.FindAllString(text, -1)
    urls = matches
    return urls
}

// filterEmpty 过滤空字符串
func filterEmpty(strs []string) []string {
    var result []string
    for _, str := range strs {
        if str != "" {
            result = append(result, str)
        }
    }
    return result
}

func main() {
    // 启动 Revel 框架
    revel.Run(
        // 配置 Revel
        map[string]string{
            "revel.import.path": "path/to/your/app",
            "revel.mode": "dev",
        },
    )
}

// 以下是一个简单的 Revel 控制器示例
type TextAnalysis struct {
    *revel.Controller
}

// AnalyzeAction 分析指定的文件
func (c *TextAnalysis) AnalyzeAction(filePath string) revel.Result {
    analyzer := &TextFileAnalyzer{}
    result, err := analyzer.AnalyzeFile(filePath)
    if err != nil {
        return c.RenderError(err)
    }

    // 将结果转换为JSON并返回
    return c.RenderJSON(result)
}

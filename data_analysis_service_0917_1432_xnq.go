// 代码生成时间: 2025-09-17 14:32:06
package main

import (
    "github.com/revel/revel"
    "math"
    "strings"
)

// AnalysisService contains methods for analyzing data
type AnalysisService struct {
    *revel.Controller
}

// NewAnalysisService creates a new instance of AnalysisService
func NewAnalysisService() *AnalysisService {
    return &AnalysisService{
        Controller: &revel.Controller{},
    }
}

// AnalyzeData takes a string of comma-separated values (CSV) and returns statistical data
// such as mean, median, and standard deviation.
func (service *AnalysisService) AnalyzeData(csv string) revel.Result {
    values, err := parseCSV(csv)
    if err != nil {
        return service.RenderError(err)
    }

    mean := calculateMean(values)
    median := calculateMedian(values)
    standardDeviation := calculateStandardDeviation(values)

    return service.RenderJSON(map[string]float64{
        "mean": mean,
        "median": median,
        "standardDeviation": standardDeviation,
    })
}

// parseCSV converts a string of comma-separated values into a slice of float64
func parseCSV(csv string) ([]float64, error) {
    values := strings.Split(csv, ",")
    result := make([]float64, 0, len(values))
    for _, value := range values {
        num, err := strconv.ParseFloat(value, 64)
        if err != nil {
            return nil, err
        }
        result = append(result, num)
    }
    return result, nil
}

// calculateMean calculates the mean of a slice of float64 values
func calculateMean(values []float64) float64 {
    sum := 0.0
    for _, value := range values {
        sum += value
    }
    return sum / float64(len(values))
}

// calculateMedian calculates the median of a slice of float64 values
func calculateMedian(values []float64) float64 {
    sort.Float64s(values)
    length := len(values)
    mid := length / 2
    if length%2 == 0 {
        return (values[mid-1] + values[mid]) / 2
    }
    return values[mid]
}

// calculateStandardDeviation calculates the standard deviation of a slice of float64 values
func calculateStandardDeviation(values []float64) float64 {
    mean := calculateMean(values)
    var sum float64
    for _, value := range values {
        sum += math.Pow(value-mean, 2)
    }
    return math.Sqrt(sum / float64(len(values)-1))
}

// RenderError renders an error response with a JSON payload
func (service *AnalysisService) RenderError(err error) revel.Result {
    return service.RenderJSON(map[string]string{
        "error": err.Error(),
    },
        revel.StatusNotFound,
    )
}

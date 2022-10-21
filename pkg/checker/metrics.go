package checker

import (
	"strings"
)

type tokenMetrics struct {
	metricsMap map[string]int `json:"metrics_map"`
}

func getMetrics(rawTokens string) tokenMetrics {
	tokens := strings.Split(rawTokens, "|")

	metrics := tokenMetrics{
		metricsMap: make(map[string]int),
	}
	for i := range tokens {
		metrics.metricsMap[tokens[i]]++
	}

	return metrics
}

func MetricsCheck(tokens1, tokens2 string) float64 {
	metrics1 := getMetrics(tokens1)
	metrics2 := getMetrics(tokens2)

	var keys []string
	keys1 := make([]string, 0)
	for key := range metrics1.metricsMap {
		keys1 = append(keys1, key)
	}
	keys2 := make([]string, 0)
	for key := range metrics2.metricsMap {
		keys2 = append(keys2, key)
	}
	if len(keys1) > len(keys2) {
		keys = keys1
	} else {
		keys = keys2
	}
	res := 0.0
	for _, key := range keys {
		if metrics1.metricsMap[key] > metrics2.metricsMap[key] {
			res += float64(metrics2.metricsMap[key]) / float64(metrics1.metricsMap[key])
		} else {
			res += float64(metrics1.metricsMap[key]) / float64(metrics2.metricsMap[key])
		}
	}

	return res * 100 / float64(len(keys))
}

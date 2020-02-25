package math

import (
	"fmt"
	"strconv"
)

// Average is responsible to calculate a average of values
func Average(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	average := total / float64(len(numbers))
	roundAverage, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", average), 64)
	return roundAverage
}

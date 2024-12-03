package helpers

import (
	"fmt"
	"time"
)

func MeasureTime(start time.Time) {
	endTime := time.Now()
	elapsed := endTime.Sub(start)
	fmt.Printf("Execution time: %v\n", elapsed)
}

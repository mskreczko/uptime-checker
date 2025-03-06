package pkg

import (
	"fmt"
	"math"
	"time"
)

func Retry[T any](attempts int, fn func() (T, error)) (T, error) {
	var result T
	var err error

	baseDelay := 1 * time.Second

	for i := 0; i < attempts; i++ {
		result, err = fn()
		if err == nil {
			return result, nil
		}

		step := time.Duration(math.Pow(2, float64(i)))
		delay := baseDelay * step
		time.Sleep(delay)
	}
	return result, fmt.Errorf("failed after %d attempts: %w", attempts, err)
}

package utils

import (
	"fmt"
	"time"
)

type Executable func() error

func Retry(fn Executable, maxRetry int, duration int) {
	for i := 0; i < maxRetry; i++ {
		err := fn()
		if err == nil {
			break
		}
		fmt.Printf("retrying function %d out of %d\n", i+1, maxRetry)
		time.Sleep(time.Duration(duration) * time.Second)
	}
}

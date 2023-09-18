package utils

import (
	"fmt"
	"time"
)

func DoWithTries(fn func() error, attempts int, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = fn(); err != nil {
			fmt.Printf("Trying to connect to PostgreSQL: %d attempts left", attempts)
			time.Sleep(delay)
			attempts--

			continue
		}

		return nil
	}

	return
}

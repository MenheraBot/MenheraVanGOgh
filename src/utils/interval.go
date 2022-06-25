package utils

import "time"

func SetInterval(someFunc func(), milliseconds int) chan bool {

	interval := time.Duration(milliseconds) * time.Millisecond

	ticker := time.NewTicker(interval)
	clear := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				someFunc()
			case <-clear:
				ticker.Stop()
				return
			}
		}
	}()

	return clear
}

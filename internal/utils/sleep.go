package utils

import (
	"sync"
	"time"
)

func AsyncSleep() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(2 * time.Second)
			wg.Done()
		}()
	}
	wg.Wait()
}

func SyncSleep() {
	for i := 0; i < 10; i++ {
		func() {
			time.Sleep(2 * time.Second)
		}()
	}
}

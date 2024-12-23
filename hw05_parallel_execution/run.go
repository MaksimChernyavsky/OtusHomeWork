package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	i := 0      // current task index to execute
	errNum := 0 // current count of errors

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for j := 0; j < n; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				mu.Lock()
				if i == len(tasks) || errNum >= m {
					mu.Unlock()
					return
				}
				task := tasks[i]
				i++
				mu.Unlock()
				err := task()
				mu.Lock()
				if err != nil {
					errNum++
				}
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	if errNum >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}

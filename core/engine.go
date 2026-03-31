package core

import (
	"sync"
)

func RunEngine(jobs []Job, workers int) string {
	jobChan := make(chan Job)
	resultChan := make(chan string)

	var wg sync.WaitGroup

	// Workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobChan {
				found, result := Process(job)
				if found {
					resultChan <- result
					return
				}
			}
		}()
	}

	// Enviar jobs
	go func() {
		for _, job := range jobs {
			jobChan <- job
		}
		close(jobChan)
	}()

	// Encerrar resultChan
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Captura resultado
	for res := range resultChan {
		return res
	}

	return ""
}
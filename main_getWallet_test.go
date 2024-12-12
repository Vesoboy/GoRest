package main

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestGetRPSLoad(t *testing.T) {
	testGetRPS(100, 1*time.Second)
}

func testGetRPS(rps int, duration time.Duration) {
	var wg sync.WaitGroup
	var successCount, failureCount int

	client := &http.Client{}
	ticker := time.NewTicker(time.Second / time.Duration(rps))
	defer ticker.Stop()

	// Канал для завершения генерации запросов
	done := make(chan bool)

	fmt.Printf("Starting RPS test with %d RPS for %v\n", rps, duration)

	// Запуск генерации запросов
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				wg.Add(1)
				go func() {
					defer wg.Done()
					resp, err := client.Get("http://localhost:8080/api/v1/wallets/0990b60d-405b-4798-a79a-197983835325")
					if err != nil {
						fmt.Printf("Error: %v\n", err)
						failureCount++
						return
					}
					defer resp.Body.Close()

					if resp.StatusCode == http.StatusOK {
						successCount++
					} else {
						fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
						failureCount++
					}
				}()
			}
		}
	}()

	// Таймер для завершения теста
	time.Sleep(duration)
	close(done)
	wg.Wait()

	// Вывод результатов
	fmt.Printf("Test completed: Success=%d, Failures=%d\n", successCount, failureCount)
}

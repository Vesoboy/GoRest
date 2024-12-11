// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// 	"testing"
// 	"time"
// )

// func TestRPSLoad(t *testing.T) {
// 	testRPS(1000, 10*time.Second) // Измените URL на нужный
// }

// func testRPS(rps int, duration time.Duration) {
// 	var wg sync.WaitGroup
// 	var successCount, failureCount int

// 	client := &http.Client{}
// 	ticker := time.NewTicker(time.Second / time.Duration(rps))
// 	defer ticker.Stop()

// 	// Канал для завершения генерации запросов
// 	done := make(chan bool)

// 	fmt.Printf("Starting RPS test with %d RPS for %v\n", rps, duration)

// 	// Запуск генерации запросов
// 	go func() {
// 		for {
// 			select {
// 			case <-done:
// 				return
// 			case <-ticker.C:
// 				wg.Add(1)
// 				go func() {
// 					defer wg.Done()
// 					resp, err := client.Get("http://localhost:8080/api/v1/wallets/a04b5495-6df5-4696-b235-edb9d37f12f3")
// 					if err != nil {
// 						fmt.Printf("Error: %v\n", err)
// 						failureCount++
// 						return
// 					}
// 					defer resp.Body.Close()

// 					if resp.StatusCode == http.StatusOK {
// 						successCount++
// 					} else {
// 						fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
// 						failureCount++
// 					}
// 				}()
// 			}
// 		}
// 	}()

// 	// Таймер для завершения теста
// 	time.Sleep(duration)
// 	close(done)
// 	wg.Wait()

// 	// Вывод результатов
// 	fmt.Printf("Test completed: Success=%d, Failures=%d\n", successCount, failureCount)
// }

package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestRPSLoad(t *testing.T) {
	testRPS("http://localhost:8080/api/v1/wallet", 1000, 10*time.Second)
}

func testRPS(url string, rps int, duration time.Duration) {
	var wg sync.WaitGroup
	var successCount, failureCount int

	requestBody := []byte(`{"valletId": "217ff946-1157-40ef-afa4-02d09d0964af", "operationType": "DEPOSIT", "amount": 1}`)
	client := &http.Client{}

	ticker := time.NewTicker(time.Second / time.Duration(rps))
	defer ticker.Stop()

	// Таймер для завершения теста
	endTime := time.Now().Add(duration)

	fmt.Printf("Starting RPS test with %d RPS for %v\n", rps, duration)

	for time.Now().Before(endTime) {
		<-ticker.C
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := client.Post(url, "application/json", bytes.NewBuffer(requestBody))
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				failureCount++
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				successCount++
			} else {
				failureCount++
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Test completed: Success=%d, Failures=%d\n", successCount, failureCount)
}

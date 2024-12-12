package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestPostRPSLoad(t *testing.T) {
	testPostRPS("http://localhost:8080/api/v1/wallet", 100, 1*time.Second)
}

func testPostRPS(url string, rps int, duration time.Duration) {
	var wg sync.WaitGroup
	var successCount, failureCount int

	requestBody := []byte(`{"valletId": "0990b60d-405b-4798-a79a-197983835325", "operationType": "DEPOSIT", "amount": 1}`)
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

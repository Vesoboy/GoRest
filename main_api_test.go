package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestWalletAPI(t *testing.T) {
	baseURL := "http://localhost:8080/api/v1"
	client := &http.Client{}

	// Шаг 1: Тест создания кошелька
	walletID := uuid.New()
	createWalletPayload := map[string]interface{}{
		"valletId": walletID,
		"allSum":   100.0,
	}
	payloadBytes, _ := json.Marshal(createWalletPayload)

	resp, err := client.Post(fmt.Sprintf("%s/addwallet", baseURL), "application/json", bytes.NewReader(payloadBytes))
	if err != nil {
		t.Fatalf("Ошибка при запросе создания кошелька: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Ожидался статус 201 Created, получен: %d", resp.StatusCode)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Созданный кошелек:", string(body))

	// Шаг 2: Тест пополнения баланса
	updateWalletPayload := map[string]interface{}{
		"valletId":      walletID,
		"operationType": "deposit",
		"amount":        50.0,
	}
	payloadBytes, _ = json.Marshal(updateWalletPayload)

	resp, err = client.Post(fmt.Sprintf("%s/wallet", baseURL), "application/json", bytes.NewReader(payloadBytes))
	if err != nil {
		t.Fatalf("Ошибка при запросе пополнения баланса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Ожидался статус 200 OK, получен: %d", resp.StatusCode)
	}

	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println("Пополнение баланса:", string(body))

	// Шаг 3: Тест получения баланса
	resp, err = client.Get(fmt.Sprintf("%s/wallets/%s", baseURL, walletID))
	if err != nil {
		t.Fatalf("Ошибка при запросе получения баланса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Ожидался статус 200 OK, получен: %d", resp.StatusCode)
	}

	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println("Текущий баланс:", string(body))
}

package main

import (
	_ "bytes"
	_ "encoding/json"
	"log"
	_ "log"
	"net/http"
	_ "net/http"
	"restApp/DataContext"
	restapi "restApp/RestApi"

	"github.com/gorilla/mux"
)

// main - запуск REST-сервера.
//
// Функция main инициализирует БД,
// создает mux-роутер,
// регистрирует на нем 3 эндпоинта:
// 1. POST /api/v1/addwallet - добавление кошелька;
// 2. POST /api/v1/wallet - обновление суммы на счете;
// 3. GET /api/v1/wallets/{valletId} - получение информации о кошельке;
// и запускает сервер на порту 8080.
func main() {

	db := DataContext.DataContextDB()
	muxRout := mux.NewRouter()

	muxRout.HandleFunc("/api/v1/addwallet", func(w http.ResponseWriter, r *http.Request) {
		restapi.AddWallet(w, r, db)
	}).Methods(http.MethodPost)

	muxRout.HandleFunc("/api/v1/wallet", func(w http.ResponseWriter, r *http.Request) {
		restapi.UpdateWallet(w, r, db)
	}).Methods(http.MethodPost)

	muxRout.HandleFunc("/api/v1/wallets/{valletId}", func(w http.ResponseWriter, r *http.Request) {

		restapi.GetWallets(w, r, db, mux.Vars(r)["valletId"])
	}).Methods(http.MethodGet)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", muxRout))
}

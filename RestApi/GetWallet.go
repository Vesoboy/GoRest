package restapi

import (
	"encoding/json"
	"net/http"
	"restApp/DataContext"

	"gorm.io/gorm"
)

// GetWallets - получение информации о кошельке.
//
// Метод ожидает ID кошелька в параметрах.
//
// Возвращает json-объект с информацией о кошельке, если запрос успешен.
// Если запрос не успешен, возвращает json-объект с ошибкой.
func GetWallets(w http.ResponseWriter, r *http.Request, db *gorm.DB, vars string) {

	if len(vars) == 0 {
		respondWithError(w, http.StatusInternalServerError, "Ошибка в данных http запроса")
		return
	}

	wallet, err := DataContext.GetWallets(db, vars)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Ошибка получения кошелька")
		return
	}

	json.NewEncoder(w).Encode(wallet.AllSum)
}

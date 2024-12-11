package restapi

import (
	"encoding/json"
	"net/http"

	"restApp/DataContext"
	"restApp/Models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AddWallet - добавление кошелька в базу данных.
//
// Метод ожидает в теле запроса json-объект со следующими полями:
// ValletId - ID кошелька (uuid). Если ID пустой, будет сгенерирован новый.
// AllSum - начальная сумма на счету.
//
// Возвращает json-объект с информацией о созданном кошельке, если запрос успешен.
// Если запрос не успешен, возвращает json-объект с ошибкой.
func AddWallet(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "application/json")

	var wallet Models.Wallets
	if err := json.NewDecoder(r.Body).Decode(&wallet); err != nil {
		respondWithError(w, http.StatusBadRequest, "Ошибка в данных http запроса")
		return
	}

	if wallet.ValletId == uuid.Nil {
		wallet.ValletId = uuid.New()
	}

	// Проверка на существование кошелька
	if walletExists(db, wallet.ValletId) {
		respondWithError(w, http.StatusBadRequest, "Кошелек с таким ID уже существует")
		return
	}

	// Добавление кошелька в базу данных
	if err := DataContext.AddWallet(db, wallet.ValletId, wallet.AllSum); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Ответ при успешном создании кошелька
	respondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "Кошелек успешно добавлен",
		"wallet":  wallet,
	})
}

func walletExists(db *gorm.DB, walletID uuid.UUID) bool {
	var existingWallet Models.Wallets
	result := db.Where("vallet_id = ?", walletID).First(&existingWallet)
	return result.RowsAffected > 0
}

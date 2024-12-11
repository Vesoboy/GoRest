package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"restApp/DataContext"
	"restApp/Models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// UpdateWallet - обновление суммы на счете.
//
// Метод ожидает в теле запроса json-объект со следующими полями:
// ValletId - ID кошелька (uuid).
// OperationType - тип операции (deposit или withdraw).
// Amount - сумма операции.
//
// Возвращает json-объект с информацией о кошельке, если запрос успешен.
// Если запрос не успешен, возвращает json-объект с ошибкой.
func UpdateWallet(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "application/json")

	var wallet Models.Wallets
	if err := json.NewDecoder(r.Body).Decode(&wallet); err != nil {
		respondWithError(w, http.StatusBadRequest, "Ошибка в данных запроса")
		return
	}

	if wallet.ValletId == uuid.Nil {
		respondWithError(w, http.StatusBadRequest, "ID кошелька не может быть пустым")
		return
	}

	if wallet.ValletId == uuid.Nil {
		respondWithError(w, http.StatusBadRequest, "ID кошелька не может быть пустым")
		return
	}

	upWallet, err := getWalletWithLock(db, wallet.ValletId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			respondWithError(w, http.StatusNotFound, "Кошелек не найден")
		} else {
			respondWithError(w, http.StatusInternalServerError, "Ошибка получения кошелька")
		}
		return
	}

	allSum, err := calculateNewBalance(upWallet.AllSum, wallet.OperationType, wallet.Amount)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := DataContext.UpdateWallet(db, wallet.ValletId, strings.ToLower(wallet.OperationType), wallet.Amount, allSum); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Ошибка обновления кошелька")
		return
	}

	upWallet.OperationType = strings.ToLower(wallet.OperationType)
	upWallet.Amount = wallet.Amount
	upWallet.AllSum = allSum

	respondWithJSON(w, http.StatusOK, upWallet)
}

func getWalletWithLock(db *gorm.DB, walletID uuid.UUID) (Models.Wallets, error) {
	var wallet Models.Wallets
	result := db.Session(&gorm.Session{PrepareStmt: true}).Model(&Models.Wallets{}).
		Where("vallet_id = ?", walletID).
		Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).
		First(&wallet)
	return wallet, result.Error
}

func calculateNewBalance(currentBalance float64, operationType string, amount float64) (float64, error) {
	switch strings.ToLower(operationType) {
	case Models.Deposit:
		return currentBalance + amount, nil
	case Models.Withdraw:
		if currentBalance < amount {
			return 0, fmt.Errorf("Недостаточно средств для вывода")
		}
		return currentBalance - amount, nil
	default:
		return 0, fmt.Errorf("Некорректный тип операции")
	}
}

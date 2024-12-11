package Models

import (
	"github.com/google/uuid"
)

type Wallets struct {
	ValletId      uuid.UUID `json:"valletId" gorm:"primaryKey";index`
	OperationType string    `json:"operationType"`
	Amount        float64   `json:"amount"`
	AllSum        float64   `json:"allSum"`
}

var PathGet = "/api/v1/wallets/"

const (
	Deposit  = "deposit"
	Withdraw = "withdraw"
)

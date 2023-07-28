package models

import "github.com/google/uuid"

type Deposit struct {
	UserID uuid.UUID `json:"user_id"`
	CoinID int       `json:"coin_id"`
	Wallet string    `json:"wallet"`
	Amount float64   `json:"amount"`
}

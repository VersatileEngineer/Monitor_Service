package models

import "github.com/google/uuid"

type Wallet struct {
	UserID      uuid.UUID `json:"user_id"`
	CoinID      int       `json:"coin_id"`
	PublicKey   string    `json:"public_key"`
	CryptoValue float64   `json:"crypto_value"`
	FiatValue   float64   `json:"fiat_value"`
}

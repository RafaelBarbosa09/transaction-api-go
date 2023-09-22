package models

type Account struct {
	ID             int64   `json:"id" db:"id"`
	AccountHolder  string  `json:"accountHolder" db:"account_holder"`
	ActiveCard     bool    `json:"activeCard" db:"active_card" `
	AvailableLimit float64 `json:"availableLimit" db:"available_limit"`
}

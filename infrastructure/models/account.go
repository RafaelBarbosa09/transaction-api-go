package models

type Account struct {
	ID             int64   `json:"id"`
	AccountHolder  string  `json:"accountHolder"`
	ActiveCard     bool    `json:"activeCard"`
	AvailableLimit float64 `json:"availableLimit"`
}

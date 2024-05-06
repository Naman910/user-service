package model

// User represents a user entity
type User struct {
	ID        uint64  `json:"id"`
	FirstName string  `json:"fname"`
	City      string  `json:"city"`
	Phone     uint64  `json:"phone"`
	Height    float32 `json:"height"`
	Married   bool    `json:"married"`
}

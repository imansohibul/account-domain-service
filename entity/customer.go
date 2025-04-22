package entity

import "time"

// Customer represents a customer entity
type Customer struct {
	ID          uint
	Fullname    string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

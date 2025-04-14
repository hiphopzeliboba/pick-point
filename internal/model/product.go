package model

import "time"

// Product представляет товар
type Product struct {
	ID         int
	IntakeID   int
	ReceivedAt time.Time
	Type       string
}

package entity

import "time"

const (
	NewStatus        = "NEW"
	ProcessingStatus = "PROCESSING"
	InvalidStatus    = "INVALID"
	ProcessedStatus  = "PROCESSED"
)

type Order struct {
	ID         string    `db:"id"`
	Status     string    `db:"status"`
	Accrual    float64   `db:"accrual"`
	UploadedAt time.Time `db:"uploaded_at"`
	UserID     string    `db:"user_id"`
}

func CreateNewOrder(
	id string,
	userID string,
) Order {
	return Order{
		ID:         id,
		Status:     NewStatus,
		Accrual:    0,
		UploadedAt: time.Now(),
		UserID:     userID,
	}
}

package entities

import "time"

var (
	StatusPending   = "pending"
	StatusActive    = "active"
	StatusExpired   = "expired"
	StatusCancelled = "cancelled"
)

type CreateSubscriptionRequest struct {
	SubscriptionID string `json:"subscription_id" binding:"required"`
	DurationInDays int    `json:"duration_in_days" binding:"required,gt=0"`
}

type UserSubscriptions struct {
	ID             string    `json:"id"`
	UserId         string    `json:"user_id"`
	SubscriptionId string    `json:"subscription_id"`
	StartAt        time.Time `json:"start_at"`
	EndAt          time.Time `json:"end_at"`
	Status         string    `json:"status"`
}

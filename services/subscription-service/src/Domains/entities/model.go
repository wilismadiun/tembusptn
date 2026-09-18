package entities

import "time"

type Subscriptions struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type UserSubscriptions struct {
	ID             string `json:"id"`
	UserId         string `json:"user_id"`
	SubscriptionId string `json:"subscription_id"`
	StartAt        time.Time
	EndAt          time.Time
	Status         string `json:"status"`
}

type UserContext struct {
	UserId string
	Role   string
}

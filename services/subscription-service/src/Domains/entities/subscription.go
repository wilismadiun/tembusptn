package entities

import "errors"

type Subscriptions struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

func VerifySubscription(subs Subscriptions) error {
	if subs.Name == "" {
		return errors.New("name is required")
	}

	if subs.Price == 0 {
		return errors.New("price is required")
	}

	if subs.Price < 0 {
		return errors.New("minimum price is 0")
	}

	return nil
}

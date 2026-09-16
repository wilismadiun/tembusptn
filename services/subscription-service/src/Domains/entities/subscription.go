package entities

import "errors"

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

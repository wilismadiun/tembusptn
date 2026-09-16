package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifySubscription(t *testing.T) {
	t.Run("should return error when name is empty", func(t *testing.T) {
		subs := Subscriptions{
			Name:  "",
			Price: 50000,
		}

		err := VerifySubscription(subs)

		assert.Error(t, err)
		assert.Equal(t, "name is required", err.Error())
	})

	t.Run("should return error when price is zero", func(t *testing.T) {
		subs := Subscriptions{
			Name:  "Gold",
			Price: 0,
		}

		err := VerifySubscription(subs)

		assert.Error(t, err)
		assert.Equal(t, "price is required", err.Error())
	})

	t.Run("should return nil when subscription is valid", func(t *testing.T) {
		subs := Subscriptions{
			Name:  "Gold",
			Price: 50000,
		}

		err := VerifySubscription(subs)

		assert.NoError(t, err)
	})
}

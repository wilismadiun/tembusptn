package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyUser(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr bool
	}{
		{
			name: "valid user",
			user: User{
				ID:       "123",
				Email:    "user@gmail.com",
				Name:     "John Doe",
				Phone:    "081234567890",
				Password: "password123",
			},
			wantErr: false,
		},
		{
			name: "empty id",
			user: User{
				ID:       "",
				Email:    "user@gmail.com",
				Name:     "John Doe",
				Phone:    "081234567890",
				Password: "password123",
			},
			wantErr: true,
		},
		{
			name: "empty email",
			user: User{
				ID:       "123",
				Email:    "",
				Name:     "John Doe",
				Phone:    "081234567890",
				Password: "password123",
			},
			wantErr: true,
		},
		{
			name: "invalid email",
			user: User{
				ID:       "123",
				Email:    "user",
				Name:     "John Doe",
				Phone:    "081234567890",
				Password: "password123",
			},
			wantErr: true,
		},
		{
			name: "empty name",
			user: User{
				ID:       "123",
				Email:    "user@gmail.com",
				Name:     "",
				Phone:    "081234567890",
				Password: "password123",
			},
			wantErr: true,
		},
		{
			name: "empty telephone",
			user: User{
				ID:       "123",
				Email:    "user@gmail.com",
				Name:     "John Doe",
				Phone:    "",
				Password: "password123",
			},
			wantErr: true,
		},
		{
			name: "invalid telephone",
			user: User{
				ID:       "123",
				Email:    "user@gmail.com",
				Name:     "John Doe",
				Phone:    "08123abc",
				Password: "password123",
			},
			wantErr: true,
		},
		{
			name: "empty password",
			user: User{
				ID:       "123",
				Email:    "user@gmail.com",
				Name:     "John Doe",
				Phone:    "081234567890",
				Password: "",
			},
			wantErr: true,
		},
		{
			name: "password less than 8 characters",
			user: User{
				ID:       "123",
				Email:    "user@gmail.com",
				Name:     "John Doe",
				Phone:    "081234567890",
				Password: "1234567",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := VerifyUser(tt.user)

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

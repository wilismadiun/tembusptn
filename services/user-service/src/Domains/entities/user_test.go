package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyUser(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr string
	}{
		{
			name: "id kosong",
			user: User{
				ID:       "",
				RoleId:   2,
				Email:    "jaya@gmail.com",
				Name:     "Jaya",
				Phone:    "081234567890",
				Password: "password123",
			},
			wantErr: "id is required",
		},
		{
			name: "email kosong",
			user: User{
				ID:       "user-123",
				RoleId:   2,
				Email:    "",
				Name:     "Jaya",
				Phone:    "081234567890",
				Password: "password123",
			},
			wantErr: "email is required",
		},
		{
			name: "format email salah",
			user: User{
				ID:       "user-123",
				RoleId:   2,
				Email:    "jaya@gmail",
				Name:     "Jaya",
				Phone:    "081234567890",
				Password: "password123",
			},
			wantErr: "invalid email format",
		},
		{
			name: "name kosong",
			user: User{
				ID:       "user-123",
				RoleId:   2,
				Email:    "jaya@gmail.com",
				Name:     "",
				Phone:    "081234567890",
				Password: "password123",
			},
			wantErr: "name is required",
		},
		{
			name: "telephone kosong",
			user: User{
				ID:       "user-123",
				RoleId:   2,
				Email:    "jaya@gmail.com",
				Name:     "Jaya",
				Phone:    "",
				Password: "password123",
			},
			wantErr: "telephone is required",
		},
		{
			name: "telephone tidak sesuai format",
			user: User{
				ID:       "user-123",
				RoleId:   2,
				Email:    "jaya@gmail.com",
				Name:     "Jaya",
				Phone:    "08123abc",
				Password: "password123",
			},
			wantErr: "telephone must contain 10 to 15 digits",
		},
		{
			name: "password kosong",
			user: User{
				ID:       "user-123",
				RoleId:   2,
				Email:    "jaya@gmail.com",
				Name:     "Jaya",
				Phone:    "081234567890",
				Password: "",
			},
			wantErr: "password is required",
		},
		{
			name: "password kurang dari 8 karakter",
			user: User{
				ID:       "user-123",
				RoleId:   2,
				Email:    "jaya@gmail.com",
				Name:     "Jaya",
				Phone:    "081234567890",
				Password: "1234567",
			},
			wantErr: "password must be at least 8 characters",
		},
		{
			name: "role id kosong atau 0",
			user: User{
				ID:       "user-123",
				RoleId:   0,
				Email:    "jaya@gmail.com",
				Name:     "Jaya",
				Phone:    "081234567890",
				Password: "password123",
			},
			wantErr: "role id is required",
		},
		{
			name: "sukses",
			user: User{
				ID:       "user-123",
				RoleId:   2,
				Email:    "jaya@gmail.com",
				Name:     "Jaya",
				Phone:    "081234567890",
				Password: "password123",
			},
			wantErr: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := VerifyUser(test.user)

			if test.wantErr == "" {
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
			assert.Equal(t, test.wantErr, err.Error())
		})
	}
}

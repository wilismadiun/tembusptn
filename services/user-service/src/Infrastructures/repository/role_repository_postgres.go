package repository

import (
	"github.com/wilismadiun/tembusptn/services/user-service/src/Domains/entities"
	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func (h *RoleRepository) FindRoleById(id int) (string, error) {
	var role entities.Role

	err := h.DB.Where("id = ?", id).First(&role).Error
	if err != nil {
		return "", err
	}

	return role.Name, nil
}

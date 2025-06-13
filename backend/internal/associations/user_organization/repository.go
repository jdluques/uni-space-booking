package user_organization

import "gorm.io/gorm"

type UserOrganizationRepository interface {
}

type userOrganizationRepository struct {
	db *gorm.DB
}

func NewUserOrganizationRepository(db *gorm.DB) UserOrganizationRepository {
	return &userOrganizationRepository{db: db}
}

package user_organization

import (
	"fmt"
	"gorm.io/gorm"
)

type UserOrganizationRepository interface {
	GetOrganizationIDsByUserID(userID string) ([]string, error)
	IsUserInOrganization(userID, orgID string) bool
}

type userOrganizationRepository struct {
	db *gorm.DB
}

func NewUserOrganizationRepository(db *gorm.DB) UserOrganizationRepository {
	return &userOrganizationRepository{db: db}
}

func (r *userOrganizationRepository) GetOrganizationIDsByUserID(userID string) ([]string, error) {
	var orgIDs []string
	if err := r.db.Table("user_organizations").Where("user_id = ?", userID).Pluck("organization_id", &orgIDs).Error; err != nil {
		return nil, fmt.Errorf("failed to find organization IDs of user %s: %w", userID, err)
	}
	return orgIDs, nil
}

func (r *userOrganizationRepository) IsUserInOrganization(userID, orgID string) bool {
	var count int64
	if err := r.db.Table("user_organizations").Where("user_id = ? AND organization_id = ?", userID, orgID).Count(&count).Error; err != nil {
		fmt.Printf("Error checking if user %s is in organization %s: %v\n", userID, orgID, err)
		return false
	}
	return count > 0
}

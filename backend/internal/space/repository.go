package space

import "gorm.io/gorm"

type SpaceRepository interface {
}

type spaceRepository struct {
	db *gorm.DB
}

func NewSpaceRepository(db *gorm.DB) SpaceRepository {
	return &spaceRepository{db: db}
}

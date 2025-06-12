package db

import "gorm.io/gorm"

func Migrate(db *gorm.DB, models ...interface{}) error {
	if err := db.AutoMigrate(models...); err != nil {
		return err
	}
	return nil
}

package account

import (
	"errors"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// Generate mysql DB repository
func NewMysqlRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Get user email by id
func (repo *repository) GetEmailByID(id int) (string, error) {
	var result string
	err := repo.db.Table("users").Select("email").Find(&result, "id", id).Error
	return result, err
}

// Update user email
func (repo *repository) UpdateEmail(id int, email string) error {
	query := repo.db.Table("users").Where("id", id).Update("email", email)
	if query.Error == nil && query.RowsAffected < 1 {
		return errors.New("not found")
	}
	return query.Error
}

// Update user password
func (repo *repository) UpdatePassword(id int, password string) error {
	query := repo.db.Table("users").Where("id", id).Update("password", password)
	if query.Error == nil && query.RowsAffected < 1 {
		return errors.New("not found")
	}
	return query.Error
}

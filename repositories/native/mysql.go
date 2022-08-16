package native

import (
	"plant-api/api/v1/native/response"
	"plant-api/business/native"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// Generate mysql DB repository
func NewMysqlRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Create new native and store into database
func (repo *repository) Create(native *native.Native) (uint, error) {
	if err := repo.db.Create(&native).Error; err != nil {
		return 0, err
	}
	return native.ID, nil
}

// Get all natives
func (repo *repository) GetAll() ([]response.Native, error) {
	natives := []response.Native{}
	if err := repo.db.Find(&natives).Error; err != nil {
		return nil, err
	}
	return natives, nil
}

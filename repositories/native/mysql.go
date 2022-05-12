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
func (repo *repository) Create(native *native.Native) error {
	if err := repo.db.Create(&native).Error; err != nil {
		return err
	}
	return nil
}

// Get all natives
func (repo *repository) GetAll() ([]response.Native, error) {
	natives := []response.Native{}
	if err := repo.db.Find(&natives).Error; err != nil {
		return nil, err
	}
	return natives, nil
}

// Get native by given id. It's return nil if not found
func (repo *repository) Get(id int) (*response.Native, error) {
	native := response.Native{}
	if err := repo.db.First(&native, id).Error; err != nil {
		return nil, err
	}
	return &native, nil
}

// Get native by given name. It's return nil if not found
func (repo *repository) GetByName(name string) (*native.Native, error) {
	native := native.Native{}
	if err := repo.db.Where("name", name).First(&native).Error; err != nil {
		return nil, err
	}
	return &native, nil
}
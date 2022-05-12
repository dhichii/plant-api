package plantnative

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

// Generate mysql DB repository
func NewMysqlRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Delete plant native by given id
func (repo *repository) Delete(id int) error {
	if err := repo.db.Table("plant_natives").Where("plant_id", id).Delete(id).Error; err != nil {
		return err
	}
	return nil
}
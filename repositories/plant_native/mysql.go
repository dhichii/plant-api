package plantnative

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

// Generate mysql DB repository
func NewMysqlRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Delete plant native by given plant id
func (repo *repository) Delete(plantID int) error {
	if err := repo.db.Table("plant_natives").Where("plant_id", plantID).Delete(plantID).Error; err != nil {
		return err
	}
	return nil
}
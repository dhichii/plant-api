package plant

import (
	"plant-api/api/v1/plant/response"
	"plant-api/business/plant"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// Generate mysql DB repository
func NewMysqlRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Type used to update plant
type plantModel plant.Plant

// Create new plant and store into database
func (repo *repository) Create(plant *plant.Plant) error {
	if err := repo.db.Create(plant).Error; err != nil {
		return err
	}
	return nil
}

// Get all plants
func (repo *repository) GetAll() ([]response.Plant, error) {
	plants := []response.Plant{}
	if err := repo.db.Find(&plants).Error; err != nil {
		return nil, err
	}
	return plants, nil
}

// Get all plants by given name
func (repo *repository) GetByName(name string) ([]response.Plant, error) {
	plant := []response.Plant{}
	if err := repo.db.Where("name LIKE '%" + name + "%'").Find(&plant).Error; err != nil {
		return nil, err
	}
	return plant, nil
}

// Get plant by given id
func (repo *repository) GetDetail(id int) (*response.PlantDetail, error) {
	plant := &response.PlantDetail{}
	if err := repo.db.Table("plants").First(&plant, id).Error; err != nil {
		return nil, err
	}
	// Find plant natives that related to plant
	natives := []*response.Native{}
	if err := repo.db.Table("natives").
		Joins("JOIN plant_natives ON natives.id = plant_natives.native_id").
		Where("plant_id", id).
		Find(&natives).
		Error; err != nil {
		return nil, err
	}
	plant.Natives = natives
	return plant, nil
}

// Update plant and store it into database
func (repo *repository) Update(id int, plant plant.Plant) error {
	if err := repo.db.Model(&plant).
		Where("id", id).
		Updates(
			plantModel{
				Name:          plant.Name,
				BotanicalName: plant.BotanicalName,
				Type:          plant.Type,
				Difficulty:    plant.Difficulty,
				Description:   plant.Description,
				WateringTime:  plant.WateringTime,
				HowToGrow:     plant.HowToGrow,
				Soil:          plant.Soil,
			},
		).Error; err != nil {
		return err
	}
	return nil
}

// Delete plant by given id
func (repo *repository) Delete(id int) error {
	plant := &plant.Plant{}
	if err := repo.db.Delete(&plant, id).Error; err != nil {
		return err
	}
	return nil
}
